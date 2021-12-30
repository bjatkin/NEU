package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bjatkin/neu_interpreter/core"
)

type debugCMD struct {
	cmd         []string
	description string
	usage       string
	fn          func(*Debugger, []string) bool
}

type Debugger struct {
	i        *Interp
	cmds     []debugCMD
	previous string
}

func NewDebugger(interp *Interp) *Debugger {
	return &Debugger{
		i: interp,
		cmds: []debugCMD{
			{
				cmd:         []string{"?", "help"},
				description: "print out help info",
				usage:       "?|help",
				fn: func(d *Debugger, args []string) bool {
					fmt.Println("  ==== available commands ====")
					var usageLen int
					for _, cmd := range d.cmds {
						if len(cmd.usage) > usageLen {
							usageLen = len(cmd.usage)
						}
					}

					pad := func(base string, i int) string {
						for len(base) < i {
							base += " "
						}
						return base
					}

					for _, cmd := range d.cmds {
						fmt.Printf("    %s - %s\n", pad(cmd.usage, usageLen), cmd.description)
					}
					return false
				},
			},
			{
				cmd:         []string{">", "step"},
				description: "execute the next expression and then break again",
				usage:       ">|step",
				fn: func(d *Debugger, args []string) bool {
					d.i.Break = true
					return true
				},
			},
			{
				cmd:         []string{">>", "resume"},
				description: "resume executing as normal",
				usage:       ">>|resume",
				fn: func(d *Debugger, args []string) bool {
					return true
				},
			},
			{
				cmd:         []string{"q", "quit", "exit"},
				description: "quit the interpreter",
				usage:       "q|quit",
				fn: func(d *Debugger, args []string) bool {
					fmt.Println("  quitting interpreter")
					os.Exit(0)
					return true
				},
			},
			{
				cmd:         []string{"cmd"},
				description: "show the next command that will be run (or the cmd at addr if provided)",
				usage:       "cmd [?add]",
				fn: func(d *Debugger, args []string) bool {
					addr := d.i.ExePointer + d.i.ReadOnlyOffset
					if len(args) > 0 {
						conv, err := parseNumeric(args[0])
						if err != nil {
							fmt.Printf("  invalid arg %s: %s\n", args[0], err)
							return false
						}

						addr = uint(conv)
					}

					op := d.i.Memory[addr]
					nextOp := core.OpCodes[op]
					fmt.Printf("  0x%x: %s\n", addr, opString(d.i, nextOp))
					return false
				},
			},
			{
				cmd:         []string{"s", "stack"},
				description: "print out the stack",
				usage:       "s|stack",
				fn: func(d *Debugger, args []string) bool {
					if d.i.StackPointer == d.i.ReadOnlyOffset {
						fmt.Println("  [ empty ]")
					}

					for i := d.i.StackPointer; i < d.i.ReadOnlyOffset; i++ {
						if i == d.i.StackPointer {
							fmt.Printf("  %08x | %02x | <[%d]\n", i, d.i.Memory[i], d.i.ReadOnlyOffset-d.i.StackPointer)
							continue
						}

						sep := "|"
						if ((i-d.i.StackPointer)+1)%4 == 0 {
							sep = "+"
						}
						fmt.Printf("  %08x %s %02x %s\n", i, sep, d.i.Memory[i], sep)
					}
					fmt.Print("\n")

					return false
				},
			},
			{
				cmd:         []string{"spt"},
				description: "print out the stack pointer",
				usage:       "spt",
				fn: func(d *Debugger, args []string) bool {
					fmt.Printf("  [ 0x%08x | %d ]\n", d.i.StackPointer, d.i.StackPointer)
					return false
				},
			},
			{
				cmd:         []string{"ept"},
				description: "print out the execution pointer",
				usage:       "ept",
				fn: func(d *Debugger, args []string) bool {
					fmt.Printf("  local:  [ 0x%08x | %d ]\n  global: [ 0x%x|%d ]\n",
						d.i.ExePointer,
						d.i.ExePointer,
						d.i.ExePointer+d.i.ReadOnlyOffset,
						d.i.ExePointer+d.i.ReadOnlyOffset,
					)
					return false
				},
			},
			{

				cmd:         []string{"m", "mem"},
				description: "print out values in memory addresses start to end",
				usage:       "m|mem   [start][?:end]",
				fn: func(d *Debugger, args []string) bool {
					if len(args) < 1 {
						fmt.Println("  no memory range provided")
						return false
					}
					addrs := strings.Split(args[0], ":")

					startArg := addrs[0]
					start, err := parseNumeric(startArg)
					if err != nil {
						fmt.Printf("  invalid starting memory address %s: %s\n", startArg, err)
						return false
					}

					endArg := startArg
					if len(addrs) < 1 {
						endArg = addrs[1]
					}
					end, err := parseNumeric(endArg)
					if err != nil {
						fmt.Printf("  invalid ending memory address %s: %s\n", endArg, err)
						return false
					}

					if start > end {
						fmt.Printf("  end address is less than start address: 0x%08x, 0x%08x\n", start, end)
					}

					if end >= uint(len(d.i.Memory)) {
						fmt.Printf("  end address is invalid 0x%08x, final memory address 0x%08x\n", end, len(d.i.Memory))
					}

					if start == end {
						end++
					}

					lines := core.FmtByteArray(start, d.i.Memory[start:end])
					for i := 0; i < len(lines); i++ {
						lines[i] = "  " + lines[i]
					}

					fmt.Printf("\nMEMORY:\n%s\n\n", strings.Join(lines, "\n"))
					return false
				},
			},
			{
				cmd:         []string{"!"},
				description: "re-run the last successful command",
				usage:       "!",
			},
		},
	}
}

