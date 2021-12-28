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

type opCodeTest struct {
	opCode     byte
	name       string
	memory     []byte
	ePt        int
	sPt        int
	wantEPt    int
	wantSPt    int
	wantMemory []byte
}

func TestOpCodeFns(t *testing.T) {
	for _, tt := range opCodeFnTests {
		t.Run(tt.name, func(t *testing.T) {
			if OpCodes[tt.opCode].Pat == "" {
				t.Fatalf("%x| op code pattern is not set", tt.opCode)
			}
			if OpCodes[tt.opCode].Fn == nil {
				t.Fatalf("%s[0x%x]| op code function not implemented yet", OpCodes[tt.opCode].Pat, tt.opCode)
			}
			ePtDelta, sPtDelta := OpCodes[tt.opCode].Fn(tt.memory, tt.ePt, tt.sPt)
			if tt.ePt+ePtDelta != tt.wantEPt {
				t.Errorf("%s[0x%x]| ePt was wrong, got: %d, want: %d", OpCodes[tt.opCode].Pat, tt.opCode, tt.ePt+ePtDelta, tt.wantEPt)
			}
			if tt.sPt+sPtDelta != tt.wantSPt {
				t.Errorf("%s[0x%x]| sPt was wrong, got: %d, want: %d", OpCodes[tt.opCode].Pat, tt.opCode, tt.sPt+sPtDelta, tt.wantSPt)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("%s[0x%x]| memory change was wrong,\ngot:  %v,\nwant: %v", OpCodes[tt.opCode].Pat, tt.opCode, tt.memory, tt.wantMemory)
			}
		})
	}
}

