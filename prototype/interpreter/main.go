package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filepath := os.Args[1]
	code, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#v\n\n", code)
	stack := &stack{ptr: -1}
	opCodes := []OpCode{
		{
			name: "stack size",
			code: 0x05,
			run:  setStackSize,
		},
		{
			name: "add",
			code: 0x02,
			run:  add,
		},
		{
			name: "push",
			code: 0x01,
			run:  push,
		},
	}
	for ptr := 0; ptr < len(code); ptr++ {
		var run bool
		for _, op := range opCodes {
			if code[ptr] == op.code {
				ptr += op.run(stack, code, ptr)
				run = true
				break
			}
		}
		if !run {
			log.Fatalf("unknown op code 0x%x\n", code[ptr])
		}
		stack.print()
	}
}

type stack struct {
	data []int
	ptr  int
}

func (s *stack) print() {
	fmt.Printf("\n")
	for i, d := range s.data {
		if s.ptr == i {
			fmt.Printf(" > | %d |\n", d)
		} else {
			fmt.Printf("   | %d |\n", d)
		}
	}
	fmt.Printf("\n")
}

type OpCode struct {
	name string
	code byte
	run  func(s *stack, code []byte, ptr int) int
}

func setStackSize(s *stack, code []byte, ptr int) int {
	size := bytesToInt(code[ptr+1 : ptr+9])
	fmt.Printf("setting length to %d, %d", size, code[ptr+1:ptr+9])
	s.data = make([]int, size)
	return 8
}

func add(s *stack, code []byte, ptr int) int {
	s.ptr--
	s.data[s.ptr] = s.data[s.ptr] + s.data[s.ptr+1]
	return 0
}

func push(s *stack, code []byte, ptr int) int {
	s.ptr++
	s.data[s.ptr] = bytesToInt(code[ptr+1 : ptr+9])
	return 8
}