func (d *Debugger) Run() error {
	reader := bufio.NewReader(os.Stdin)

	opIndex := d.i.Memory[d.i.ExePointer+d.i.ReadOnlyOffset]
	op := core.OpCodes[opIndex]
	fmt.Printf("break point hit: %s (type ? for help)\n", opString(d.i, op))

	for {
		fmt.Printf("debug: ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if cmd == "\n" { // empty command
			continue
		}

		if cmd == "!\n" { // run the last valid command
			cmd = d.previous
		}

		cmdSlice := strings.Split(cmd[:len(cmd)-1], " ")
		if len(cmdSlice) == 0 { // empty command
			continue
		}

		debug, valid := d.CMD(cmdSlice[0])
		if !valid {
			fmt.Printf("  unknown command '%s'\n", cmd)
			continue
		}

		d.previous = cmd

		done := debug.fn(d, cmdSlice[1:])

		if done {
			return nil
		}
	}

}

func (d *Debugger) CMD(find string) (*debugCMD, bool) {
	for _, cmd := range d.cmds {
		for _, s := range cmd.cmd {
			if find == s {
				return &cmd, true
			}
		}
	}
	return nil, false
}

func opString(i *Interp, op core.OpCode) string {
	pat := op.Pat
	var arg string
	switch op.ArgSize {
	case 8:
		b := i.Memory[i.ExePointer+i.ReadOnlyOffset+1]
		arg = fmt.Sprintf(" %d | 0x%02x", b, b)
	case 16:
		i16 := *core.Asi16(&i.Memory[i.ExePointer+i.ReadOnlyOffset+1])
		arg = fmt.Sprintf(" %d | 0x%04x", i16, i16)
	case 32:
		i32 := *core.Asi32(&i.Memory[i.ExePointer+i.ReadOnlyOffset+1])
		arg = fmt.Sprintf(" %d | 0x%08x", i32, i32)
	case 64:
		i64 := *core.Asi64(&i.Memory[i.ExePointer+i.ReadOnlyOffset+1])
		arg = fmt.Sprintf(" %d | 0x%016x", i64, i64)
	}

	return fmt.Sprintf("[ %s%s ]", pat, arg)
}

func parseNumeric(addr string) (uint, error) {
	base := 10
	switch {
	case strings.HasPrefix(addr, "0b"):
		base = 2
		addr = strings.TrimPrefix(addr, "0b")
	case strings.HasPrefix(addr, "0x"):
		base = 16
		addr = strings.TrimPrefix(addr, "0x")
	}

	conv, err := strconv.ParseUint(addr, base, 64)
	if err != nil {
		return 0, err
	}

	return uint(conv), nil
}
