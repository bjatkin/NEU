package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "", "the output file to write the compiled phone to")
	rootCmd.PersistentFlags().BoolVar(&print, "print", false, "print the binary data to the console")

	// TODO: I need a different setup to get this to work the way I want
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
			printByteArray(asm)
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

	labelMap := map[string]uint{}
	labelReplace := map[string][]int{}
	// pass 3: convert the code to binary
	// fill out the label map, and the label arg index
	for _, expr := range expressions {
		cmd := expr[0]
		if core.IsLabel(cmd) {
			labelMap[cmd] = uint(len(bin))
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
			for _, arg := range expr[1:] {
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

		replace := core.I64tob(labelMap[label])
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

	i, err := strconv.ParseUint(num, base, int(size))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to convert number '%s' as a base %d number: %s", num, base, err))
	}

	switch size {
	case 8:
		return []byte{byte(i)}, nil
	case 16:
		return core.I16tob(uint16(i)), nil
	case 32:
		return core.I32tob(uint32(i)), nil
	case 64:
		return core.I64tob(uint(i)), nil
	default:
		return nil, errors.New(fmt.Sprintf("invalid size %d for converting number", size))
	}
}

func printByteArray(asm []byte) {
	var lines, line []string
	var memoryOffset int
	for i, b := range asm {
		line = append(line, padStr(fmt.Sprintf("%X", b), 2))

		if (i+1)%16 == 0 || i == len(asm)-1 {
			lines = append(lines, fmt.Sprintf("%s | %s", padStr(fmt.Sprintf("%X", memoryOffset), 8), strings.Join(line, " ")))
			line = []string{}
			memoryOffset += 16
		}
	}
	fmt.Printf("\nEXE:\n%s\n\n", strings.Join(lines, "\n"))
}

func padStr(s string, l int) string {
	for len(s) < l {
		s = "0" + s
	}

	if len(s) > l {
		log.Fatalf("input value '%s' was longer than specified length %d\n", s, l)
	}

	return s
}

func Execute() error {
	return rootCmd.Execute()
}
