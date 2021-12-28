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
	StackPointer   int
	ExePointer     int
	ReadOnlyOffset int
}

func (i *Interp) Run() error {
	// TODO: this should probably be done in a for/while loop
	// TODO: should I have a done code or just when execution
	// gets to the end of the code?
	if i.ExePointer >= len(i.Memory) {
		// execution is finished
		return nil
	}
	op := i.Memory[i.ExePointer]
	if core.OpCodes[op].Fn == nil {
		fmt.Printf("unimplmented op code %s\n", core.OpCodes[op].Pat)
	}
	deltaEPtr, deltaSPtr := core.OpCodes[op].Fn(i.Memory[:], i.ExePointer, i.StackPointer)
	i.ExePointer += deltaEPtr
	i.StackPointer += deltaSPtr
	// NOTE: this stuff isn't really adding any security since the pop command can write to any address
	// we need to rethink this, but i'm taking it out for now
	//
	// if i.ExePointer < i.ReadOnlyOffset {
	// 	return errors.New(fmt.Sprintf("invalid execution pointer %d, read only region starts at %d", i.ExePointer, i.ReadOnlyOffset))
	// }
	// if i.StackPointer > i.ReadOnlyOffset {
	// 	return errors.New(fmt.Sprintf("invalid stack pointer %d, read only region starts at %d", i.StackPointer, i.ReadOnlyOffset))
	// }
	fmt.Println("OP: ", core.OpCodes[op].Pat)
	fmt.Println("EXEPT: ", i.ExePointer, "STACKPT: ", i.StackPointer)
	fmt.Printf("MEM: ...%#v\n", i.Memory[8185:])
	return nil
}

func (i *Interp) LoadCode(byteCode []byte) error {
	if len(byteCode) > maxByteCodeLen {
		return errors.New(fmt.Sprintf("byte code is %d bytes, max length is %d", len(byteCode), maxByteCodeLen))
	}

	offset := len(i.Memory) - len(byteCode)
	for c := 0; c < len(byteCode); c++ {
		i.Memory[offset+c] = byteCode[c]
	}

	i.ReadOnlyOffset = offset
	i.StackPointer = offset
	i.ExePointer = offset
	return nil
}
