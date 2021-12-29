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
	name           string
	memory         []byte
	readOnlyOffset uint
	ePt            uint
	sPt            uint
	wantEPt        uint
	wantSPt        uint
	wantMemory     []byte
}

func TestOpCodeFns(t *testing.T) {
	for _, tt := range opCodeFnTests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.ePt) >= len(tt.memory) {
				t.Fatalf("execution pointer is out of range %d", tt.ePt)
			}

			opCode := tt.memory[tt.ePt+tt.readOnlyOffset]
			if OpCodes[opCode].Pat == "" {
				t.Fatalf("%x| op code pattern is not set", opCode)
			}
			if OpCodes[opCode].Fn == nil {
				t.Fatalf("%s[0x%x]| op code function not implemented yet", OpCodes[opCode].Pat, opCode)
			}
			newEPt, newSPt := OpCodes[opCode].Fn(tt.memory[:tt.readOnlyOffset], tt.memory[tt.readOnlyOffset:], tt.ePt, tt.sPt)
			if newEPt != tt.wantEPt {
				t.Errorf("%s[0x%x]| ePt was wrong, got: %d, want: %d", OpCodes[opCode].Pat, opCode, newEPt, tt.wantEPt)
			}
			if newSPt != tt.wantSPt {
				t.Errorf("%s[0x%x]| sPt was wrong, got: %d, want: %d", OpCodes[opCode].Pat, opCode, newSPt, tt.wantSPt)
			}
			if !reflect.DeepEqual(tt.memory, tt.wantMemory) {
				t.Errorf("%s[0x%x]| memory change was wrong,\ngot:  %v,\nwant: %v", OpCodes[opCode].Pat, opCode, tt.memory, tt.wantMemory)
			}
		})
	}
}

