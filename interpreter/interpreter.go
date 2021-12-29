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
	Memory       [8 * 1024]byte // main memory
	Code         []byte         // the code to execute
	StackPointer int
	ZeroStack    int
	ExePointer   int
}

func (i *Interp) Run() error {
	// TODO: this should probably be done in a for/while loop
	// TODO: should I have a done code or just when execution
	// gets to the end of the code?
	if i.ExePointer >= len(i.Memory) {
		// execution is finished
		return nil
	}
	op := i.Code[i.ExePointer]
	if core.OpCodes[op].Fn == nil {
		fmt.Printf("uknown byte code 0x%x\n", op)
	}

	fmt.Printf("do: %s\n", core.OpCodes[op].Pat)
	fmt.Printf("stack: 0x%x, [0x%x]\n", i.Memory[i.StackPointer], i.StackPointer)
	fmt.Printf("#i: 0x%x\n", i.Memory[0x410])

	deltaEPtr, deltaSPtr := core.OpCodes[op].Fn(i.Memory[:], i.Code, i.ExePointer, i.StackPointer)
	i.ExePointer += deltaEPtr
	i.StackPointer += deltaSPtr
	if i.StackPointer > len(i.Memory) {
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

	i.Code = byteCode
	i.StackPointer = len(i.Memory)
	return nil
}
