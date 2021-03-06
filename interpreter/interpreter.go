package main

import (
	"errors"
	"fmt"

	"github.com/bjatkin/neu_interpreter/core"
)

const (
	maxByteCodeLen = 20 * 1024 // 20k is the max size of your code (leaves 12k of execution memory)
)

type Interp struct {
	Memory         [32 * 1024]byte // 32k of main memory
	StackPointer   uint
	ExePointer     uint
	ReadOnlyOffset uint
	Break          bool // break after the next step
	Debugger       *Debugger
}

func (i *Interp) Run() error {
	for int(i.ExePointer+i.ReadOnlyOffset) < len(i.Memory) {
		if i.ExePointer < 0 {
			return errors.New(fmt.Sprintf("the execution pointer is invalid ePt: %d", i.ExePointer))
		}

		op := i.Memory[i.ExePointer+i.ReadOnlyOffset]
		if core.OpCodes[op].Fn == nil {
			return errors.New(fmt.Sprintf("uknown byte code 0x%x\n", op))
		}

		if op == 0x49 { // break point
			i.Break = true
		}

		i.StackPointer, i.ExePointer = core.OpCodes[op].Fn(i.Memory[:i.ReadOnlyOffset], i.Memory[i.ReadOnlyOffset:], i.StackPointer, i.ExePointer)
		if int(i.StackPointer) > int(i.ReadOnlyOffset) {
			return errors.New(fmt.Sprintf("stack size is less than zero sPt: 0x%x, memorySize: 0x%x", i.StackPointer, len(i.Memory)))
		}

		if i.Break { // break execution and run the debugger
			i.Break = false

			err := i.Debugger.Run()
			if err != nil {
				return err
			}
		}
	}

	// print to debug bytes to see if stuff is working as expected
	// reset execution to run again on the next update loop
	i.ExePointer = 0
	return nil
}

func (i *Interp) LoadCode(byteCode []byte) error {
	if len(byteCode) > maxByteCodeLen {
		return errors.New(fmt.Sprintf("byte code is %d bytes, max length is %d", len(byteCode), maxByteCodeLen))
	}

	offset := uint(len(i.Memory) - len(byteCode))
	for c := 0; c < len(byteCode); c++ {
		i.Memory[int(offset)+c] = byteCode[c]
	}

	i.StackPointer = offset
	i.ReadOnlyOffset = offset
	return nil
}
