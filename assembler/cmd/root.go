package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/bjatkin/neu_interpreter/core"
	"github.com/spf13/cobra"
)

var (
	outputFile string
	print      bool
)

func init() {
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "", "the output file to write the compiled .n file to")
	rootCmd.PersistentFlags().BoolVar(&print, "print", false, "print the binary data to the console")

	rootCmd.AddCommand(fmtCmd)
}

var rootCmd = &cobra.Command{
	Use:   "NeuBi [byte code file]",
	Short: "NeuBi is an assembler that assembles neu byte code",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		if !strings.HasSuffix(filepath, ".nb") {
			return errors.New(fmt.Sprintf("Invalid neu byte code file %s, file must end with .nb extension\n", filepath))
		}

		if outputFile == "" {
			outputFile = strings.Split(filepath, ".")[0]
		}

		code, err := ioutil.ReadFile(filepath)
		if err != nil {
			return errors.New(fmt.Sprintf("Error: unable to read %s, %s\n", filepath, err))
		}

		asm, err := assemble(string(code))
		if err != nil {
			return errors.New(fmt.Sprintf("Error: unable to assemble code %s\n", err))
		}

		if print {
			lines := core.FmtByteArray(0, asm)
			fmt.Printf("\nEXE:\n%s\n\n", strings.Join(lines, "\n"))
		}

		i := strings.Index(filepath, ".nb")
		outFilepath := filepath[:i] + ".n"
		err = ioutil.WriteFile(outFilepath, asm, 0777)
		if err != nil {
			return errors.New(fmt.Sprintf("Error: unable to write code to %s, %s\n", outFilepath, err))
		}

		return nil
	},
}

func assemble(code string) ([]byte, error) {
	var bin []byte

	// filter out white space so nb code can be
	// alligned with spaces
	// also remove comments
	filter := func(base []string) []string {
		var f []string
		for i := 0; i < len(base); i++ {
			if base[i] == "" {
				continue
			}
			if strings.HasPrefix(base[i], "//") {
				break // ignore everythign after comments
			}

			f = append(f, base[i])
		}
		return f
	}

	// replace tabs with spaces
	code = strings.ReplaceAll(code, "\t", "  ")

	// pass 1: filter out spaces and comments
	// fill out the named constants map
	// rewrite pointer (#) vs literal commands
	namedConsts := map[string]string{}
	var expressions [][]string
	for _, line := range strings.Split(code, "\n") {
		expr := filter(strings.Split(line, " "))
		if len(expr) == 0 {
			// don't include empty lines
			continue
		}

		if core.IsNamedConst(expr) {
			// TODO: error out of there are conflicting names
			namedConsts[expr[0]] = expr[2]

			// named constand declaration lines are not included
			// in the final byte code and are just available
			// for the convience of the writer
			continue
		}

		if core.IsAddrCMD(expr) {
			expr[0] += "#"
			expr[1] = expr[1][1:]
		}

		expressions = append(expressions, expr)
	}

	// pass 2: replace all the named constants in the code
	for i := 0; i < len(expressions); i++ {
		expr := expressions[i]

		if len(expr) < 2 {
			continue
		}

		if c, ok := namedConsts[expr[1]]; ok {
			expr[1] = c
		}
	}

	labelMap := map[string]int{}
	labelReplace := map[string][]int{}
	// pass 3: convert the code to binary
	// fill out the label map, and the label arg index
	for _, expr := range expressions {
		cmd := expr[0]
		if core.IsLabel(cmd) {
			labelMap[cmd] = len(bin)
			// labels aren't included in the byte code
			continue
		}

		var found bool
		for _, op := range core.OpCodes {
			if cmd != op.Pat {
				continue
			}

			found = true
			bin = append(bin, op.Op)
			var argCount int
			for _, arg := range expr[1:] {
				argCount++
				if core.IsLabel(arg) {
					labelReplace[arg] = append(labelReplace[arg], len(bin))
					bin = append(bin, core.I64tob(0)...)
					continue
				}

				b, err := convertNum(arg, op.ArgSize)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("unable to parse expression '%s' %s", strings.Join(expr, " "), err))
				}
				bin = append(bin, b...)
			}

			if (argCount > 0) != core.ExpectArg(op.Op) {
				return nil, errors.New(fmt.Sprintf("invalid expression '%s', op %s has invalid number of arguments", strings.Join(expr, " "), cmd))
			}
		}
		if !found {
			return nil, errors.New(fmt.Sprintf("unable to write command to exe '%s'", cmd))
		}
	}

	// replace all the labels
	for label, idxs := range labelReplace {
		if _, ok := labelMap[label]; !ok {
			return nil, errors.New("undefined label " + label)
		}

		l := labelMap[label]
		replace := core.I64asb(&l)
		for _, i := range idxs {
			bin[i] = replace[0]
			bin[i+1] = replace[1]
			bin[i+2] = replace[2]
			bin[i+3] = replace[3]
			bin[i+4] = replace[4]
			bin[i+5] = replace[5]
			bin[i+6] = replace[6]
			bin[i+7] = replace[7]
		}
	}

	return bin, nil
}

func convertNum(num string, size byte) ([]byte, error) {
	base := 10
	switch {
	case strings.HasPrefix(num, "0b"):
		base = 2
		num = strings.TrimPrefix(num, "0b")
	case strings.HasPrefix(num, "0x"):
		base = 16
		num = strings.TrimPrefix(num, "0x")
	}

	parseErr := func(err error) error {
		return errors.New(fmt.Sprintf("unable to convert number '%s' as a base %d and size %d number: %s", num, base, size, err))
	}

	switch size {
	case 8:
		i, err := strconv.ParseUint(num, base, int(size))
		if err != nil {
			return nil, parseErr(err)
		}
		return []byte{byte(i)}, nil
	case 16:
		i, err := strconv.ParseInt(num, base, int(size))
		if err != nil {
			return nil, parseErr(err)
		}

		i16 := int16(i)
		b16 := core.I16asb(&i16)
		return b16[:], nil
	case 32:
		i, err := strconv.ParseInt(num, base, int(size))
		if err != nil {
			return nil, parseErr(err)
		}

		i32 := int32(i)
		b32 := core.I32asb(&i32)
		return b32[:], nil
	case 64:
		i, err := strconv.ParseInt(num, base, int(size))
		if err != nil {
			return nil, parseErr(err)
		}

		i64 := int(i)
		b64 := core.I64asb(&i64)
		return b64[:], nil
	default:
		return nil, errors.New(fmt.Sprintf("invalid size %d for converting number", size))
	}
}

func Execute() error {
	return rootCmd.Execute()
}
