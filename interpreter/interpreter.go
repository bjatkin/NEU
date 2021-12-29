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
	// TODO: this should probably be done in a for/while loop
	// TODO: should I have a done code or just when execution
	// gets to the end of the code?
	if int(i.ExePointer) >= len(i.Memory) {
		// execution is finished
		return nil
	}
	if i.ExePointer < i.ReadOnlyOffset {
		return errors.New(fmt.Sprintf("the execution pointer is invalid ePt: %d, readOnlyOffset: %d", i.ExePointer, i.ReadOnlyOffset))
	}

	op := i.Memory[i.ExePointer]
	if core.OpCodes[op].Fn == nil {
		fmt.Printf("uknown byte code 0x%x\n", op)
	}

	fmt.Printf("do: %s\n", core.OpCodes[op].Pat)
	fmt.Printf("stack: 0x%x, [0x%x]\n", i.Memory[i.StackPointer], i.StackPointer)
	fmt.Printf("#i: 0x%x\n", i.Memory[0x410])

	i.StackPointer, i.ExePointer = core.OpCodes[op].Fn(i.Memory[:i.ReadOnlyOffset], i.Memory[i.ReadOnlyOffset:], i.StackPointer, i.ExePointer)
	if int(i.StackPointer) > int(i.ReadOnlyOffset) {
		return errors.New(fmt.Sprintf("stack size is less than zero sPt: 0x%x, memorySize: 0x%x", i.StackPointer, len(i.Memory)))
	}

	fmt.Printf("#i: 0x%x\n", i.Memory[0x410])
	fmt.Printf("stack: 0x%x, [0x%x]\n", i.Memory[i.StackPointer], i.StackPointer)
	fmt.Printf("do next: %s %x[%x]\n", core.OpCodes[i.Memory[i.ExePointer]].Pat, i.Memory[i.ExePointer], i.ExePointer)
	fmt.Println("--------------------------------------")
	return nil
}

func (i *Interp) LoadCode(byteCode []byte) error {
	if len(byteCode) > maxByteCodeLen {
		return errors.New(fmt.Sprintf("byte code is %d bytes, max length is %d", len(byteCode), maxByteCodeLen))
	}

	offset := uint(len(byteCode) - len(i.Memory))
	for c := 0; c < len(byteCode); c++ {
		i.Memory[int(offset)+c] = byteCode[c]
	}

	i.StackPointer = offset
	i.ExePointer = offset
	i.ReadOnlyOffset = offset
	return nil
}
