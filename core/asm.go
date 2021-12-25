package core

type OpCode struct {
	Pat     string
	Op      byte
	ArgSize byte
	Fn      func(memory []byte, ePtr, sPtr int) (ePtDelta int, sPtDelta int)
}

var OpCodes = [0xff]OpCode{
	{ // Byte Add
		Pat:     "+.",
		Op:      0x00,
		ArgSize: 8,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			sPt += 1
			memory[sPt] = memory[sPt] + memory[sPt-1]
			return 1, 1
		},
	},
	{ // Int16 Add
		Pat:     "+o",
		Op:      0x01,
		ArgSize: 16,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			add := I16tob(Btoi16(memory[sPt:sPt+2]) + Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = add[0]
			memory[sPt+3] = add[1]
			return 1, 2
		},
	},
	{ // Int32 Add
		Pat:     "+O",
		Op:      0x02,
		ArgSize: 32,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			add := I32tob(Btoi32(memory[sPt:sPt+4]) + Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = add[0]
			memory[sPt+5] = add[1]
			memory[sPt+6] = add[2]
			memory[sPt+7] = add[3]
			return 1, 4
		},
	},
	{ // Int64 Add
		Pat:     "+",
		Op:      0x03,
		ArgSize: 64,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
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
		Pat:     "-.",
		Op:      0x04,
		ArgSize: 8,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			sPt += 1
			memory[sPt] = memory[sPt-1] - memory[sPt]
			return 1, 1
		},
	},
	{ // Int16 Minus
		Pat:     "-o",
		Op:      0x05,
		ArgSize: 16,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			add := I16tob(Btoi16(memory[sPt:sPt+2]) - Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = add[0]
			memory[sPt+3] = add[1]
			return 1, 2
		},
	},
	{ // Int32 Minus
		Pat:     "-O",
		Op:      0x06,
		ArgSize: 32,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			add := I32tob(Btoi32(memory[sPt:sPt+4]) - Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = add[0]
			memory[sPt+5] = add[1]
			memory[sPt+6] = add[2]
			memory[sPt+7] = add[3]
			return 1, 4
		},
	},
	{ // Int64 Minus
		Pat:     "-",
		Op:      0x07,
		ArgSize: 64,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
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
		Pat:     "*.",
		Op:      0x08,
		ArgSize: 8,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			sPt += 1
			memory[sPt] = memory[sPt-1] * memory[sPt]
			return 1, 1
		},
	},
	{ // Int16 Multiply
		Pat:     "*o",
		Op:      0x09,
		ArgSize: 16,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			mul := I16tob(Btoi16(memory[sPt:sPt+2]) * Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = mul[0]
			memory[sPt+3] = mul[1]
			return 1, 2
		},
	},
	{ // Int32 Multiply
		Pat:     "*O",
		Op:      0x0a,
		ArgSize: 32,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			mul := I32tob(Btoi32(memory[sPt:sPt+4]) * Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = mul[0]
			memory[sPt+5] = mul[1]
			memory[sPt+6] = mul[2]
			memory[sPt+7] = mul[3]
			return 1, 4
		},
	},
	{ // Int64 Multiply
		Pat:     "*",
		Op:      0x0b,
		ArgSize: 64,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
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
		Pat:     "/.",
		Op:      0x0c,
		ArgSize: 8,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			sPt += 1
			memory[sPt] = memory[sPt-1] / memory[sPt]
			return 1, 1
		},
	},
	{ // Int16 divide
		Pat:     "/o",
		Op:      0x0d,
		ArgSize: 16,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			div := I16tob(Btoi16(memory[sPt:sPt+2]) / Btoi16(memory[sPt+2:sPt+4]))
			memory[sPt+2] = div[0]
			memory[sPt+3] = div[1]
			return 1, 2
		},
	},
	{ // Int32 divide
		Pat:     "/O",
		Op:      0x0e,
		ArgSize: 32,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
			div := I32tob(Btoi32(memory[sPt:sPt+4]) / Btoi32(memory[sPt+4:sPt+8]))
			memory[sPt+4] = div[0]
			memory[sPt+5] = div[1]
			memory[sPt+6] = div[2]
			memory[sPt+7] = div[3]
			return 1, 4
		},
	},
	{ // Int64 divide
		Pat:     "/",
		Op:      0x0f,
		ArgSize: 64,
		Fn: func(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
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
	{ // Byte push
		Pat:     "<.",
		Op:      0x10,
		ArgSize: 8,
		Fn: func(memory []byte, ePt, sPt int) (ePtDelta int, sPtDelta int) {
			memory[sPt-1] = memory[ePt+1]
			return 2, -1
		},
	},
	{ // Int16 push
		Pat:     "<o",
		Op:      0x11,
		ArgSize: 16,
	},
	{ // Int32 push
		Pat:     "<O",
		Op:      0x12,
		ArgSize: 32,
	},
	{ // Int64 push
		Pat:     "<",
		Op:      0x13,
		ArgSize: 64,
	},
	{ // Byte pop
		Pat:     ">.",
		Op:      0x14,
		ArgSize: 64,
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
		Pat:     ">o",
		Op:      0x15,
		ArgSize: 16,
	},
	{ // Int32 pop
		Pat:     ">O",
		Op:      0x16,
		ArgSize: 32,
	},
	{ // Int64 pop
		Pat:     ">",
		Op:      0x17,
		ArgSize: 64,
	},
	{ // Bitwise OR
		Pat:     "|.",
		Op:      0x18,
		ArgSize: 8,
	},
	{ // Int16 Bitwise Or
		Pat:     "|o",
		Op:      0x19,
		ArgSize: 16,
	},
	{ // Int32 Bitwise Or
		Pat:     "|O",
		Op:      0x1a,
		ArgSize: 32,
	},
	{ // Int64 Bitwise Or
		Pat:     "|",
		Op:      0x1b,
		ArgSize: 64,
	},
	{ // Bitwise And
		Pat:     "&.",
		Op:      0x1c,
		ArgSize: 8,
	},
	{ // Int16 Bitwise And
		Pat:     "&o",
		Op:      0x1d,
		ArgSize: 16,
	},
	{ // Int32 Bitwise Or
		Pat:     "&O",
		Op:      0x1e,
		ArgSize: 32,
	},
	{ // Int64 Bitwise Or
		Pat:     "&",
		Op:      0x1f,
		ArgSize: 64,
	},
	{ // Bitwise Xor
		Pat:     "^.",
		Op:      0x20,
		ArgSize: 8,
	},
	{ // Int16 Bitwise Xor
		Pat:     "^o",
		Op:      0x21,
		ArgSize: 16,
	},
	{ // Int32 Bitwise Xor
		Pat:     "^O",
		Op:      0x22,
		ArgSize: 32,
	},
	{ // Int64 Bitwise Xor
		Pat:     "^",
		Op:      0x23,
		ArgSize: 64,
	},
	{ // Bitwise LeftShift
		Pat:     "<<.",
		Op:      0x24,
		ArgSize: 8,
	},
	{ // Int16 Bitwise LeftShift
		Pat:     "<<o",
		Op:      0x25,
		ArgSize: 16,
	},
	{ // Int32 Bitwise LeftShift
		Pat:     "<<O",
		Op:      0x26,
		ArgSize: 32,
	},
	{ // Int64 Bitwise LeftShift
		Pat:     "<<",
		Op:      0x27,
		ArgSize: 64,
	},
	{ // Bitwise RightShift
		Pat:     ">>",
		Op:      0x28,
		ArgSize: 8,
	},
	{ // Int16 jBitwise RightShift
		Pat:     ">>",
		Op:      0x29,
		ArgSize: 16,
	},
	{ // Int32 Bitwise RightShift
		Pat:     ">>",
		Op:      0x2a,
		ArgSize: 32,
	},
	{ // Bitwise RightShift
		Pat:     ">>",
		Op:      0x2b,
		ArgSize: 8,
	},
	{ // Int16 Bitwise RightShift
		Pat:     ">>",
		Op:      0x2c,
		ArgSize: 16,
	},
	{ // Int32 Bitwise RightShift
		Pat:     ">>",
		Op:      0x2d,
		ArgSize: 32,
	},
	{ // Int64 Bitwise RightShift
		Pat:     ">>",
		Op:      0x2e,
		ArgSize: 64,
	},
	{ // Jump If Greater
		Pat:     "?>.",
		Op:      0x2f,
		ArgSize: 64,
	},
	{ // Int16 Jump If Greater
		Pat:     "?>o",
		Op:      0x30,
		ArgSize: 64,
	},
	{ // Int32 Jump If Greater
		Pat:     "?>O",
		Op:      0x31,
		ArgSize: 64,
	},
	{ // Int64 Jump If Greater
		Pat:     "?>",
		Op:      0x32,
		ArgSize: 64,
	},
	{ // Jump If Less
		Pat:     "?<.",
		Op:      0x33,
		ArgSize: 64,
	},
	{ // Int16 Jump If Less
		Pat:     "?<o",
		Op:      0x34,
		ArgSize: 64,
	},
	{ // Int32 Jump If Less
		Pat:     "?<O",
		Op:      0x35,
		ArgSize: 64,
	},
	{ // Int64 Jump If Less
		Pat:     "?<",
		Op:      0x36,
		ArgSize: 64,
	},
	{ // Jump
		Pat:     "|>",
		Op:      0x37,
		ArgSize: 64,
	},
}

func IsLabel(test string) bool {
	if len(test) < 3 {
		return false
	}

	return test[0] == '[' && test[2] == ']'
}

func IsAddressAssignemtn(test []string) bool {
	if len(test) != 3 {
		return false
	}

	if IsAddressName(test[0]) {
		return false
	}

	if test[1] != "=" {
		return false
	}

	if len(test[2]) == 0 {
		return false
	}

	return false
}

func IsAddressName(test string) bool {
	if len(test) < 2 || len(test) > 20 {
		return false
	}

	if test[0] != '\\' {
		return false
	}

	return true
}
