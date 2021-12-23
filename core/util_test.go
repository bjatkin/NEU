package core

import (
	"testing"
)

func TestI64tob(t *testing.T) {
	tests := []uint{
		50,
		1029384,
		0xffffffffffffffff,
		0,
	}

	for _, i := range tests {
		if c := Btoi64(I64tob(i)); c != i {
			t.Errorf("I64tob| failed to convert from into to binary and back, got: %d, want: %d", c, i)
		}
	}
}

func TestI32tob(t *testing.T) {
	tests := []uint32{
		2837,
		102984,
		0xffffffff,
		0,
	}

	for _, i := range tests {
		if c := Btoi32(I32tob(i)); c != i {
			t.Errorf("I32tob| failed to convert from into to binary and back, got: %d, want: %d", c, i)
		}
	}
}

func TestI16tob(t *testing.T) {
	tests := []uint16{
		2837,
		10293,
		0xffff,
		0,
	}

	for _, i := range tests {
		if c := Btoi16(I16tob(i)); c != i {
			t.Errorf("I16tob| failed to convert from into to binary and back, got: %d, want: %d", c, i)
		}
	}
}
