package core

import (
	"testing"
)

func TestI64tob(t *testing.T) {
	tests := []int{
		1029384,
		-50,
		0x7fffffffffffffff,
		-0x7fffffffffffffff,
		0,
	}

	for _, i := range tests {
		if c := Btoi64(I64tob(i)); c != i {
			t.Errorf("I64tob| failed to convert from into to binary and back, got: %d, want: %d", c, i)
		}
	}
}

func TestI32tob(t *testing.T) {
	tests := []int32{
		102984,
		-2837,
		0x7fffffff,
		-0x7fffffff,
		0,
	}

	for _, i := range tests {
		if c := Btoi32(I32tob(i)); c != i {
			t.Errorf("I32tob| failed to convert from into to binary and back, got: %d, want: %d", c, i)
		}
	}
}

func TestI16tob(t *testing.T) {
	tests := []int16{
		10293,
		-2837,
		0x7fff,
		-0x7fff,
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

func TestIsAddrCMD(t *testing.T) {
	tests := []struct {
		name string
		expr []string
		want bool
	}{
		{
			"push cmd",
			[]string{"<.", `\test`},
			false,
		},
		{
			"addr cmd",
			[]string{"<.", "#0x00"},
			true,
		},
		{
			"addr with label",
			[]string{"<.", `#\test`},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAddrCMD(tt.expr); got != tt.want {
				t.Errorf("IsAddrCMD| got %v but wanted %v", got, tt.want)
			}
		})
	}
}

func TestAsi64(t *testing.T) {
	tests := []int{
		10,
		-10,
		0x7fffffffffffffff,
		-0x7fffffffffffffff,
		0,
	}

	for _, tt := range tests {
		b := I64asb(&tt)
		if got := Asi64(&b[0]); *got != tt {
			t.Errorf("TestAsiI64| got %d but wanted %d", got, tt)
		}
	}
}

func TestAsi32(t *testing.T) {
	tests := []int32{
		10,
		-10,
		0x7fffffff,
		-0x7fffffff,
		0,
	}

	for _, tt := range tests {
		b := I32asb(&tt)
		if got := Asi32(&b[0]); *got != tt {
			t.Errorf("TestAsi32| got %d but wanted %d", got, tt)
		}
	}
}

func TestAsi16(t *testing.T) {
	tests := []int16{
		10,
		-10,
		0x7fff,
		-0x7fff,
		0,
	}

	for _, tt := range tests {
		b := I16asb(&tt)
		if got := Asi16(&b[0]); *got != tt {
			t.Errorf("TestAsi16| got %d but wanted %d", got, tt)
		}
	}
}

func TestAsu64(t *testing.T) {
	tests := []uint{
		10,
		0xffffffffffffffff,
		0,
	}

	for _, tt := range tests {
		b := U64asb(&tt)
		if got := Asu64(&b[0]); *got != tt {
			t.Errorf("TestAsu64| got %d but wanted %d", got, tt)
		}
	}
}

func TestAsu32(t *testing.T) {
	tests := []uint32{
		10,
		0xffffffff,
		0,
	}

	for _, tt := range tests {
		b := U32asb(&tt)
		if got := Asu32(&b[0]); *got != tt {
			t.Errorf("TestAsu32| got %d but wanted %d", got, tt)
		}
	}
}

func TestAsu16(t *testing.T) {
	tests := []uint16{
		10,
		0xffff,
		0,
	}

	for _, tt := range tests {
		b := U16asb(&tt)
		if got := Asu16(&b[0]); *got != tt {
			t.Errorf("TestAsu16| got %d but wanted %d", got, tt)
		}
	}
}
