package core

type OpCode struct {
	Pat     string
	Op      byte
	ArgSize byte
	Fn      func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint)
}

var OpCodes = [0xff]OpCode{
	{ // Byte Add
		Pat: "+.",
		Op:  0x00,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt] + memory[sPt+1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 Add
		Pat: "+o",
		Op:  0x01,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			add := I16tob(Btoi16(memory[sPt:sPt+2]) + Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = add[0]
			memory[sPt+3] = add[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Add
		Pat: "+O",
		Op:  0x02,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			add := I32tob(Btoi32(memory[sPt:sPt+4]) + Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = add[0]
			memory[sPt+5] = add[1]
			memory[sPt+6] = add[2]
			memory[sPt+7] = add[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Add
		Pat: "+",
		Op:  0x03,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			add := I64tob(Btoi64(memory[sPt:sPt+8]) + Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = add[0]
			memory[sPt+9] = add[1]
			memory[sPt+10] = add[2]
			memory[sPt+11] = add[3]
			memory[sPt+12] = add[3]
			memory[sPt+13] = add[3]
			memory[sPt+14] = add[3]
			memory[sPt+15] = add[3]
			return sPt + 8, ePt + 1
		},
	},
	{ // Byte Minus
		Pat: "-.",
		Op:  0x04,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt] - memory[sPt+1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 Minus
		Pat: "-o",
		Op:  0x05,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			add := I16tob(Btoi16(memory[sPt:sPt+2]) - Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = add[0]
			memory[sPt+3] = add[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Minus
		Pat: "-O",
		Op:  0x06,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			add := I32tob(Btoi32(memory[sPt:sPt+4]) - Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = add[0]
			memory[sPt+5] = add[1]
			memory[sPt+6] = add[2]
			memory[sPt+7] = add[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Minus
		Pat: "-",
		Op:  0x07,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			add := I64tob(Btoi64(memory[sPt:sPt+8]) - Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = add[0]
			memory[sPt+9] = add[1]
			memory[sPt+10] = add[2]
			memory[sPt+11] = add[3]
			memory[sPt+12] = add[3]
			memory[sPt+13] = add[3]
			memory[sPt+14] = add[3]
			memory[sPt+15] = add[3]
			return sPt + 8, ePt + 1
		},
	},
	{ // Byte Multiply
		Pat: "*.",
		Op:  0x08,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt+1] * memory[sPt]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 Multiply
		Pat: "*o",
		Op:  0x09,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mul := I16tob(Btoi16(memory[sPt:sPt+2]) * Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = mul[0]
			memory[sPt+3] = mul[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Multiply
		Pat: "*O",
		Op:  0x0a,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mul := I32tob(Btoi32(memory[sPt:sPt+4]) * Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = mul[0]
			memory[sPt+5] = mul[1]
			memory[sPt+6] = mul[2]
			memory[sPt+7] = mul[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Multiply
		Pat: "*",
		Op:  0x0b,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mul := I64tob(Btoi64(memory[sPt:sPt+8]) * Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = mul[0]
			memory[sPt+9] = mul[1]
			memory[sPt+10] = mul[2]
			memory[sPt+11] = mul[3]
			memory[sPt+12] = mul[4]
			memory[sPt+13] = mul[5]
			memory[sPt+14] = mul[6]
			memory[sPt+15] = mul[7]
			return sPt + 8, ePt + 1
		},
	},
	{ // Byte divide
		Pat: "/.",
		Op:  0x0c,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt] / memory[sPt+1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 divide
		Pat: "/o",
		Op:  0x0d,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			div := I16tob(Btoi16(memory[sPt:sPt+2]) / Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = div[0]
			memory[sPt+3] = div[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 divide
		Pat: "/O",
		Op:  0x0e,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			div := I32tob(Btoi32(memory[sPt:sPt+4]) / Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = div[0]
			memory[sPt+5] = div[1]
			memory[sPt+6] = div[2]
			memory[sPt+7] = div[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 divide
		Pat: "/",
		Op:  0x0f,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			div := I64tob(Btoi64(memory[sPt:sPt+8]) / Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = div[0]
			memory[sPt+9] = div[1]
			memory[sPt+10] = div[2]
			memory[sPt+11] = div[3]
			memory[sPt+12] = div[3]
			memory[sPt+13] = div[3]
			memory[sPt+14] = div[3]
			memory[sPt+15] = div[3]
			return sPt + 8, ePt + 1
		},
	},
	{ // Byte Push
		Pat:     "<.",
		Op:      0x10,
		ArgSize: 8,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt-1] = code[ePt+1]
			return sPt - 1, ePt + 2
		},
	},
	{ // Int16 Push
		Pat:     "<o",
		Op:      0x11,
		ArgSize: 16,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt-2] = code[ePt+1]
			memory[sPt-1] = code[ePt+2]
			return sPt - 2, ePt + 3
		},
	},
	{ // Int32 push
		Pat:     "<O",
		Op:      0x12,
		ArgSize: 32,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt-4] = code[ePt+1]
			memory[sPt-3] = code[ePt+2]
			memory[sPt-2] = code[ePt+3]
			memory[sPt-1] = code[ePt+4]
			return sPt - 4, ePt + 5
		},
	},
	{ // Int64 push
		Pat:     "<",
		Op:      0x13,
		ArgSize: 64,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt-8] = code[ePt+1]
			memory[sPt-7] = code[ePt+2]
			memory[sPt-6] = code[ePt+3]
			memory[sPt-5] = code[ePt+4]
			memory[sPt-4] = code[ePt+5]
			memory[sPt-3] = code[ePt+6]
			memory[sPt-2] = code[ePt+7]
			memory[sPt-1] = code[ePt+8]
			return sPt - 8, ePt + 9
		},
	},
	{ // Byte pop
		Pat: ">.",
		Op:  0x14,
		// TODO: would 32 bit pointers be enough (probably)?
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := Btoi64(memory[sPt : sPt+8])
			memory[addr] = memory[sPt+8]
			return sPt + 9, ePt + 1
		},
	},
	{ // Int16 pop
		Pat: ">o",
		Op:  0x15,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := Btoi64(memory[sPt : sPt+8])
			memory[addr] = memory[sPt+8]
			memory[addr+1] = memory[sPt+9]
			return sPt + 10, ePt + 1
		},
	},
	{ // Int32 pop
		Pat: ">O",
		Op:  0x16,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := Btoi64(memory[sPt : sPt+8])
			memory[addr] = memory[sPt+8]
			memory[addr+1] = memory[sPt+9]
			memory[addr+2] = memory[sPt+10]
			memory[addr+3] = memory[sPt+11]
			return sPt + 12, ePt + 1
		},
	},
	{ // Int64 pop
		Pat: ">",
		Op:  0x17,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := Btoi64(memory[sPt : sPt+8])
			memory[addr] = memory[sPt+8]
			memory[addr+1] = memory[sPt+9]
			memory[addr+2] = memory[sPt+10]
			memory[addr+3] = memory[sPt+11]
			memory[addr+4] = memory[sPt+12]
			memory[addr+5] = memory[sPt+13]
			memory[addr+6] = memory[sPt+14]
			memory[addr+7] = memory[sPt+15]
			return sPt + 16, ePt + 1
		},
	},
	{ // Bitwise OR
		Pat: "|.",
		Op:  0x18,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt] | memory[sPt+1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 Bitwise Or
		Pat: "|o",
		Op:  0x19,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I16tob(Btoi16(memory[sPt:sPt+2]) | Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = or[0]
			memory[sPt+3] = or[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Bitwise Or
		Pat: "|O",
		Op:  0x1a,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I32tob(Btoi32(memory[sPt:sPt+4]) | Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = or[0]
			memory[sPt+5] = or[1]
			memory[sPt+6] = or[2]
			memory[sPt+7] = or[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Bitwise Or
		Pat: "|",
		Op:  0x1b,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I64tob(Btoi64(memory[sPt:sPt+8]) | Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = or[0]
			memory[sPt+9] = or[1]
			memory[sPt+10] = or[2]
			memory[sPt+11] = or[3]
			memory[sPt+12] = or[4]
			memory[sPt+13] = or[5]
			memory[sPt+14] = or[6]
			memory[sPt+15] = or[7]
			return sPt + 8, ePt + 1
		},
	},
	{ // Bitwise And
		Pat: "&.",
		Op:  0x1c,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt] & memory[sPt+1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 Bitwise And
		Pat: "&o",
		Op:  0x1d,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I16tob(Btoi16(memory[sPt:sPt+2]) & Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = or[0]
			memory[sPt+3] = or[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Bitwise And
		Pat: "&O",
		Op:  0x1e,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I32tob(Btoi32(memory[sPt:sPt+4]) & Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = or[0]
			memory[sPt+5] = or[1]
			memory[sPt+6] = or[2]
			memory[sPt+7] = or[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Bitwise And
		Pat: "&",
		Op:  0x1f,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I64tob(Btoi64(memory[sPt:sPt+8]) & Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = or[0]
			memory[sPt+9] = or[1]
			memory[sPt+10] = or[2]
			memory[sPt+11] = or[3]
			memory[sPt+12] = or[4]
			memory[sPt+13] = or[5]
			memory[sPt+14] = or[6]
			memory[sPt+15] = or[7]
			return sPt + 8, ePt + 1
		},
	},
	{ // Bitwise Xor
		Pat: "^.",
		Op:  0x20,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt] ^ memory[sPt+1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 Bitwise Xor
		Pat: "^o",
		Op:  0x21,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I16tob(Btoi16(memory[sPt:sPt+2]) ^ Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = or[0]
			memory[sPt+3] = or[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Bitwise Xor
		Pat: "^O",
		Op:  0x22,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I32tob(Btoi32(memory[sPt:sPt+4]) ^ Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = or[0]
			memory[sPt+5] = or[1]
			memory[sPt+6] = or[2]
			memory[sPt+7] = or[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Bitwise Xor
		Pat: "^",
		Op:  0x23,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := I64tob(Btoi64(memory[sPt:sPt+8]) ^ Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = or[0]
			memory[sPt+9] = or[1]
			memory[sPt+10] = or[2]
			memory[sPt+11] = or[3]
			memory[sPt+12] = or[4]
			memory[sPt+13] = or[5]
			memory[sPt+14] = or[6]
			memory[sPt+15] = or[7]
			return sPt + 8, ePt + 1
		},
	},
	{ // Bitwise LeftShift
		Pat: "<<.",
		Op:  0x24,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt+1] << memory[sPt]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 Bitwise LeftShift
		Pat: "<<o",
		Op:  0x25,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := I16tob(Btoi16(memory[sPt+1:sPt+3]) << memory[sPt])
			memory[sPt+1] = shift[0]
			memory[sPt+2] = shift[1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int32 Bitwise LeftShift
		Pat: "<<O",
		Op:  0x26,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := I32tob(Btoi32(memory[sPt+1:sPt+5]) << memory[sPt])
			memory[sPt+1] = shift[0]
			memory[sPt+2] = shift[1]
			memory[sPt+3] = shift[2]
			memory[sPt+4] = shift[3]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int64 Bitwise LeftShift
		Pat: "<<",
		Op:  0x27,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := I64tob(Btoi64(memory[sPt+1:sPt+9]) << memory[sPt])
			memory[sPt+1] = shift[0]
			memory[sPt+2] = shift[1]
			memory[sPt+3] = shift[2]
			memory[sPt+4] = shift[3]
			memory[sPt+5] = shift[4]
			memory[sPt+6] = shift[5]
			memory[sPt+7] = shift[6]
			memory[sPt+8] = shift[7]
			return sPt + 1, ePt + 1
		},
	},
	{ // Bitwise RightShift
		Pat: ">>.",
		Op:  0x28,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt+1] >> memory[sPt]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 Bitwise RightShift
		Pat: ">>o",
		Op:  0x29,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := I16tob(Btoi16(memory[sPt+1:sPt+3]) >> memory[sPt])
			memory[sPt+1] = shift[0]
			memory[sPt+2] = shift[1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int32 Bitwise RightShift
		Pat: ">>O",
		Op:  0x2a,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := I32tob(Btoi32(memory[sPt+1:sPt+5]) >> memory[sPt])
			memory[sPt+1] = shift[0]
			memory[sPt+2] = shift[1]
			memory[sPt+3] = shift[2]
			memory[sPt+4] = shift[3]
			return sPt + 1, ePt + 1
		},
	},
	{ // Bitwise RightShift
		Pat: ">>",
		Op:  0x2b,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := I64tob(Btoi64(memory[sPt+1:sPt+9]) >> memory[sPt])
			memory[sPt+1] = shift[0]
			memory[sPt+2] = shift[1]
			memory[sPt+3] = shift[2]
			memory[sPt+4] = shift[3]
			memory[sPt+5] = shift[4]
			memory[sPt+6] = shift[5]
			memory[sPt+7] = shift[6]
			memory[sPt+8] = shift[7]
			return sPt + 1, ePt + 1
		},
	},
	{ // Jump If Greater
		Pat: "?>.",
		Op:  0x2c,
		// TODO: we should probably change ePt and sPt and their deltas to be uint's
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if memory[sPt] > memory[sPt+1] {
				return sPt + 10, Btoi64(memory[sPt+2 : sPt+10])
			}
			return sPt + 10, ePt + 1
		},
	},
	{ // Int16 Jump If Greater
		Pat: "?>o",
		Op:  0x2d,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if Btoi16(memory[sPt:sPt+2]) > Btoi16(memory[sPt+2:sPt+4]) {
				return sPt + 12, Btoi64(memory[sPt+4 : sPt+12])
			}
			return sPt + 12, ePt + 1
		},
	},
	{ // Int32 Jump If Greater
		Pat: "?>O",
		Op:  0x2e,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if Btoi32(memory[sPt:sPt+4]) > Btoi32(memory[sPt+4:sPt+8]) {
				return sPt + 16, Btoi64(memory[sPt+8 : sPt+16])
			}
			return sPt + 16, ePt + 1
		},
	},
	{ // Int64 Jump If Greater
		Pat: "?>",
		Op:  0x2f,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if Btoi64(memory[sPt:sPt+8]) > Btoi64(memory[sPt+8:sPt+16]) {
				return sPt + 24, Btoi64(memory[sPt+16 : sPt+24])
			}
			return sPt + 24, ePt + 1
		},
	},
	{ // Jump If Less
		Pat: "?<.",
		Op:  0x30,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if memory[sPt] < memory[sPt+1] {
				return sPt + 10, Btoi64(memory[sPt+2 : sPt+10])
			}
			return sPt + 10, ePt + 1
		},
	},
	{ // Int16 Jump If Less
		Pat: "?<o",
		Op:  0x31,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if Btoi16(memory[sPt:sPt+2]) < Btoi16(memory[sPt+2:sPt+4]) {
				return sPt + 12, Btoi64(memory[sPt+4 : sPt+12])
			}
			return sPt + 12, ePt + 1
		},
	},
	{ // Int32 Jump If Less
		Pat: "?<O",
		Op:  0x32,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if Btoi32(memory[sPt:sPt+4]) < Btoi32(memory[sPt+4:sPt+8]) {
				return sPt + 16, Btoi64(memory[sPt+8 : sPt+16])
			}
			return sPt + 16, ePt + 1
		},
	},
	{ // Int64 Jump If Less
		Pat: "?<",
		Op:  0x33,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if Btoi64(memory[sPt:sPt+8]) < Btoi64(memory[sPt+8:sPt+16]) {
				return sPt + 24, Btoi64(memory[sPt+16 : sPt+24])
			}
			return sPt + 24, ePt + 1
		},
	},
	{ // Jump
		Pat: "|>",
		Op:  0x34,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			return sPt + 8, Btoi64(memory[sPt : sPt+8])
		},
	},
	{ // byte mod
		Pat: "%.",
		Op:  0x35,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt+1] = memory[sPt] % memory[sPt+1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int16 mod
		Pat: "%o",
		Op:  0x36,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mod := I16tob(Btoi16(memory[sPt:sPt+2]) % Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = mod[0]
			memory[sPt+3] = mod[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 mod
		Pat: "%O",
		Op:  0x37,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mod := I32tob(Btoi32(memory[sPt:sPt+4]) % Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = mod[0]
			memory[sPt+5] = mod[1]
			memory[sPt+6] = mod[2]
			memory[sPt+7] = mod[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 mod
		Pat: "%",
		Op:  0x38,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mod := I64tob(Btoi64(memory[sPt:sPt+8]) % Btoi64(memory[sPt+8:sPt+16]))
			memory[sPt+8] = mod[0]
			memory[sPt+9] = mod[1]
			memory[sPt+10] = mod[2]
			memory[sPt+11] = mod[3]
			memory[sPt+12] = mod[4]
			memory[sPt+13] = mod[5]
			memory[sPt+14] = mod[6]
			memory[sPt+15] = mod[7]
			return sPt + 8, ePt + 1
		},
	},
	{ // Push Byte 0
		Pat: "<0.",
		Op:  0x39,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt-1] = 0
			return sPt - 1, ePt + 1
		},
	},
	{ // Push Int16 0
		Pat: "<0o",
		Op:  0x3a,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt-2] = 0
			memory[sPt-1] = 0
			return sPt - 2, ePt + 1
		},
	},
	{ // Push Int32 0
		Pat: "<0O",
		Op:  0x3b,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt-4] = 0
			memory[sPt-3] = 0
			memory[sPt-2] = 0
			memory[sPt-1] = 0
			return sPt - 4, ePt + 1
		},
	},
	{ // Push Int64 0
		Pat: "<0",
		Op:  0x3c,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt-8] = 0
			memory[sPt-7] = 0
			memory[sPt-6] = 0
			memory[sPt-5] = 0
			memory[sPt-4] = 0
			memory[sPt-3] = 0
			memory[sPt-2] = 0
			memory[sPt-1] = 0
			return sPt - 8, ePt + 1
		},
	},
	{ // Dec Byte
		Pat: "--.",
		Op:  0x3d,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt]--
			return sPt, ePt + 1
		},
	},
	{ // Dec Int16
		Pat: "--o",
		Op:  0x3e,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			dec := I16tob(Btoi16(memory[sPt:sPt+2]) - 1)
			memory[sPt] = dec[0]
			memory[sPt+1] = dec[1]
			return sPt, ePt + 1
		},
	},
	{ // Dec Int32
		Pat: "--O",
		Op:  0x3f,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			dec := I32tob(Btoi32(memory[sPt:sPt+4]) - 1)
			memory[sPt] = dec[0]
			memory[sPt+1] = dec[1]
			memory[sPt+2] = dec[2]
			memory[sPt+3] = dec[3]
			return sPt, ePt + 1
		},
	},
	{ // Dec Int64
		Pat: "--",
		Op:  0x40,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			dec := I64tob(Btoi64(memory[sPt:sPt+8]) - 1)
			memory[sPt] = dec[0]
			memory[sPt+1] = dec[1]
			memory[sPt+2] = dec[2]
			memory[sPt+3] = dec[3]
			memory[sPt+4] = dec[4]
			memory[sPt+5] = dec[5]
			memory[sPt+6] = dec[6]
			memory[sPt+7] = dec[7]
			return sPt, ePt + 1
		},
	},
	{ // Inc Byte
		Pat: "++.",
		Op:  0x41,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			memory[sPt]++
			return sPt, ePt + 1
		},
	},
	{ // Inc Int16
		Pat: "++o",
		Op:  0x42,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			dec := I16tob(Btoi16(memory[sPt:sPt+2]) + 1)
			memory[sPt] = dec[0]
			memory[sPt+1] = dec[1]
			return sPt, ePt + 1
		},
	},
	{ // Inc Int32
		Pat: "++O",
		Op:  0x43,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			dec := I32tob(Btoi32(memory[sPt:sPt+4]) + 1)
			memory[sPt] = dec[0]
			memory[sPt+1] = dec[1]
			memory[sPt+2] = dec[2]
			memory[sPt+3] = dec[3]
			return sPt, ePt + 1
		},
	},
	{ // Inc Int64
		Pat: "++",
		Op:  0x44,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			dec := I64tob(Btoi64(memory[sPt:sPt+8]) + 1)
			memory[sPt] = dec[0]
			memory[sPt+1] = dec[1]
			memory[sPt+2] = dec[2]
			memory[sPt+3] = dec[3]
			memory[sPt+4] = dec[4]
			memory[sPt+5] = dec[5]
			memory[sPt+6] = dec[6]
			memory[sPt+7] = dec[7]
			return sPt, ePt + 1
		},
	},
	{ // Byte push (addr)
		Pat:     "<.#",
		Op:      0x45,
		ArgSize: 64,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := Btoi64(code[ePt+1 : ePt+9])
			memory[sPt-1] = memory[addr]
			return sPt - 1, ePt + 9
		},
	},
	{ // Int16 push (addr)
		Pat:     "<o#",
		Op:      0x46,
		ArgSize: 64,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := Btoi64(code[ePt+1 : ePt+9])
			memory[sPt-2] = memory[addr]
			memory[sPt-1] = memory[addr+1]
			return sPt - 2, ePt + 9
		},
	},
	{ // Int32 push (addr)
		Pat:     "<O#",
		Op:      0x47,
		ArgSize: 64,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := Btoi64(code[ePt+1 : ePt+9])
			memory[sPt-4] = memory[addr]
			memory[sPt-3] = memory[addr+1]
			memory[sPt-2] = memory[addr+2]
			memory[sPt-1] = memory[addr+3]
			return sPt - 4, ePt + 9
		},
	},
	{ // Int64 push (addr)
		Pat:     "<#",
		Op:      0x48,
		ArgSize: 64,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := Btoi64(code[ePt+1 : ePt+9])
			memory[sPt-8] = memory[addr]
			memory[sPt-7] = memory[addr+1]
			memory[sPt-6] = memory[addr+2]
			memory[sPt-5] = memory[addr+3]
			memory[sPt-4] = memory[addr+4]
			memory[sPt-3] = memory[addr+5]
			memory[sPt-2] = memory[addr+6]
			memory[sPt-1] = memory[addr+7]
			return sPt - 8, ePt + 9
		},
	},
	{ // Break
		Pat: "(/)",
		Op:  0x49,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			// no-op
			return sPt, ePt + 1
		},
	},
}
