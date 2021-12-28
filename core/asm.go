package core

type OpCode struct {
	Pat     string
	Op      byte
	ArgSize byte
	Fn      func(memory []byte, ePtr, sPtr int) (ePtDelta, sPtDelta int)
}

var OpCodes = [0xff]OpCode{
	{ // Byte Add
		Pat: "+.",
		Op:  0x00,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			sPt += 1
			memory[sPt] = memory[sPt] + memory[sPt-1]
			return 1, 1
		},
	},
	{ // Int16 Add
		Pat: "+o",
		Op:  0x01,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			add := I16tob(Btoi16(memory[sPt:sPt+2]) + Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = add[0]
			memory[sPt+3] = add[1]
			return 1, 2
		},
	},
	{ // Int32 Add
		Pat: "+O",
		Op:  0x02,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			add := I32tob(Btoi32(memory[sPt:sPt+4]) + Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = add[0]
			memory[sPt+5] = add[1]
			memory[sPt+6] = add[2]
			memory[sPt+7] = add[3]
			return 1, 4
		},
	},
	{ // Int64 Add
		Pat: "+",
		Op:  0x03,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			add := I64tob(Btoi64(memory[sPt:sPt+8]) + Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = add[0]
			memory[sPt+9] = add[1]
			memory[sPt+10] = add[2]
			memory[sPt+11] = add[3]
			memory[sPt+12] = add[3]
			memory[sPt+13] = add[3]
			memory[sPt+14] = add[3]
			memory[sPt+15] = add[3]
			return 1, 8
		},
	},
	{ // Byte Minus
		Pat: "-.",
		Op:  0x04,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			sPt += 1
			memory[sPt] = memory[sPt-1] - memory[sPt]
			return 1, 1
		},
	},
	{ // Int16 Minus
		Pat: "-o",
		Op:  0x05,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			add := I16tob(Btoi16(memory[sPt:sPt+2]) - Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = add[0]
			memory[sPt+3] = add[1]
			return 1, 2
		},
	},
	{ // Int32 Minus
		Pat: "-O",
		Op:  0x06,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			add := I32tob(Btoi32(memory[sPt:sPt+4]) - Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = add[0]
			memory[sPt+5] = add[1]
			memory[sPt+6] = add[2]
			memory[sPt+7] = add[3]
			return 1, 4
		},
	},
	{ // Int64 Minus
		Pat: "-",
		Op:  0x07,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			add := I64tob(Btoi64(memory[sPt:sPt+8]) - Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = add[0]
			memory[sPt+9] = add[1]
			memory[sPt+10] = add[2]
			memory[sPt+11] = add[3]
			memory[sPt+12] = add[3]
			memory[sPt+13] = add[3]
			memory[sPt+14] = add[3]
			memory[sPt+15] = add[3]
			return 1, 8
		},
	},
	{ // Byte Multiply
		Pat: "*.",
		Op:  0x08,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			sPt += 1
			memory[sPt] = memory[sPt-1] * memory[sPt]
			return 1, 1
		},
	},
	{ // Int16 Multiply
		Pat: "*o",
		Op:  0x09,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			mul := I16tob(Btoi16(memory[sPt:sPt+2]) * Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = mul[0]
			memory[sPt+3] = mul[1]
			return 1, 2
		},
	},
	{ // Int32 Multiply
		Pat: "*O",
		Op:  0x0a,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			mul := I32tob(Btoi32(memory[sPt:sPt+4]) * Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = mul[0]
			memory[sPt+5] = mul[1]
			memory[sPt+6] = mul[2]
			memory[sPt+7] = mul[3]
			return 1, 4
		},
	},
	{ // Int64 Multiply
		Pat: "*",
		Op:  0x0b,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			mul := I64tob(Btoi64(memory[sPt:sPt+8]) * Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = mul[0]
			memory[sPt+9] = mul[1]
			memory[sPt+10] = mul[2]
			memory[sPt+11] = mul[3]
			memory[sPt+12] = mul[4]
			memory[sPt+13] = mul[5]
			memory[sPt+14] = mul[6]
			memory[sPt+15] = mul[7]
			return 1, 8
		},
	},
	{ // Byte divide
		Pat: "/.",
		Op:  0x0c,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			sPt += 1
			memory[sPt] = memory[sPt-1] / memory[sPt]
			return 1, 1
		},
	},
	{ // Int16 divide
		Pat: "/o",
		Op:  0x0d,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			div := I16tob(Btoi16(memory[sPt:sPt+2]) / Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = div[0]
			memory[sPt+3] = div[1]
			return 1, 2
		},
	},
	{ // Int32 divide
		Pat: "/O",
		Op:  0x0e,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			div := I32tob(Btoi32(memory[sPt:sPt+4]) / Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = div[0]
			memory[sPt+5] = div[1]
			memory[sPt+6] = div[2]
			memory[sPt+7] = div[3]
			return 1, 4
		},
	},
	{ // Int64 divide
		Pat: "/",
		Op:  0x0f,
		Fn: func(memory []byte, _, sPt int) (ePtDelta, sPtDelta int) {
			div := I64tob(Btoi64(memory[sPt:sPt+8]) / Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = div[0]
			memory[sPt+9] = div[1]
			memory[sPt+10] = div[2]
			memory[sPt+11] = div[3]
			memory[sPt+12] = div[3]
			memory[sPt+13] = div[3]
			memory[sPt+14] = div[3]
			memory[sPt+15] = div[3]
			return 1, 8
		},
	},
	{ // Byte Push
		Pat:     "<.",
		Op:      0x10,
		ArgSize: 8,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			memory[sPt-1] = memory[ePt+1]
			return 2, -1
		},
	},
	{ // Int16 Push
		Pat:     "<o",
		Op:      0x11,
		ArgSize: 16,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			memory[sPt-2] = memory[ePt+1]
			memory[sPt-1] = memory[ePt+2]
			return 3, -2
		},
	},
	{ // Int32 push
		Pat:     "<O",
		Op:      0x12,
		ArgSize: 32,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			memory[sPt-4] = memory[ePt+1]
			memory[sPt-3] = memory[ePt+2]
			memory[sPt-2] = memory[ePt+3]
			memory[sPt-1] = memory[ePt+4]
			return 5, -4
		},
	},
	{ // Int64 push
		Pat:     "<",
		Op:      0x13,
		ArgSize: 64,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			memory[sPt-8] = memory[ePt+1]
			memory[sPt-7] = memory[ePt+2]
			memory[sPt-6] = memory[ePt+3]
			memory[sPt-5] = memory[ePt+4]
			memory[sPt-4] = memory[ePt+5]
			memory[sPt-3] = memory[ePt+6]
			memory[sPt-2] = memory[ePt+7]
			memory[sPt-1] = memory[ePt+8]
			return 9, -8
		},
	},
	{ // Byte pop
		Pat: ">.",
		Op:  0x14,
		// TODO: I need to rethink this for serveral reasons,
		// this allows me to write to any spot in memory, even if it's marked as read only
		// also, would 32 bit pointers be enough here?
		// Also, I need to write a test for this function here
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			addr := Btoi64(memory[sPt : sPt+8])
			memory[addr] = memory[sPt+8]
			return 1, 9
		},
	},
	{ // Int16 pop
		Pat: ">o",
		Op:  0x15,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			addr := Btoi64(memory[sPt : sPt+8])
			memory[addr] = memory[sPt+8]
			memory[addr+1] = memory[sPt+9]
			return 1, 10
		},
	},
	{ // Int32 pop
		Pat: ">O",
		Op:  0x16,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			addr := Btoi64(memory[sPt : sPt+8])
			memory[addr] = memory[sPt+8]
			memory[addr+1] = memory[sPt+9]
			memory[addr+2] = memory[sPt+10]
			memory[addr+3] = memory[sPt+11]
			return 1, 12
		},
	},
	{ // Int64 pop
		Pat: ">",
		Op:  0x17,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			addr := Btoi64(memory[sPt : sPt+8])
			memory[addr] = memory[sPt+8]
			memory[addr+1] = memory[sPt+9]
			memory[addr+2] = memory[sPt+10]
			memory[addr+3] = memory[sPt+11]
			memory[addr+4] = memory[sPt+12]
			memory[addr+5] = memory[sPt+13]
			memory[addr+6] = memory[sPt+14]
			memory[addr+7] = memory[sPt+15]
			return 1, 16
		},
	},
	{ // Bitwise OR
		Pat: "|.",
		Op:  0x18,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			memory[sPt+1] = memory[sPt] | memory[sPt+1]
			return 1, 1
		},
	},
	{ // Int16 Bitwise Or
		Pat: "|o",
		Op:  0x19,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			or := I16tob(Btoi16(memory[sPt:sPt+2]) | Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = or[0]
			memory[sPt+3] = or[1]
			return 1, 2
		},
	},
	{ // Int32 Bitwise Or
		Pat: "|O",
		Op:  0x1a,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			or := I32tob(Btoi32(memory[sPt:sPt+4]) | Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = or[0]
			memory[sPt+5] = or[1]
			memory[sPt+6] = or[2]
			memory[sPt+7] = or[3]
			return 1, 4
		},
	},
	{ // Int64 Bitwise Or
		Pat: "|",
		Op:  0x1b,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta, sPtDelta int) {
			or := I64tob(Btoi64(memory[sPt:sPt+8]) | Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = or[0]
			memory[sPt+9] = or[1]
			memory[sPt+10] = or[2]
			memory[sPt+11] = or[3]
			memory[sPt+12] = or[4]
			memory[sPt+13] = or[5]
			memory[sPt+14] = or[6]
			memory[sPt+15] = or[7]
			return 1, 8
		},
	},
	{ // Bitwise And
		Pat: "&.",
		Op:  0x1c,
	},
	{ // Int16 Bitwise And
		Pat: "&o",
		Op:  0x1d,
	},
	{ // Int32 Bitwise Or
		Pat: "&O",
		Op:  0x1e,
	},
	{ // Int64 Bitwise Or
		Pat: "&",
		Op:  0x1f,
	},
	{ // Bitwise Xor
		Pat: "^.",
		Op:  0x20,
	},
	{ // Int16 Bitwise Xor
		Pat: "^o",
		Op:  0x21,
	},
	{ // Int32 Bitwise Xor
		Pat: "^O",
		Op:  0x22,
	},
	{ // Int64 Bitwise Xor
		Pat: "^",
		Op:  0x23,
	},
	{ // Bitwise LeftShift
		Pat: "<<.",
		Op:  0x24,
	},
	{ // Int16 Bitwise LeftShift
		Pat: "<<o",
		Op:  0x25,
	},
	{ // Int32 Bitwise LeftShift
		Pat: "<<O",
		Op:  0x26,
	},
	{ // Int64 Bitwise LeftShift
		Pat: "<<",
		Op:  0x27,
	},
	{ // Bitwise RightShift
		Pat: ">>.",
		Op:  0x28,
	},
	{ // Int16 jBitwise RightShift
		Pat: ">>o",
		Op:  0x29,
	},
	{ // Int32 Bitwise RightShift
		Pat: ">>O",
		Op:  0x2a,
	},
	{ // Bitwise RightShift
		Pat: ">>",
		Op:  0x2b,
	},
	{ // Jump If Greater
		Pat: "?>.",
		Op:  0x2c,
	},
	{ // Int16 Jump If Greater
		Pat: "?>o",
		Op:  0x2d,
	},
	{ // Int32 Jump If Greater
		Pat: "?>O",
		Op:  0x2e,
	},
	{ // Int64 Jump If Greater
		Pat: "?>",
		Op:  0x2f,
	},
	{ // Jump If Less
		Pat: "?<.",
		Op:  0x30,
	},
	{ // Int16 Jump If Less
		Pat: "?<o",
		Op:  0x31,
	},
	{ // Int32 Jump If Less
		Pat: "?<O",
		Op:  0x32,
	},
	{ // Int64 Jump If Less
		Pat: "?<",
		Op:  0x33,
	},
	{ // Jump
		Pat: "|>",
		Op:  0x34,
	},
	{ // byte mod
		Pat: "%.",
		Op:  0x35,
	},
	{ // Int16 mod
		Pat: "%o",
		Op:  0x36,
	},
	{ // Int32 mod
		Pat: "%O",
		Op:  0x37,
	},
	{ // Int64 mod
		Pat: "%",
		Op:  0x38,
	},
	{ // Push Byte 0
		Pat: "<0.",
		Op:  0x39,
	},
	{ // Push Int16 0
		Pat: "<0o",
		Op:  0x3a,
	},
	{ // Push Int32 0
		Pat: "<0O",
		Op:  0x3b,
	},
	{ // Push Int64 0
		Pat: "<0",
		Op:  0x3c,
	},
	{ // Dec Byte
		Pat: "--.",
		Op:  0x3d,
	},
	{ // Dec Int16
		Pat: "--o",
		Op:  0x3e,
	},
	{ // Dec Int32
		Pat: "--O",
		Op:  0x3f,
	},
	{ // Dec Int64
		Pat: "--",
		Op:  0x40,
	},
	{ // Inc Byte
		Pat: "++.",
		Op:  0x41,
	},
	{ // Inc Int16
		Pat: "++o",
		Op:  0x42,
	},
	{ // Inc Int32
		Pat: "++O",
		Op:  0x43,
	},
	{ // Inc Int64
		Pat: "++",
		Op:  0x44,
	},
	{ // Byte push (addr)
		Pat:     "<.#",
		Op:      0x45,
		ArgSize: 64,
	},
	{ // Int16 push (addr)
		Pat:     "<o#",
		Op:      0x46,
		ArgSize: 64,
	},
	{ // Int32 push (addr)
		Pat:     "<O#",
		Op:      0x47,
		ArgSize: 64,
	},
	{ // Int64 push (addr)
		Pat:     "<#",
		Op:      0x48,
		ArgSize: 64,
	},
}
