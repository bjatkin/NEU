package main

import (
	"errors"
	"fmt"

	"github.com/bjatkin/neu_interpreter/core"
)

const (
	maxByteCodeLen = 4096
)

type Interp struct {
	Memory         [8 * 1024]byte // main memory
	StackPointer   uint
	ExePointer     uint
	ReadOnlyOffset uint
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

		i.StackPointer, i.ExePointer = core.OpCodes[op].Fn(i.Memory[:i.ReadOnlyOffset], i.Memory[i.ReadOnlyOffset:], i.StackPointer, i.ExePointer)
		if int(i.StackPointer) > int(i.ReadOnlyOffset) {
			return errors.New(fmt.Sprintf("stack size is less than zero sPt: 0x%x, memorySize: 0x%x", i.StackPointer, len(i.Memory)))
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
