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
	for _, line := range strings.Split(code, "\n") {
		statement := strings.Split(line, " ")
		cmd := statement[0]
		for _, op := range core.OpCodes {
			if cmd != op.Pat {
				continue
			}

			bin = append(bin, op.Op)
			for _, arg := range statement[1:] {
				b, err := convertNum(arg, op.ArgSize)
				if err != nil {
					return nil, err
				}
				bin = append(bin, b...)
			}
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
