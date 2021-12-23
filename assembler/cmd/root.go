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

var rootCmd = &cobra.Command{
	Use:   "NeuBi",
	Short: "NeuBi is an assembler that assembles neu byte code",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		if !strings.HasSuffix(filepath, ".nb") {
			return errors.New(fmt.Sprintf("Invalid neu byte code file %s, file must end with .nb extension\n", filepath))
		}

		code, err := ioutil.ReadFile(filepath)
		if err != nil {
			return errors.New(fmt.Sprintf("Error: unable to read %s, %s\n", filepath, err))
		}

		asm, err := assemble(string(code))
		if err != nil {
			return errors.New(fmt.Sprintf("Error: unable to assemble code %s\n", err))
		}

		fmt.Printf("%#v\n", asm)

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
	for _, base := range []int{2, 10, 16} {
		i, err := strconv.ParseInt(num, base, int(size))
		if err != nil {
			continue
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

	return nil, errors.New("unable to convert number as either base 2, 10 or 16")
}

func Execute() error {
	return rootCmd.Execute()
}
