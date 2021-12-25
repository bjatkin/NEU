package core

import (
	"reflect"
	"testing"
)

func TestOpCodes(t *testing.T) {
	for i, op := range OpCodes {
		if i == 0 {
			continue
		}
		if op.Op != 0 && op.Op != byte(i) {
			t.Errorf("OpCode| op code '%s' is at the wrong index, got: %d, want: %d", op.Pat, op.Op, i)
		}
		if op.Op != 0 && op.Fn == nil {
			t.Errorf("OpCode| op code '%s' has no associated function", op.Pat)
		}
	}
	// check 0x00 explicitly so we don't error on empty byte codes
	if OpCodes[0].Pat != "+." {
		t.Errorf("OpCode| op code '%s' at index 0 should be '+.'", OpCodes[0].Pat)
	}
	if OpCodes[0].Op != 0 {
		t.Errorf("OpCode| op code '+.' is at the wrong index, got: %d, want: 0", OpCodes[0].Op)
	}
}

func TestByteAdd(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			[]byte{2, 3},
			0,
			1,
			1,
			[]byte{2, 5}, // 0x02 is not overwritten
		},
		{
			"overflow",
			[]byte{0xff, 5},
			0,
			1,
			1,
			[]byte{0xff, 4}, // 0xff is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x00].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("ByteAdd| rPtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("ByteAdd| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("ByteAdd| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt16Add(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I16tob(50), I16tob(100)...),
			0,
			1,
			2,
			append(I16tob(50), I16tob(150)...), // 50 is not overwritten
		},
		{
			"overflow",
			append(I16tob(0xffff), I16tob(5)...),
			0,
			1,
			2,
			append(I16tob(0xffff), I16tob(4)...), // 0xffff is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x01].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int16Add| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int16Add| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int16Add| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt32Add(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I32tob(50), I32tob(100)...),
			0,
			1,
			4,
			append(I32tob(50), I32tob(150)...), // 50 is not overwritten
		},
		{
			"overflow",
			append(I32tob(0xffffffff), I32tob(5)...),
			0,
			1,
			4,
			append(I32tob(0xffffffff), I32tob(4)...), // 0xffffffff is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x02].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int32Add| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int32Add| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int32Add| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt64Add(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I64tob(50), I64tob(100)...),
			0,
			1,
			8,
			append(I64tob(50), I64tob(150)...), // 50 is not overwritten
		},
		{
			"overflow",
			append(I64tob(0xffffffffffffffff), I64tob(5)...),
			0,
			1,
			8,
			append(I64tob(0xffffffffffffffff), I64tob(4)...), // 0xffffffffffffffff is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x03].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int64Add| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int64Add| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int64Add| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestByteMinus(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			[]byte{3, 2},
			0,
			1,
			1,
			[]byte{3, 1}, // 3 is not overwritten
		},
		{
			"underflow",
			[]byte{0, 9},
			0,
			1,
			1,
			[]byte{0, 0xf7}, // 0 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x04].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("ByteMinus| rPtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("ByteMinus| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("ByteMinus| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt16Minus(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I16tob(50), I16tob(20)...),
			0,
			1,
			2,
			append(I16tob(50), I16tob(30)...), // 50 is not overwritten
		},
		{
			"underflow",
			append(I16tob(0), I16tob(9)...),
			0,
			1,
			2,
			append(I16tob(0), I16tob(0xfff7)...), // 0 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x05].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int16Minus| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int16Minus| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int16Minus| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt32Minus(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I32tob(50), I32tob(20)...),
			0,
			1,
			4,
			append(I32tob(50), I32tob(30)...), // 50 is not overwritten
		},
		{
			"underflow",
			append(I32tob(0), I32tob(9)...),
			0,
			1,
			4,
			append(I32tob(0), I32tob(0xfffffff7)...), // 0 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x06].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int32Minus| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int32Minus| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int32Minus| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt64Minus(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I64tob(50), I64tob(20)...),
			0,
			1,
			8,
			append(I64tob(50), I64tob(30)...), // 50 is not overwritten
		},
		{
			"underflow",
			append(I64tob(0), I64tob(9)...),
			0,
			1,
			8,
			append(I64tob(0), I64tob(0xfffffffffffffff7)...), // 0 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x07].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int64Minus| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int64Minus| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int64Minus| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestByteMultiply(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			[]byte{3, 2},
			0,
			1,
			1,
			[]byte{3, 6}, // 3 is not overwritten
		},
		{
			"overflow",
			[]byte{2, 0xff},
			0,
			1,
			1,
			[]byte{2, 0xfe}, // 2 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x08].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("ByteMultiply| rPtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("ByteMultiply| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("ByteMultiply| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt16Multiply(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I16tob(10), I16tob(5)...),
			0,
			1,
			2,
			append(I16tob(10), I16tob(50)...), // 10 is not overwritten
		},
		{
			"overflow",
			append(I16tob(0xffff), I16tob(2)...),
			0,
			1,
			2,
			append(I16tob(0xffff), I16tob(0xfffe)...), // 0xffff is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x09].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int16Multiply| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int16Multiply| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int16Multiply| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt32Multiply(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I32tob(10), I32tob(5)...),
			0,
			1,
			4,
			append(I32tob(10), I32tob(50)...), // 50 is not overwritten
		},
		{
			"overflow",
			append(I32tob(0xffffffff), I32tob(2)...),
			0,
			1,
			4,
			append(I32tob(0xffffffff), I32tob(0xfffffffe)...), // 0xffffffff is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x0a].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int32Multiply| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int32Multiply| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int32Multiply| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt64Multiply(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I64tob(10), I64tob(5)...),
			0,
			1,
			8,
			append(I64tob(10), I64tob(50)...), // 10 is not overwritten
		},
		{
			"overflow",
			append(I64tob(0xffffffffffffffff), I64tob(2)...),
			0,
			1,
			8,
			append(I64tob(0xffffffffffffffff), I64tob(0xfffffffffffffffe)...), // 0xffffffffffffffff is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x0b].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int64Multiply| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int64Multiply| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int64Multiply| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestByteDivide(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			[]byte{10, 2},
			0,
			1,
			1,
			[]byte{10, 5}, // 3 is not overwritten
		},
		{
			"uneven",
			[]byte{3, 2},
			0,
			1,
			1,
			[]byte{3, 1}, // 3 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x0c].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("ByteDivide| rPtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("ByteDivide| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("ByteDivide| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt16Divide(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I16tob(100), I16tob(5)...),
			0,
			1,
			2,
			append(I16tob(100), I16tob(20)...), // 100 is not overwritten
		},
		{
			"uneven",
			append(I16tob(100), I16tob(3)...),
			0,
			1,
			2,
			append(I16tob(100), I16tob(33)...), // 100 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x0d].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int16Divide| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int16Divide| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int16Divide| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt32Divide(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I32tob(100), I32tob(5)...),
			0,
			1,
			4,
			append(I32tob(100), I32tob(20)...), // 100 is not overwritten
		},
		{
			"uneven",
			append(I32tob(100), I32tob(3)...),
			0,
			1,
			4,
			append(I32tob(100), I32tob(33)...), // 100 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x0e].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int32Divide| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int32Divide| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int32Divide| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestInt64Divide(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			append(I64tob(100), I64tob(5)...),
			0,
			1,
			8,
			append(I64tob(100), I64tob(20)...), // 100 is not overwritten
		},
		{
			"uneven",
			append(I64tob(100), I64tob(3)...),
			0,
			1,
			8,
			append(I64tob(100), I64tob(33)...), // 100 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x0f].Fn(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("Int64Minus| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("Int64Minus| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("Int64Minus| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestBytePush(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		ePt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			[]byte{0, 0x10, 5, 0},
			1,
			1,
			2,
			-1,
			[]byte{5, 0x10, 5, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x10].Fn(tt.memory, tt.ePt, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("BytePush| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("BytePush| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("BytePush| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}

func TestBytePop(t *testing.T) {
	tests := []struct {
		name         string
		memory       []byte
		sPt          int
		ePt          int
		wantEPtDelta int
		wantSPtDelta int
		wantMemory   []byte
	}{
		{
			"normal",
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0x14},
			1,
			10,
			1,
			9,
			[]byte{5, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0x14},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := OpCodes[0x14].Fn(tt.memory, tt.ePt, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("BytePop| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
			}
			if sPtDelta != tt.wantSPtDelta {
				t.Errorf("BytePop| sPtDelta was wrong, got: %d, want: %d", sPtDelta, tt.wantSPtDelta)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("BytePop| memory change was wrong, got: %v, want: %v", tt.memory, tt.wantMemory)
			}
		})
	}
}
