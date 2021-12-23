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
			[]byte{0x02, 0x03},
			0,
			1,
			1,
			[]byte{0x02, 0x05}, // 0x02 is not overwritten
		},
		{
			"overflow",
			[]byte{0xff, 0x05},
			0,
			1,
			1,
			[]byte{0xff, 0x04}, // 0xff is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := ByteAdd(tt.memory, 0, tt.sPt)
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
			ePtDelta, sPtDelta := Int16Add(tt.memory, 0, tt.sPt)
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
			ePtDelta, sPtDelta := Int32Add(tt.memory, 0, tt.sPt)
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
			ePtDelta, sPtDelta := Int64Add(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("IntAdd| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
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
			[]byte{0x03, 0x02},
			0,
			1,
			1,
			[]byte{0x03, 0x01}, // 0x03 is not overwritten
		},
		{
			"underflow",
			[]byte{0x00, 0x09},
			0,
			1,
			1,
			[]byte{0x00, 0xf7}, // 0x00 is not overwritten
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ePtDelta, sPtDelta := ByteMinus(tt.memory, 0, tt.sPt)
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
			ePtDelta, sPtDelta := Int16Minus(tt.memory, 0, tt.sPt)
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
			ePtDelta, sPtDelta := Int32Minus(tt.memory, 0, tt.sPt)
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
			ePtDelta, sPtDelta := Int64Minus(tt.memory, 0, tt.sPt)
			if ePtDelta != tt.wantEPtDelta {
				t.Errorf("IntMinus| ePtDelta was wrong, got: %d, want: %d", ePtDelta, tt.wantEPtDelta)
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
