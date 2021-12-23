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
	op := i.Memory[i.ExePointer]
	deltaEPtr, deltaSPtr := core.OpCodes[op].Fn(i.Memory[:], i.ExePointer, i.StackPointer)
	i.ExePointer += deltaEPtr
	i.StackPointer += deltaSPtr
	if i.ExePointer < i.ReadOnlyOffset {
		return errors.New(fmt.Sprintf("invalid execution pointer %d, read only region starts at %d", i.ExePointer, i.ReadOnlyOffset))
	}
	if i.StackPointer > i.ReadOnlyOffset {
		return errors.New(fmt.Sprintf("invalid stack pointer %d, read only region starts at %d", i.StackPointer, i.ReadOnlyOffset))
	}
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
