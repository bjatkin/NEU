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

func TestIsLabel(t *testing.T) {
	tests := []struct {
		name  string
		label string
		want  bool
	}{
		{
			"too short",
			"[]",
			false,
		},
		{
			"too long",
			"[this label is super super super long, like longer than any label really needs to be ever ever ever ever, like never, you should never make a label this long]",
			false,
		},
		{
			"missing brackes",
			"LABEL",
			false,
		},
		{
			"nested brackets",
			"[[[][]]]",
			true,
		},
		{
			"symboles",
			"[~>_test]",
			true,
		},
		{
			"normal",
			"[test_label]",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLabel(tt.label); got != tt.want {
				t.Errorf("IsLabel| got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNamedConst(t *testing.T) {
	tests := []struct {
		name string
		line []string
		want bool
	}{
		{
			"too short",
			[]string{},
			false,
		},
		{
			"too long",
			[]string{`\test`, "=", "0x00", "0x00"},
			false,
		},
		{
			"missing equal",
			[]string{`\test`, "is", "0x00"},
			false,
		},
		{
			"invalid name",
			[]string{"bad_name", "=", "0x00"},
			false,
		},
		{
			"empty value",
			[]string{`\test`, "=", ""},
			false,
		},
		{
			"good",
			[]string{`\test`, "=", "0x00"},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNamedConst(tt.line); got != tt.want {
				t.Errorf("IsNamedConst| got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsName(t *testing.T) {
	tests := []struct {
		name string
		test string
		want bool
	}{
		{
			"invalid",
			"fail",
			false,
		},
		{
			"valid",
			`\test`,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsName(tt.test); got != tt.want {
				t.Errorf("IsName| got %v, want %v", got, tt.want)
			}
		})
	}
}