// all the op code tests
var opCodeFnTests = []opCodeTest{
	{
		name:           "byte add",
		memory:         []byte{2, 3, 0x00},
		readOnlyOffset: 2,
		ePt:            0,
		sPt:            0,
		wantEPt:        1,
		wantSPt:        1,
		wantMemory:     []byte{2, 5, 0x00}, // 2 is not overwritten
	},
	{
		name:           "overflow byte add",
		memory:         []byte{0xff, 5, 0x00},
		readOnlyOffset: 2,
		ePt:            0,
		sPt:            0,
		wantEPt:        1,
		wantSPt:        1,
		wantMemory:     []byte{0xff, 4, 0x00}, // 0xff is not overwritten
	},
	{
		name:           "int16 add",
		memory:         []byte{50, 0, 100, 0, 0x01},
		readOnlyOffset: 4,
		ePt:            0,
		sPt:            0,
		wantEPt:        1,
		wantSPt:        2,
		wantMemory:     []byte{50, 0, 150, 0, 0x01}, // 50 is not overwritten
	},
	// {
	// 	name:       "overflow int16 add",
	// 	memory:     []byte{0xff, 0xff, 5, 0, 0x01},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{0xff, 0xff, 4, 0, 0x01}, // 0xffff is not overwritten
	// },
	// {
	// 	name:       "int32 add",
	// 	memory:     []byte{50, 0, 0, 0, 100, 0, 0, 0, 0x02},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{50, 0, 0, 0, 150, 0, 0, 0, 0x02}, // 50 is not overwritten
	// },
	// {
	// 	name:       "overflow int32 add",
	// 	memory:     []byte{0xff, 0xff, 0xff, 0xff, 5, 0, 0, 0, 0x02},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{0xff, 0xff, 0xff, 0xff, 4, 0, 0, 0, 0x02},
	// },
	// {
	// 	name:       "int64 add",
	// 	memory:     []byte{50, 0, 0, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0x03},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{50, 0, 0, 0, 0, 0, 0, 0, 150, 0, 0, 0, 0, 0, 0, 0, 0x03}, // 50 is not overwritten
	// },
	// {
	// 	name:       "overflow int64 add",
	// 	memory:     []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 5, 0, 0, 0, 0, 0, 0, 0, 0x03},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 4, 0, 0, 0, 0, 0, 0, 0, 0x03}, // 0xffffffffffffffff is not overwritten
	// },
	// {
	// 	name:       "subtract byte",
	// 	memory:     []byte{3, 2, 0x04},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{3, 1, 0x04}, // 3 is not overwritten
	// },
	// {
	// 	name:       "underflow subtract byte",
	// 	memory:     []byte{0, 9, 0x04},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0, 0xf7, 0x04}, // 0 is not overwritten
	// },
	// {
	// 	name:       "subtract int16",
	// 	memory:     []byte{50, 0, 20, 0, 0x05},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{50, 0, 30, 0, 0x05}, // 50 is not overwritten
	// },
	// {
	// 	name:       "underflow subtract int16",
	// 	memory:     []byte{0, 0, 9, 0, 0x05},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{0, 0, 0xf7, 0xff, 0x05},
	// },
	// {
	// 	name:       "subtract int32",
	// 	memory:     []byte{50, 0, 0, 0, 20, 0, 0, 0, 0x06},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{50, 0, 0, 0, 30, 0, 0, 0, 0x06}, // 50 is not overwritten
	// },
	// {
	// 	name:       "underflow subtract int32",
	// 	memory:     []byte{0, 0, 0, 0, 9, 0, 0, 0, 0x06},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{0, 0, 0, 0, 0xf7, 0xff, 0xff, 0xff, 0x06},
	// },
	// {
	// 	name:       "subtract int64",
	// 	memory:     []byte{50, 0, 0, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0x07},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{50, 0, 0, 0, 0, 0, 0, 0, 30, 0, 0, 0, 0, 0, 0, 0, 0x07}, // 50 is not overwritten
	// },
	// {
	// 	name:       "underflow subtract int64",
	// 	memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0x07},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x07},
	// },
	// {
	// 	name:       "multiply byte",
	// 	memory:     []byte{3, 2, 0x08},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{3, 6, 0x08}, // 3 is not overwritten
	// },
	// {
	// 	name:       "overflow multiply byte",
	// 	memory:     []byte{2, 0xff, 0x08},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{2, 0xfe, 0x08}, // 2 is not overwritten
	// },
	// {
	// 	name:       "multiply int16",
	// 	memory:     []byte{10, 0, 5, 0, 0x09},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{10, 0, 50, 0, 0x09}, // 10 is not overwritten
	// },
	// {
	// 	name:       "overflow multiply int16",
	// 	memory:     []byte{0xff, 0xff, 2, 0, 0x09},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{0xff, 0xff, 0xfe, 0xff, 0x09}, // 0xffff is not overwritten
	// },
	// {
	// 	name:       "multiply 32",
	// 	memory:     []byte{10, 0, 0, 0, 5, 0, 0, 0, 0x0a},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{10, 0, 0, 0, 50, 0, 0, 0, 0x0a}, // 50 is not overwritten
	// },
	// {
	// 	name:       "overflow multiply 32",
	// 	memory:     []byte{0xff, 0xff, 0xff, 0xff, 2, 0, 0, 0, 0x0a},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0x0a}, // 0xffffffff is not overwritten
	// },
	// {
	// 	name:       "multiply int64",
	// 	memory:     []byte{10, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0x0b},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{10, 0, 0, 0, 0, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 0x0b}, // 10 is not overwritten
	// },
	// {
	// 	name:       "overflow multiply int64",
	// 	memory:     []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 2, 0, 0, 0, 0, 0, 0, 0, 0x0b},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0b}, // 0xffffffffffffffff is not overwritten
	// },
	// {
	// 	name:       "byte divide",
	// 	memory:     []byte{10, 2, 0x0c},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{10, 5, 0x0c}, // 10 is not overwritten
	// },
	// {
	// 	name:       "uneven byte divide",
	// 	memory:     []byte{3, 2, 0x0c},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{3, 1, 0x0c}, // 3 is not overwritten
	// },
	// {
	// 	name:       "int16 divide",
	// 	memory:     []byte{100, 0, 5, 0, 0x0d},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{100, 0, 20, 0, 0x0d},
	// },
	// {
	// 	name:       "uneven int16 divide",
	// 	memory:     []byte{100, 0, 3, 0, 0x0d},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{100, 0, 33, 0, 0x0d}, // 100 is not overwritten
	// },
	// {
	// 	name:       "int32 divide",
	// 	memory:     []byte{100, 0, 0, 0, 5, 0, 0, 0, 0x0e},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{100, 0, 0, 0, 20, 0, 0, 0, 0x0e}, // 100 is not overwritten
	// },
	// {
	// 	name:       "uneven int32 divide",
	// 	memory:     []byte{100, 0, 0, 0, 3, 0, 0, 0, 0x0e},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{100, 0, 0, 0, 33, 0, 0, 0, 0x0e}, // 100 is not overwritten
	// },
	// {
	// 	name:       "int64 divide",
	// 	memory:     []byte{100, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0x0f},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{100, 0, 0, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0x0f}, // 100 is not overwritten
	// },
	// {
	// 	name:       "uneven int64 divide",
	// 	memory:     []byte{100, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0x0f},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{100, 0, 0, 0, 0, 0, 0, 0, 33, 0, 0, 0, 0, 0, 0, 0, 0x0f}, // 100 is not overwritten
	// },
	// {
	// 	name:       "byte push",
	// 	memory:     []byte{0, 0x10, 5},
	// 	ePt:        1,
	// 	sPt:        1,
	// 	wantEPt:    3,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{5, 0x10, 5},
	// },
	// {
	// 	name:       "int16 push",
	// 	memory:     []byte{0, 0, 0x11, 8, 7},
	// 	ePt:        2,
	// 	sPt:        2,
	// 	wantEPt:    5,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{8, 7, 0x11, 8, 7},
	// },
	// {
	// 	name:       "int32 push",
	// 	memory:     []byte{0, 0, 0, 0, 0x12, 8, 7, 6, 5},
	// 	ePt:        4,
	// 	sPt:        4,
	// 	wantEPt:    9,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{8, 7, 6, 5, 0x12, 8, 7, 6, 5},
	// },
	// {
	// 	name:       "int64 push",
	// 	memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0x13, 8, 7, 6, 5, 4, 3, 2, 1},
	// 	ePt:        8,
	// 	sPt:        8,
	// 	wantEPt:    17,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{8, 7, 6, 5, 4, 3, 2, 1, 0x13, 8, 7, 6, 5, 4, 3, 2, 1},
	// },
	// {
	// 	name: "byte pop",
	// 	//                 0  1  2  3  4  5  6  7  8  9  0  1  2
	// 	memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 8, 0x14},
	// 	ePt:        9,
	// 	sPt:        0,
	// 	wantEPt:    10,
	// 	wantSPt:    9,
	// 	wantMemory: []byte{8, 0, 0, 0, 0, 0, 0, 0, 8, 0x14},
	// },
	// {
	// 	name: "int16 pop",
	// 	//                 0  1  2  3  4  5  6  7  8  9  0  1  2
	// 	memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0x15},
	// 	ePt:        10,
	// 	sPt:        0,
	// 	wantEPt:    11,
	// 	wantSPt:    10,
	// 	wantMemory: []byte{8, 7, 0, 0, 0, 0, 0, 0, 8, 7, 0x15},
	// },
	// {
	// 	name: "int32 pop",
	// 	//                 0  1  2  3  4  5  6  7  8  9  0  1  2
	// 	memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 6, 5, 0x16},
	// 	ePt:        12,
	// 	sPt:        0,
	// 	wantEPt:    13,
	// 	wantSPt:    12,
	// 	wantMemory: []byte{8, 7, 6, 5, 0, 0, 0, 0, 8, 7, 6, 5, 0x16},
	// },
	// {
	// 	name:       "int64 pop",
	// 	memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 6, 5, 4, 3, 2, 1, 0x17},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    16,
	// 	wantMemory: []byte{8, 7, 6, 5, 4, 3, 2, 1, 8, 7, 6, 5, 4, 3, 2, 1, 0x17},
	// },
	// {
	// 	name:       "byte bitwise or",
	// 	memory:     []byte{0xf0, 0x0f, 0x18},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0xf0, 0xff, 0x18},
	// },
	// {
	// 	name:       "int16 bitwise or",
	// 	memory:     []byte{0xf0, 0xf0, 0x0f, 0x0f, 0x19},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{0xf0, 0xf0, 0xff, 0xff, 0x19},
	// },
	// {
	// 	name:       "int32 bitwise or",
	// 	memory:     []byte{0xf0, 0xf0, 0xf0, 0xf0, 0x0f, 0x0f, 0x0f, 0x0f, 0x1a},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xff, 0xff, 0xff, 0xff, 0x1a},
	// },
	// {
	// 	name:       "int64 bitwise or",
	// 	memory:     []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x1b},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1b},
	// },
	// {
	// 	name:       "byte bitwise and",
	// 	memory:     []byte{0xf0, 0xff, 0x1c},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0xf0, 0xf0, 0x1c},
	// },
	// {
	// 	name:       "int16 bitwise and",
	// 	memory:     []byte{0xf0, 0xf0, 0xff, 0xff, 0x1d},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0x1d},
	// },
	// {
	// 	name:       "int32 bitwise and",
	// 	memory:     []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xff, 0xff, 0xff, 0xff, 0x1e},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0x1e},
	// },
	// {
	// 	name:       "int64 bitwise and",
	// 	memory:     []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1f},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0x1f},
	// },
	// {
	// 	name:       "byte bitwise xor",
	// 	memory:     []byte{0xf3, 0xcf, 0x20},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0xf3, 0x3c, 0x20},
	// },
	// {
	// 	name:       "int16 bitwise xor",
	// 	memory:     []byte{0xf3, 0xf3, 0xcf, 0xcf, 0x21},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{0xf3, 0xf3, 0x3c, 0x3c, 0x21},
	// },
	// {
	// 	name:       "int32 bitwise xor",
	// 	memory:     []byte{0xf3, 0xf3, 0xf3, 0xf3, 0xcf, 0xcf, 0xcf, 0xcf, 0x22},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{0xf3, 0xf3, 0xf3, 0xf3, 0x3c, 0x3c, 0x3c, 0x3c, 0x22},
	// },
	// {
	// 	name:       "int64 bitwise xor",
	// 	memory:     []byte{0xf3, 0xf3, 0xf3, 0xf3, 0xf3, 0xf3, 0xf3, 0xf3, 0xcf, 0xcf, 0xcf, 0xcf, 0xcf, 0xcf, 0xcf, 0xcf, 0x23},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{0xf3, 0xf3, 0xf3, 0xf3, 0xf3, 0xf3, 0xf3, 0xf3, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x3c, 0x23},
	// },
	// {
	// 	name:       "byte leftshift",
	// 	memory:     []byte{0x04, 0xfe, 0x24},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0x04, 0xe0, 0x24},
	// },
	// {
	// 	name:       "int16 leftshift",
	// 	memory:     []byte{0x04, 0xfe, 0xdc, 0x25},
	// 	ePt:        3,
	// 	sPt:        0,
	// 	wantEPt:    4,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0x04, 0xe0, 0xcf, 0x25},
	// },
	// {
	// 	name:       "int32 leftshift",
	// 	memory:     []byte{0x04, 0xfe, 0xdc, 0xba, 0x98, 0x26},
	// 	ePt:        5,
	// 	sPt:        0,
	// 	wantEPt:    6,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0x04, 0xe0, 0xcf, 0xad, 0x8b, 0x26},
	// },
	// {
	// 	name:       "int64 leftshift",
	// 	memory:     []byte{0x04, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0x27},
	// 	ePt:        9,
	// 	sPt:        0,
	// 	wantEPt:    10,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0x04, 0xe0, 0xcf, 0xad, 0x8b, 0x69, 0x47, 0x25, 0x03, 0x27},
	// },
	// {
	// 	name:       "byte rightshift",
	// 	memory:     []byte{0x04, 0xfe, 0x28},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0x04, 0x0f, 0x28},
	// },
	// {
	// 	name:       "int16 rightshift",
	// 	memory:     []byte{0x04, 0xfe, 0xdc, 0x29},
	// 	ePt:        3,
	// 	sPt:        0,
	// 	wantEPt:    4,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0x04, 0xcf, 0x0d, 0x29},
	// },
	// {
	// 	name:       "int32 righshift",
	// 	memory:     []byte{0x04, 0xfe, 0xdc, 0xba, 0x98, 0x2a},
	// 	ePt:        5,
	// 	sPt:        0,
	// 	wantEPt:    6,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0x04, 0xcf, 0xad, 0x8b, 0x09, 0x2a},
	// },
	// {
	// 	name:       "int64 rightshift",
	// 	memory:     []byte{0x04, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0x2b},
	// 	ePt:        9,
	// 	sPt:        0,
	// 	wantEPt:    10,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{0x04, 0xcf, 0xad, 0x8b, 0x69, 0x47, 0x25, 0x03, 0x01, 0x2b},
	// },
	// {
	// 	name:       "byte jump if greater",
	// 	memory:     []byte{10, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0x2c},
	// 	ePt:        10,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    10,
	// 	wantMemory: []byte{10, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0x2c},
	// },
	// {
	// 	name:       "byte jump if greater fail",
	// 	memory:     []byte{5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0x2c},
	// 	ePt:        10,
	// 	sPt:        0,
	// 	wantEPt:    11,
	// 	wantSPt:    10,
	// 	wantMemory: []byte{5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0x2c},
	// },
	// {
	// 	name:       "int16 jump if greater",
	// 	memory:     []byte{10, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2d},
	// 	ePt:        12,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    12,
	// 	wantMemory: []byte{10, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2d},
	// },
	// {
	// 	name:       "int16 jump if greater fail",
	// 	memory:     []byte{5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2d},
	// 	ePt:        12,
	// 	sPt:        0,
	// 	wantEPt:    13,
	// 	wantSPt:    12,
	// 	wantMemory: []byte{5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2d},
	// },
	// {
	// 	name:       "int32 jump if greater",
	// 	memory:     []byte{10, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2e},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    16,
	// 	wantMemory: []byte{10, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2e},
	// },
	// {
	// 	name:       "int32 jump if greater fail",
	// 	memory:     []byte{5, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2e},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    16,
	// 	wantMemory: []byte{5, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2e},
	// },
	// {
	// 	name:       "int64 jump if greater",
	// 	memory:     []byte{10, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2f},
	// 	ePt:        24,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    24,
	// 	wantMemory: []byte{10, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2f},
	// },
	// {
	// 	name:       "int64 jump if greater fail",
	// 	memory:     []byte{5, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2f},
	// 	ePt:        24,
	// 	sPt:        0,
	// 	wantEPt:    25,
	// 	wantSPt:    24,
	// 	wantMemory: []byte{5, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x2f},
	// },
	// ///////////////////////////////////
	// {
	// 	name:       "byte jump if less",
	// 	memory:     []byte{5, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0x30},
	// 	ePt:        10,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    10,
	// 	wantMemory: []byte{5, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0x30},
	// },
	// {
	// 	name:       "byte jump if less fail",
	// 	memory:     []byte{5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0x30},
	// 	ePt:        10,
	// 	sPt:        0,
	// 	wantEPt:    11,
	// 	wantSPt:    10,
	// 	wantMemory: []byte{5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0x30},
	// },
	// {
	// 	name:       "int16 jump if less",
	// 	memory:     []byte{5, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x31},
	// 	ePt:        12,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    12,
	// 	wantMemory: []byte{5, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x31},
	// },
	// {
	// 	name:       "int16 jump if less fail",
	// 	memory:     []byte{5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x31},
	// 	ePt:        12,
	// 	sPt:        0,
	// 	wantEPt:    13,
	// 	wantSPt:    12,
	// 	wantMemory: []byte{5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x31},
	// },
	// {
	// 	name:       "int32 jump if less",
	// 	memory:     []byte{5, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x32},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    16,
	// 	wantMemory: []byte{5, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x32},
	// },
	// {
	// 	name:       "int32 jump if less fail",
	// 	memory:     []byte{5, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x32},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    16,
	// 	wantMemory: []byte{5, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x32},
	// },
	// {
	// 	name:       "int64 jump if less",
	// 	memory:     []byte{5, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x33},
	// 	ePt:        24,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    24,
	// 	wantMemory: []byte{5, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x33},
	// },
	// {
	// 	name:       "int64 jump if less fail",
	// 	memory:     []byte{5, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x33},
	// 	ePt:        24,
	// 	sPt:        0,
	// 	wantEPt:    25,
	// 	wantSPt:    24,
	// 	wantMemory: []byte{5, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x33},
	// },
	// {
	// 	name:       "jump",
	// 	memory:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0x34},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    0,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0x34},
	// },
	// {
	// 	name:       "byte mod",
	// 	memory:     []byte{10, 2, 0x35},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    1,
	// 	wantMemory: []byte{10, 0, 0x35},
	// },
	// {
	// 	name:       "int16 mod",
	// 	memory:     []byte{10, 0, 2, 0, 0x36},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    2,
	// 	wantMemory: []byte{10, 0, 0, 0, 0x36},
	// },
	// {
	// 	name:       "int32 mod",
	// 	memory:     []byte{10, 0, 0, 0, 2, 0, 0, 0, 0x37},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    4,
	// 	wantMemory: []byte{10, 0, 0, 0, 0, 0, 0, 0, 0x37},
	// },
	// {
	// 	name:       "int64 mod",
	// 	memory:     []byte{10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0x38},
	// 	ePt:        16,
	// 	sPt:        0,
	// 	wantEPt:    17,
	// 	wantSPt:    8,
	// 	wantMemory: []byte{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x38},
	// },
	// {
	// 	name:       "byte push 0",
	// 	memory:     []byte{3, 0x39},
	// 	ePt:        1,
	// 	sPt:        1,
	// 	wantEPt:    2,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{0, 0x39},
	// },
	// {
	// 	name:       "int16 push 0",
	// 	memory:     []byte{3, 3, 0x3a},
	// 	ePt:        2,
	// 	sPt:        2,
	// 	wantEPt:    3,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{0, 0, 0x3a},
	// },
	// {
	// 	name:       "int32 push 0",
	// 	memory:     []byte{3, 3, 3, 3, 0x3b},
	// 	ePt:        4,
	// 	sPt:        4,
	// 	wantEPt:    5,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{0, 0, 0, 0, 0x3b},
	// },
	// {
	// 	name:       "int64 push 0",
	// 	memory:     []byte{3, 3, 3, 3, 3, 3, 3, 3, 0x3c},
	// 	ePt:        8,
	// 	sPt:        8,
	// 	wantEPt:    9,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0x3c},
	// },
	// {
	// 	name:       "dec byte",
	// 	memory:     []byte{10, 0x3d},
	// 	ePt:        1,
	// 	sPt:        0,
	// 	wantEPt:    2,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{9, 0x3d},
	// },
	// {
	// 	name:       "dec int16",
	// 	memory:     []byte{10, 0, 0x3e},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{9, 0, 0x3e},
	// },
	// {
	// 	name:       "dec int32",
	// 	memory:     []byte{10, 0, 0, 0, 0x3f},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{9, 0, 0, 0, 0x3f},
	// },
	// {
	// 	name:       "dec int64",
	// 	memory:     []byte{10, 0, 0, 0, 0, 0, 0, 0, 0x40},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{9, 0, 0, 0, 0, 0, 0, 0, 0x40},
	// },
	// {
	// 	name:       "inc byte",
	// 	memory:     []byte{10, 0x41},
	// 	ePt:        1,
	// 	sPt:        0,
	// 	wantEPt:    2,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{11, 0x41},
	// },
	// {
	// 	name:       "inc int16",
	// 	memory:     []byte{10, 0, 0x42},
	// 	ePt:        2,
	// 	sPt:        0,
	// 	wantEPt:    3,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{11, 0, 0x42},
	// },
	// {
	// 	name:       "inc int32",
	// 	memory:     []byte{10, 0, 0, 0, 0x43},
	// 	ePt:        4,
	// 	sPt:        0,
	// 	wantEPt:    5,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{11, 0, 0, 0, 0x43},
	// },
	// {
	// 	name:       "inc int64",
	// 	memory:     []byte{10, 0, 0, 0, 0, 0, 0, 0, 0x44},
	// 	ePt:        8,
	// 	sPt:        0,
	// 	wantEPt:    9,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{11, 0, 0, 0, 0, 0, 0, 0, 0x44},
	// },
	// {
	// 	name:       "byte push addr",
	// 	memory:     []byte{5, 0x45, 1, 0, 0, 0, 0, 0, 0, 0},
	// 	ePt:        1,
	// 	sPt:        1,
	// 	wantEPt:    10,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{0x45, 0x45, 1, 0, 0, 0, 0, 0, 0, 0},
	// },
	// {
	// 	name:       "int16 push addr",
	// 	memory:     []byte{5, 0, 0x46, 2, 0, 0, 0, 0, 0, 0, 0},
	// 	ePt:        2,
	// 	sPt:        2,
	// 	wantEPt:    11,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{0x46, 2, 0x46, 2, 0, 0, 0, 0, 0, 0, 0},
	// },
	// {
	// 	name:       "int32 push addr",
	// 	memory:     []byte{5, 0, 0, 0, 0x47, 4, 0, 0, 0, 0, 0, 0, 0},
	// 	ePt:        4,
	// 	sPt:        4,
	// 	wantEPt:    13,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{0x47, 4, 0, 0, 0x47, 4, 0, 0, 0, 0, 0, 0, 0},
	// },
	// {
	// 	name:       "int64 push addr",
	// 	memory:     []byte{5, 0, 0, 0, 0, 0, 0, 0, 0x48, 8, 0, 0, 0, 0, 0, 0, 0},
	// 	ePt:        8,
	// 	sPt:        8,
	// 	wantEPt:    17,
	// 	wantSPt:    0,
	// 	wantMemory: []byte{0x48, 8, 0, 0, 0, 0, 0, 0, 0x48, 8, 0, 0, 0, 0, 0, 0, 0},
	// },
}
