package main

import "fmt"

type stringMatcher struct {
	pattern string
	opCode  byte
}

func newStringMatcher(pattern string, opCode byte) *stringMatcher {
	return &stringMatcher{
		pattern: pattern,
		opCode:  opCode,
	}
}

func (m *stringMatcher) match(check string) bool {
	fmt.Println("checking ", m.pattern)
	if m.pattern == check {
		fmt.Println("\t\t- matched!")
		return true
	}
	return false
}

func (m *stringMatcher) value() []byte {
	return []byte{m.opCode}
}