// all the op code tests
var opCodeFnTests = []opCodeTest{
	{
		opCode:     0x00,
		name:       "byte add",
		memory:     []byte{2, 3, 0x00},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{2, 5, 0x00}, // 2 is not overwritten
	},
	{
		opCode:     0x00,
		name:       "overflow byte add",
		memory:     []byte{0xff, 5, 0x00},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{0xff, 4, 0x00}, // 0xff is not overwritten
	},
	{
		opCode:     0x01,
		name:       "int16 add",
		memory:     []byte{50, 0, 100, 0, 0x01},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{50, 0, 150, 0, 0x01}, // 50 is not overwritten
	},
	{
		opCode:     0x01,
		name:       "overflow int16 add",
		memory:     []byte{0xff, 0xff, 5, 0, 0x01},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{0xff, 0xff, 4, 0, 0x01}, // 0xffff is not overwritten
	},
	{
		opCode:     0x02,
		name:       "int32 add",
		memory:     []byte{50, 0, 0, 0, 100, 0, 0, 0, 0x02},
		ePt:        9,
		sPt:        0,
		wantEPt:    10,
		wantSPt:    4,
		wantMemory: []byte{50, 0, 0, 0, 150, 0, 0, 0, 0x02}, // 50 is not overwritten
	},
	{
		opCode:     0x02,
		name:       "overflow int32 add",
		memory:     []byte{0xff, 0xff, 0xff, 0xff, 5, 0, 0, 0, 0x02},
		ePt:        9,
		sPt:        0,
		wantEPt:    10,
		wantSPt:    4,
		wantMemory: []byte{0xff, 0xff, 0xff, 0xff, 4, 0, 0, 0, 0x02},
	},
	{
		opCode:     0x03,
		name:       "int64 add",
		memory:     []byte{50, 0, 0, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0x03},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{50, 0, 0, 0, 0, 0, 0, 0, 150, 0, 0, 0, 0, 0, 0, 0, 0x03}, // 50 is not overwritten
	},
	{
		opCode:     0x03,
		name:       "overflow int64 add",
		memory:     []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 5, 0, 0, 0, 0, 0, 0, 0, 0x03},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 4, 0, 0, 0, 0, 0, 0, 0, 0x03}, // 0xffffffffffffffff is not overwritten
	},
	{
		opCode:     0x04,
		name:       "subtract byte",
		memory:     []byte{3, 2, 0x04},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{3, 1, 0x04}, // 3 is not overwritten
	},
	{
		opCode:     0x04,
		name:       "underflow subtract byte",
		memory:     []byte{0, 9, 0x04},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{0, 0xf7, 0x04}, // 0 is not overwritten
	},
	{
		opCode:     0x05,
		name:       "subtract int16",
		memory:     []byte{50, 0, 20, 0, 0x05},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{50, 0, 30, 0, 0x05}, // 50 is not overwritten
	},
	{
		opCode:     0x05,
		name:       "underflow subtract int16",
		memory:     []byte{0, 0, 9, 0, 0x05},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{0, 0, 0xf7, 0xff, 0x05},
	},
	{
		opCode:     0x06,
		name:       "subtract int32",
		memory:     []byte{50, 0, 0, 0, 20, 0, 0, 0, 0x06},
		ePt:        8,
		sPt:        0,
		wantEPt:    9,
		wantSPt:    4,
		wantMemory: []byte{50, 0, 0, 0, 30, 0, 0, 0, 0x06}, // 50 is not overwritten
	},
	{
		opCode:     0x06,
		name:       "underflow subtract int32",
		memory:     []byte{0, 0, 0, 0, 9, 0, 0, 0, 0x06},
		ePt:        8,
		sPt:        0,
		wantEPt:    9,
		wantSPt:    4,
		wantMemory: []byte{0, 0, 0, 0, 0xf7, 0xff, 0xff, 0xff, 0x06},
	},
	{
		opCode:     0x07,
		name:       "subtract int64",
		memory:     []byte{50, 0, 0, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0x07},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{50, 0, 0, 0, 0, 0, 0, 0, 30, 0, 0, 0, 0, 0, 0, 0, 0x07}, // 50 is not overwritten
	},
	{
		opCode:     0x07,
		name:       "underflow subtract int64",
		memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0x07},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x07},
	},
	{
		opCode:     0x08,
		name:       "multiply byte",
		memory:     []byte{3, 2, 0x08},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{3, 6, 0x08}, // 3 is not overwritten
	},
	{
		opCode:     0x08,
		name:       "overflow multiply byte",
		memory:     []byte{2, 0xff, 0x08},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{2, 0xfe, 0x08}, // 2 is not overwritten
	},
	{
		opCode:     0x09,
		name:       "multiply int16",
		memory:     []byte{10, 0, 5, 0, 0x09},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{10, 0, 50, 0, 0x09}, // 10 is not overwritten
	},
	{
		opCode:     0x09,
		name:       "overflow multiply int16",
		memory:     []byte{0xff, 0xff, 2, 0, 0x09},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{0xff, 0xff, 0xfe, 0xff, 0x09}, // 0xffff is not overwritten
	},
	{
		opCode:     0x0a,
		name:       "multiply 32",
		memory:     []byte{10, 0, 0, 0, 5, 0, 0, 0, 0x0a},
		ePt:        8,
		sPt:        0,
		wantEPt:    9,
		wantSPt:    4,
		wantMemory: []byte{10, 0, 0, 0, 50, 0, 0, 0, 0x0a}, // 50 is not overwritten
	},
	{
		opCode:     0x0a,
		name:       "overflow multiply 32",
		memory:     []byte{0xff, 0xff, 0xff, 0xff, 2, 0, 0, 0, 0x0a},
		ePt:        8,
		sPt:        0,
		wantEPt:    9,
		wantSPt:    4,
		wantMemory: []byte{0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0x0a}, // 0xffffffff is not overwritten
	},
	{
		opCode:     0x0b,
		name:       "multiply int64",
		memory:     []byte{10, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0x0b},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{10, 0, 0, 0, 0, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 0x0b}, // 10 is not overwritten
	},
	{
		opCode:     0x0b,
		name:       "overflow multiply int64",
		memory:     []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 2, 0, 0, 0, 0, 0, 0, 0, 0x0b},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0b}, // 0xffffffffffffffff is not overwritten
	},
	{
		opCode:     0x0c,
		name:       "byte divide",
		memory:     []byte{10, 2, 0x0c},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{10, 5, 0x0c}, // 10 is not overwritten
	},
	{
		opCode:     0x0c,
		name:       "uneven byte divide",
		memory:     []byte{3, 2, 0x0c},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{3, 1, 0x0c}, // 3 is not overwritten
	},
	{
		opCode:     0x0d,
		name:       "int16 divide",
		memory:     []byte{100, 0, 5, 0, 0x0d},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{100, 0, 20, 0, 0x0d},
	},
	{
		opCode:     0x0d,
		name:       "uneven int16 divide",
		memory:     []byte{100, 0, 3, 0, 0x0d},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{100, 0, 33, 0, 0x0d}, // 100 is not overwritten
	},
	{
		opCode:     0x0e,
		name:       "int32 divide",
		memory:     []byte{100, 0, 0, 0, 5, 0, 0, 0, 0x0e},
		ePt:        8,
		sPt:        0,
		wantEPt:    9,
		wantSPt:    4,
		wantMemory: []byte{100, 0, 0, 0, 20, 0, 0, 0, 0x0e}, // 100 is not overwritten
	},
	{
		opCode:     0x0e,
		name:       "uneven int32 divide",
		memory:     []byte{100, 0, 0, 0, 3, 0, 0, 0, 0x0e},
		ePt:        8,
		sPt:        0,
		wantEPt:    9,
		wantSPt:    4,
		wantMemory: []byte{100, 0, 0, 0, 33, 0, 0, 0, 0x0e}, // 100 is not overwritten
	},
	{
		opCode:     0x0f,
		name:       "int64 divide",
		memory:     []byte{100, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0x0f},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{100, 0, 0, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0x0f}, // 100 is not overwritten
	},
	{
		opCode:     0x0f,
		name:       "uneven int64 divide",
		memory:     []byte{100, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0x0f},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{100, 0, 0, 0, 0, 0, 0, 0, 33, 0, 0, 0, 0, 0, 0, 0, 0x0f}, // 100 is not overwritten
	},
	{
		opCode:     0x10,
		name:       "byte push",
		memory:     []byte{0, 0x10, 5},
		ePt:        1,
		sPt:        1,
		wantEPt:    3,
		wantSPt:    0,
		wantMemory: []byte{5, 0x10, 5},
	},
	{
		opCode:     0x11,
		name:       "int16 push",
		memory:     []byte{0, 0, 0x11, 8, 7},
		ePt:        2,
		sPt:        2,
		wantEPt:    5,
		wantSPt:    0,
		wantMemory: []byte{8, 7, 0x11, 8, 7},
	},
	{
		opCode:     0x12,
		name:       "int32 push",
		memory:     []byte{0, 0, 0, 0, 0x12, 8, 7, 6, 5},
		ePt:        4,
		sPt:        4,
		wantEPt:    9,
		wantSPt:    0,
		wantMemory: []byte{8, 7, 6, 5, 0x12, 8, 7, 6, 5},
	},
	{
		opCode:     0x13,
		name:       "int64 push",
		memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0x13, 8, 7, 6, 5, 4, 3, 2, 1},
		ePt:        8,
		sPt:        8,
		wantEPt:    17,
		wantSPt:    0,
		wantMemory: []byte{8, 7, 6, 5, 4, 3, 2, 1, 0x13, 8, 7, 6, 5, 4, 3, 2, 1},
	},
	{
		opCode: 0x14,
		name:   "byte pop",
		//                 0  1  2  3  4  5  6  7  8  9  0  1  2
		memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 8, 0x14},
		ePt:        9,
		sPt:        0,
		wantEPt:    10,
		wantSPt:    9,
		wantMemory: []byte{8, 0, 0, 0, 0, 0, 0, 0, 8, 0x14},
	},
	{
		opCode: 0x15,
		name:   "int16 pop",
		//                 0  1  2  3  4  5  6  7  8  9  0  1  2
		memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0x15},
		ePt:        10,
		sPt:        0,
		wantEPt:    11,
		wantSPt:    10,
		wantMemory: []byte{8, 7, 0, 0, 0, 0, 0, 0, 8, 7, 0x15},
	},
	{
		opCode: 0x16,
		name:   "int32 pop",
		//                 0  1  2  3  4  5  6  7  8  9  0  1  2
		memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 6, 5, 0x16},
		ePt:        12,
		sPt:        0,
		wantEPt:    13,
		wantSPt:    12,
		wantMemory: []byte{8, 7, 6, 5, 0, 0, 0, 0, 8, 7, 6, 5, 0x16},
	},
	{
		opCode:     0x17,
		name:       "int64 pop",
		memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 6, 5, 4, 3, 2, 1, 0x17},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    16,
		wantMemory: []byte{8, 7, 6, 5, 4, 3, 2, 1, 8, 7, 6, 5, 4, 3, 2, 1, 0x17},
	},
	{
		opCode:     0x18,
		name:       "byte bitwise or",
		memory:     []byte{0xf0, 0x0f, 0x18},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{0xf0, 0xff, 0x18},
	},
	{
		opCode:     0x19,
		name:       "int16 bitwise or",
		memory:     []byte{0xf0, 0xf0, 0x0f, 0x0f, 0x19},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{0xf0, 0xf0, 0xff, 0xff, 0x19},
	},
	{
		opCode:     0x1a,
		name:       "int32 bitwise or",
		memory:     []byte{0xf0, 0xf0, 0xf0, 0xf0, 0x0f, 0x0f, 0x0f, 0x0f, 0x1a},
		ePt:        8,
		sPt:        0,
		wantEPt:    9,
		wantSPt:    4,
		wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xff, 0xff, 0xff, 0xff, 0x1a},
	},
	{
		opCode:     0x1b,
		name:       "int64 bitwise or",
		memory:     []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x1b},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1b},
	},
	{
		opCode:     0x1c,
		name:       "byte bitwise and",
		memory:     []byte{0xf0, 0xff, 0x1c},
		ePt:        2,
		sPt:        0,
		wantEPt:    3,
		wantSPt:    1,
		wantMemory: []byte{0xf0, 0xf0, 0x1c},
	},
	{
		opCode:     0x1d,
		name:       "int16 bitwise and",
		memory:     []byte{0xf0, 0xf0, 0xff, 0xff, 0x1d},
		ePt:        4,
		sPt:        0,
		wantEPt:    5,
		wantSPt:    2,
		wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0x1d},
	},
	{
		opCode:     0x1e,
		name:       "int32 bitwise and",
		memory:     []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xff, 0xff, 0xff, 0xff, 0x1e},
		ePt:        8,
		sPt:        0,
		wantEPt:    9,
		wantSPt:    4,
		wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0x1e},
	},
	{
		opCode:     0x1f,
		name:       "int64 bitwise and",
		memory:     []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1f},
		ePt:        16,
		sPt:        0,
		wantEPt:    17,
		wantSPt:    8,
		wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0x1f},
	},
}
