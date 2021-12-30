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
			add := *Asi16(&memory[sPt]) + *Asi16(&memory[sPt+2])
			addb := I16asb(&add)
			memory[sPt+2] = addb[0]
			memory[sPt+3] = addb[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Add
		Pat: "+O",
		Op:  0x02,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			add := *Asi32(&memory[sPt]) + *Asi32(&memory[sPt+4])
			addb := I32asb(&add)
			memory[sPt+4] = addb[0]
			memory[sPt+5] = addb[1]
			memory[sPt+6] = addb[2]
			memory[sPt+7] = addb[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Add
		Pat: "+",
		Op:  0x03,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			add := *Asi64(&memory[sPt]) + *Asi64(&memory[sPt+8])
			addb := I64asb(&add)
			memory[sPt+8] = addb[0]
			memory[sPt+9] = addb[1]
			memory[sPt+10] = addb[2]
			memory[sPt+11] = addb[3]
			memory[sPt+12] = addb[4]
			memory[sPt+13] = addb[5]
			memory[sPt+14] = addb[6]
			memory[sPt+15] = addb[7]
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
			sub := *Asi16(&memory[sPt]) - *Asi16(&memory[sPt+2])
			subb := I16asb(&sub)
			memory[sPt+2] = subb[0]
			memory[sPt+3] = subb[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Minus
		Pat: "-O",
		Op:  0x06,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			sub := *Asi32(&memory[sPt]) - *Asi32(&memory[sPt+4])
			subb := I32asb(&sub)
			memory[sPt+4] = subb[0]
			memory[sPt+5] = subb[1]
			memory[sPt+6] = subb[2]
			memory[sPt+7] = subb[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Minus
		Pat: "-",
		Op:  0x07,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			sub := *Asi64(&memory[sPt]) - *Asi64(&memory[sPt+8])
			subb := I64asb(&sub)
			memory[sPt+8] = subb[0]
			memory[sPt+9] = subb[1]
			memory[sPt+10] = subb[2]
			memory[sPt+11] = subb[3]
			memory[sPt+12] = subb[4]
			memory[sPt+13] = subb[5]
			memory[sPt+14] = subb[6]
			memory[sPt+15] = subb[7]
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
			mul := *Asi16(&memory[sPt]) * *Asi16(&memory[sPt+2])
			mulb := I16asb(&mul)
			memory[sPt+2] = mulb[0]
			memory[sPt+3] = mulb[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Multiply
		Pat: "*O",
		Op:  0x0a,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mul := *Asi32(&memory[sPt]) * *Asi32(&memory[sPt+4])
			mulb := I32asb(&mul)
			memory[sPt+4] = mulb[0]
			memory[sPt+5] = mulb[1]
			memory[sPt+6] = mulb[2]
			memory[sPt+7] = mulb[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Multiply
		Pat: "*",
		Op:  0x0b,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mul := *Asi64(&memory[sPt]) * *Asi64(&memory[sPt+8])
			mulb := I64asb(&mul)
			memory[sPt+8] = mulb[0]
			memory[sPt+9] = mulb[1]
			memory[sPt+10] = mulb[2]
			memory[sPt+11] = mulb[3]
			memory[sPt+12] = mulb[4]
			memory[sPt+13] = mulb[5]
			memory[sPt+14] = mulb[6]
			memory[sPt+15] = mulb[7]
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
			div := *Asi16(&memory[sPt]) / *Asi16(&memory[sPt+2])
			divb := I16asb(&div)
			memory[sPt+2] = divb[0]
			memory[sPt+3] = divb[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 divide
		Pat: "/O",
		Op:  0x0e,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			div := *Asi32(&memory[sPt]) / *Asi32(&memory[sPt+4])
			divb := I32asb(&div)
			memory[sPt+4] = divb[0]
			memory[sPt+5] = divb[1]
			memory[sPt+6] = divb[2]
			memory[sPt+7] = divb[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 divide
		Pat: "/",
		Op:  0x0f,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			div := *Asi64(&memory[sPt]) / *Asi64(&memory[sPt+8])
			divb := I64asb(&div)
			memory[sPt+8] = divb[0]
			memory[sPt+9] = divb[1]
			memory[sPt+10] = divb[2]
			memory[sPt+11] = divb[3]
			memory[sPt+12] = divb[4]
			memory[sPt+13] = divb[5]
			memory[sPt+14] = divb[6]
			memory[sPt+15] = divb[7]
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
			addr := *Asi64(&memory[sPt])
			memory[addr] = memory[sPt+8]
			return sPt + 9, ePt + 1
		},
	},
	{ // Int16 pop
		Pat: ">o",
		Op:  0x15,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := *Asi64(&memory[sPt])
			memory[addr] = memory[sPt+8]
			memory[addr+1] = memory[sPt+9]
			return sPt + 10, ePt + 1
		},
	},
	{ // Int32 pop
		Pat: ">O",
		Op:  0x16,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := *Asi64(&memory[sPt])
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
			addr := *Asi64(&memory[sPt])
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
			or := *Asi16(&memory[sPt]) | *Asi16(&memory[sPt+2])
			orb := I16asb(&or)
			memory[sPt+2] = orb[0]
			memory[sPt+3] = orb[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Bitwise Or
		Pat: "|O",
		Op:  0x1a,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := *Asi32(&memory[sPt]) | *Asi32(&memory[sPt+4])
			orb := I32asb(&or)
			memory[sPt+4] = orb[0]
			memory[sPt+5] = orb[1]
			memory[sPt+6] = orb[2]
			memory[sPt+7] = orb[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Bitwise Or
		Pat: "|",
		Op:  0x1b,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			or := *Asi64(&memory[sPt]) | *Asi64(&memory[sPt+8])
			orb := I64asb(&or)
			memory[sPt+8] = orb[0]
			memory[sPt+9] = orb[1]
			memory[sPt+10] = orb[2]
			memory[sPt+11] = orb[3]
			memory[sPt+12] = orb[4]
			memory[sPt+13] = orb[5]
			memory[sPt+14] = orb[6]
			memory[sPt+15] = orb[7]
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
			and := *Asi16(&memory[sPt]) & *Asi16(&memory[sPt+2])
			andb := I16asb(&and)
			memory[sPt+2] = andb[0]
			memory[sPt+3] = andb[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Bitwise And
		Pat: "&O",
		Op:  0x1e,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			and := *Asi32(&memory[sPt]) & *Asi32(&memory[sPt+4])
			andb := I32asb(&and)
			memory[sPt+4] = andb[0]
			memory[sPt+5] = andb[1]
			memory[sPt+6] = andb[2]
			memory[sPt+7] = andb[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Bitwise And
		Pat: "&",
		Op:  0x1f,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			and := *Asi64(&memory[sPt]) & *Asi64(&memory[sPt+8])
			andb := I64asb(&and)
			memory[sPt+8] = andb[0]
			memory[sPt+9] = andb[1]
			memory[sPt+10] = andb[2]
			memory[sPt+11] = andb[3]
			memory[sPt+12] = andb[4]
			memory[sPt+13] = andb[5]
			memory[sPt+14] = andb[6]
			memory[sPt+15] = andb[7]
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
			xor := *Asi16(&memory[sPt]) ^ *Asi16(&memory[sPt+2])
			xorb := I16asb(&xor)
			memory[sPt+2] = xorb[0]
			memory[sPt+3] = xorb[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 Bitwise Xor
		Pat: "^O",
		Op:  0x22,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			xor := *Asi32(&memory[sPt]) ^ *Asi32(&memory[sPt+4])
			xorb := I32asb(&xor)
			memory[sPt+4] = xorb[0]
			memory[sPt+5] = xorb[1]
			memory[sPt+6] = xorb[2]
			memory[sPt+7] = xorb[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 Bitwise Xor
		Pat: "^",
		Op:  0x23,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			xor := *Asi64(&memory[sPt]) ^ *Asi64(&memory[sPt+8])
			xorb := I64asb(&xor)
			memory[sPt+8] = xorb[0]
			memory[sPt+9] = xorb[1]
			memory[sPt+10] = xorb[2]
			memory[sPt+11] = xorb[3]
			memory[sPt+12] = xorb[4]
			memory[sPt+13] = xorb[5]
			memory[sPt+14] = xorb[6]
			memory[sPt+15] = xorb[7]
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
			shift := *Asi16(&memory[sPt+1]) << memory[sPt]
			shiftb := I16asb(&shift)
			memory[sPt+1] = shiftb[0]
			memory[sPt+2] = shiftb[1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int32 Bitwise LeftShift
		Pat: "<<O",
		Op:  0x26,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := *Asi32(&memory[sPt+1]) << memory[sPt]
			shiftb := I32asb(&shift)
			memory[sPt+1] = shiftb[0]
			memory[sPt+2] = shiftb[1]
			memory[sPt+3] = shiftb[2]
			memory[sPt+4] = shiftb[3]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int64 Bitwise LeftShift
		Pat: "<<",
		Op:  0x27,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := *Asi64(&memory[sPt+1]) << memory[sPt]
			shiftb := I64asb(&shift)
			memory[sPt+1] = shiftb[0]
			memory[sPt+2] = shiftb[1]
			memory[sPt+3] = shiftb[2]
			memory[sPt+4] = shiftb[3]
			memory[sPt+5] = shiftb[4]
			memory[sPt+6] = shiftb[5]
			memory[sPt+7] = shiftb[6]
			memory[sPt+8] = shiftb[7]
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
			shift := *Asu16(&memory[sPt+1]) >> memory[sPt]
			shiftB := U16asb(&shift)
			memory[sPt+1] = shiftB[0]
			memory[sPt+2] = shiftB[1]
			return sPt + 1, ePt + 1
		},
	},
	{ // Int32 Bitwise RightShift
		Pat: ">>O",
		Op:  0x2a,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := *Asu32(&memory[sPt+1]) >> memory[sPt]
			shiftB := U32asb(&shift)
			memory[sPt+1] = shiftB[0]
			memory[sPt+2] = shiftB[1]
			memory[sPt+3] = shiftB[2]
			memory[sPt+4] = shiftB[3]
			return sPt + 1, ePt + 1
		},
	},
	{ // Bitwise RightShift
		Pat: ">>",
		Op:  0x2b,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			shift := *Asu64(&memory[sPt+1]) >> uint(memory[sPt])
			shiftB := U64asb(&shift)
			memory[sPt+1] = shiftB[0]
			memory[sPt+2] = shiftB[1]
			memory[sPt+3] = shiftB[2]
			memory[sPt+4] = shiftB[3]
			memory[sPt+5] = shiftB[4]
			memory[sPt+6] = shiftB[5]
			memory[sPt+7] = shiftB[6]
			memory[sPt+8] = shiftB[7]
			return sPt + 1, ePt + 1
		},
	},
	{ // Jump If Greater
		Pat: "?>.",
		Op:  0x2c,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if memory[sPt] > memory[sPt+1] {
				return sPt + 10, *Asu64(&memory[sPt+2])
			}
			return sPt + 10, ePt + 1
		},
	},
	{ // Int16 Jump If Greater
		Pat: "?>o",
		Op:  0x2d,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if *Asi16(&memory[sPt]) > *Asi16(&memory[sPt+2]) {
				return sPt + 12, *Asu64(&memory[sPt+4])
			}
			return sPt + 12, ePt + 1
		},
	},
	{ // Int32 Jump If Greater
		Pat: "?>O",
		Op:  0x2e,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if *Asi32(&memory[sPt]) > *Asi32(&memory[sPt+4]) {
				return sPt + 16, *Asu64(&memory[sPt+8])
			}
			return sPt + 16, ePt + 1
		},
	},
	{ // Int64 Jump If Greater
		Pat: "?>",
		Op:  0x2f,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if *Asi64(&memory[sPt]) > *Asi64(&memory[sPt+8]) {
				return sPt + 24, *Asu64(&memory[sPt+16])
			}
			return sPt + 24, ePt + 1
		},
	},
	{ // Jump If Less
		Pat: "?<.",
		Op:  0x30,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if memory[sPt] < memory[sPt+1] {
				return sPt + 10, *Asu64(&memory[sPt+2])
			}
			return sPt + 10, ePt + 1
		},
	},
	{ // Int16 Jump If Less
		Pat: "?<o",
		Op:  0x31,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if *Asi16(&memory[sPt]) < *Asi16(&memory[sPt+2]) {
				return sPt + 12, *Asu64(&memory[sPt+4])
			}
			return sPt + 12, ePt + 1
		},
	},
	{ // Int32 Jump If Less
		Pat: "?<O",
		Op:  0x32,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if *Asi32(&memory[sPt]) < *Asi32(&memory[sPt+4]) {
				return sPt + 16, *Asu64(&memory[sPt+8])
			}
			return sPt + 16, ePt + 1
		},
	},
	{ // Int64 Jump If Less
		Pat: "?<",
		Op:  0x33,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			if *Asi64(&memory[sPt]) < *Asi64(&memory[sPt+8]) {
				return sPt + 24, *Asu64(&memory[sPt+16])
			}
			return sPt + 24, ePt + 1
		},
	},
	{ // Jump
		Pat: "|>",
		Op:  0x34,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			return sPt + 8, *Asu64(&memory[sPt])
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
			mod := *Asi16(&memory[sPt]) % *Asi16(&memory[sPt+2])
			modb := I16asb(&mod)
			memory[sPt+2] = modb[0]
			memory[sPt+3] = modb[1]
			return sPt + 2, ePt + 1
		},
	},
	{ // Int32 mod
		Pat: "%O",
		Op:  0x37,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mod := *Asi32(&memory[sPt]) % *Asi32(&memory[sPt+4])
			modb := I32asb(&mod)
			memory[sPt+4] = modb[0]
			memory[sPt+5] = modb[1]
			memory[sPt+6] = modb[2]
			memory[sPt+7] = modb[3]
			return sPt + 4, ePt + 1
		},
	},
	{ // Int64 mod
		Pat: "%",
		Op:  0x38,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			mod := *Asi64(&memory[sPt]) % *Asi64(&memory[sPt+8])
			modb := I64asb(&mod)
			memory[sPt+8] = modb[0]
			memory[sPt+9] = modb[1]
			memory[sPt+10] = modb[2]
			memory[sPt+11] = modb[3]
			memory[sPt+12] = modb[4]
			memory[sPt+13] = modb[5]
			memory[sPt+14] = modb[6]
			memory[sPt+15] = modb[7]
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
			dec := *Asi16(&memory[sPt]) - 1
			decb := I16asb(&dec)
			memory[sPt] = decb[0]
			memory[sPt+1] = decb[1]
			return sPt, ePt + 1
		},
	},
	{ // Dec Int32
		Pat: "--O",
		Op:  0x3f,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			dec := *Asi32(&memory[sPt]) - 1
			decb := I32asb(&dec)
			memory[sPt] = decb[0]
			memory[sPt+1] = decb[1]
			memory[sPt+2] = decb[2]
			memory[sPt+3] = decb[3]
			return sPt, ePt + 1
		},
	},
	{ // Dec Int64
		Pat: "--",
		Op:  0x40,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			dec := *Asi64(&memory[sPt]) - 1
			decb := I64asb(&dec)
			memory[sPt] = decb[0]
			memory[sPt+1] = decb[1]
			memory[sPt+2] = decb[2]
			memory[sPt+3] = decb[3]
			memory[sPt+4] = decb[4]
			memory[sPt+5] = decb[5]
			memory[sPt+6] = decb[6]
			memory[sPt+7] = decb[7]
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
			inc := *Asi16(&memory[sPt]) + 1
			incb := I16asb(&inc)
			memory[sPt] = incb[0]
			memory[sPt+1] = incb[1]
			return sPt, ePt + 1
		},
	},
	{ // Inc Int32
		Pat: "++O",
		Op:  0x43,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			inc := *Asi32(&memory[sPt]) + 1
			incb := I32asb(&inc)
			memory[sPt] = incb[0]
			memory[sPt+1] = incb[1]
			memory[sPt+2] = incb[2]
			memory[sPt+3] = incb[3]
			return sPt, ePt + 1
		},
	},
	{ // Inc Int64
		Pat: "++",
		Op:  0x44,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			inc := *Asi64(&memory[sPt]) + 1
			incb := I64asb(&inc)
			memory[sPt] = incb[0]
			memory[sPt+1] = incb[1]
			memory[sPt+2] = incb[2]
			memory[sPt+3] = incb[3]
			memory[sPt+4] = incb[4]
			memory[sPt+5] = incb[5]
			memory[sPt+6] = incb[6]
			memory[sPt+7] = incb[7]
			return sPt, ePt + 1
		},
	},
	{ // Byte push (addr)
		Pat:     "<.#",
		Op:      0x45,
		ArgSize: 64,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := *Asi64(&code[ePt+1])
			memory[sPt-1] = memory[addr]
			return sPt - 1, ePt + 9
		},
	},
	{ // Int16 push (addr)
		Pat:     "<o#",
		Op:      0x46,
		ArgSize: 64,
		Fn: func(memory, code []byte, sPt, ePt uint) (newSPt, newEPt uint) {
			addr := *Asi64(&code[ePt+1])
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
			addr := *Asi64(&code[ePt+1])
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
			addr := *Asi64(&code[ePt+1])
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
