package core

import "fmt"

const (
	IntWidth = 8
)

func I64tob(from uint) (to []byte) {
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

func Btoi64(from []byte) (to uint) {
	if len(from) != 8 {
		fmt.Printf("from must be of length 8, was %d\n", len(from))
		panic(1)
	}

	return uint(from[0]) |
		uint(from[1])<<8 |
		uint(from[2])<<16 |
		uint(from[3])<<24 |
		uint(from[4])<<32 |
		uint(from[5])<<40 |
		uint(from[6])<<48 |
		uint(from[7])<<56
}

func I32tob(from uint32) (to []byte) {
	ret := []byte{
		byte(from),
		byte(from >> 8),
		byte(from >> 16),
		byte(from >> 24),
	}
	return ret
}

func Btoi32(from []byte) (to uint32) {
	if len(from) != 4 {
		fmt.Printf("from must be of length 4, was %d\n", len(from))
		panic(1)
	}

	return uint32(from[0]) |
		uint32(from[1])<<8 |
		uint32(from[2])<<16 |
		uint32(from[3])<<24
}

func I16tob(from uint16) (to []byte) {
	ret := []byte{
		byte(from),
		byte(from >> 8),
	}
	return ret
}

func Btoi16(from []byte) (to uint16) {
	if len(from) != 2 {
		fmt.Printf("from must be of length 2, was %d\n", len(from))
		panic(1)
	}

	return uint16(from[0]) |
		uint16(from[1])<<8
}