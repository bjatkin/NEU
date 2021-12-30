package core

import (
	"fmt"
	"strings"
	"unsafe"
)

const (
	IntWidth = 8
)

// NOTE: doing this this way is much faster but opens up some potential security issues
// for example, you can now grab stuff under the stack by popping the stack down to a single
// byte and then popping an int64 off the stack, because we're doing things with unsafe
// pointers there is no bounds checking going on. How big a deal is letting you access the
// top 7 bytes of the execution portion of memory?
// the only security flaw I can think of is loading a very short program that is less than 7
// bytes to take advantage of this to read/ write data into abitrary memory.
// in practice that seems tricky to do but is still probably worth looking into.
// perhaps we just load some data between the stack and the code to prevent any funny
// business?
func Asi64(pt *byte) *int {
	return (*int)(unsafe.Pointer(pt))
}

func I64asb(pt *int) [8]byte {
	return (*(*[8]byte)(unsafe.Pointer(pt)))
}

func Asi32(pt *byte) *int32 {
	return (*int32)(unsafe.Pointer(pt))
}

func I32asb(pt *int32) [4]byte {
	return (*(*[4]byte)(unsafe.Pointer(pt)))
}

func Asi16(pt *byte) *int16 {
	return (*int16)(unsafe.Pointer(pt))
}

func I16asb(pt *int16) [2]byte {
	return (*(*[2]byte)(unsafe.Pointer(pt)))
}

func Asu64(pt *byte) *uint {
	return (*uint)(unsafe.Pointer(pt))
}

func U64asb(pt *uint) [8]byte {
	return (*(*[8]byte)(unsafe.Pointer(pt)))
}

func Asu32(pt *byte) *uint32 {
	return (*uint32)(unsafe.Pointer(pt))
}

func U32asb(pt *uint32) [4]byte {
	return (*(*[4]byte)(unsafe.Pointer(pt)))
}

func Asu16(pt *byte) *uint16 {
	return (*uint16)(unsafe.Pointer(pt))
}

func U16asb(pt *uint16) [2]byte {
	return (*(*[2]byte)(unsafe.Pointer(pt)))
}

func I64tob(from int) (to []byte) {
	ret := []byte{
		byte(from),
		byte(from >> 8),
		byte(from >> 16),
		byte(from >> 24),
		byte(from >> 32),
		byte(from >> 40),
		byte(from >> 48),
		byte(from >> 56),
	}
	return ret
}

func Btoi64(from []byte) (to int) {
	if len(from) != 8 {
		fmt.Printf("from must be of length 8, was %d\n", len(from))
		panic(1)
	}

	return int(from[0]) |
		int(from[1])<<8 |
		int(from[2])<<16 |
		int(from[3])<<24 |
		int(from[4])<<32 |
		int(from[5])<<40 |
		int(from[6])<<48 |
		int(from[7])<<56
}

func I32tob(from int32) (to []byte) {
	ret := []byte{
		byte(from),
		byte(from >> 8),
		byte(from >> 16),
		byte(from >> 24),
	}
	return ret
}

func Btoi32(from []byte) (to int32) {
	if len(from) != 4 {
		fmt.Printf("from must be of length 4, was %d\n", len(from))
		panic(1)
	}

	return int32(from[0]) |
		int32(from[1])<<8 |
		int32(from[2])<<16 |
		int32(from[3])<<24
}

func I16tob(from int16) (to []byte) {
	ret := []byte{
		byte(from),
		byte(from >> 8),
	}
	return ret
}

func Btoi16(from []byte) (to int16) {
	if len(from) != 2 {
		fmt.Printf("from must be of length 2, was %d\n", len(from))
		panic(1)
	}

	return int16(from[0]) |
		int16(from[1])<<8
}

func IsLabel(test string) bool {
	if len(test) < 3 {
		return false
	}
	if len(test) > 25 {
		return false
	}

	return test[0] == '[' && test[len(test)-1] == ']'
}

func IsNamedConst(test []string) bool {
	if len(test) != 3 {
		return false
	}

	if !IsName(test[0]) {
		return false
	}

	if test[1] != "=" {
		return false
	}

	if len(test[2]) == 0 {
		return false
	}

	return true
}

// IsName checks if the given string is a valid name
// names always start with \ and must have at leas on additional character
// names can not contain spaces
func IsName(test string) bool {
	if strings.ContainsAny(test, " ") {
		return false
	}

	if len(test) < 2 || len(test) > 25 {
		return false
	}

	if test[0] != '\\' {
		return false
	}

	return true
}

func IsAddrCMD(test []string) bool {
	// addr comands have a least 1 arg
	if len(test) < 2 || test[1] == "" {
		return false
	}

	// the first arg must start with '#'
	if test[1][0] != '#' {
		return false
	}

	return true
}

func ExpectArg(op byte) bool {
	if OpCodes[op].ArgSize > 0 {
		return true
	}
	return false
}

func FmtByteArray(memoryOffset uint, data []byte) []string {
	var lines, line []string
	for i, b := range data {
		line = append(line, fmt.Sprintf("%02X", b))

		if (i+1)%16 == 0 || i == len(data)-1 {
			lines = append(lines, fmt.Sprintf("%s | %s", fmt.Sprintf("%08X", memoryOffset), strings.Join(line, " ")))
			line = []string{}
			memoryOffset += 16
		}
	}
	return lines
}
