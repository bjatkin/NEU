package core

type OpCode struct {
	Pat     string
	Op      byte
	ArgSize byte
	Fn      func(memory []byte, ptr, stackPtr int) (ptrShift int, stackPtrShift int)
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
		Pat:     "+0",
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
		Pat:     "-0",
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
		Pat:     "*0",
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
		Pat:     "/0",
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
		Pat:     "<0",
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
			addr := Btoi64(memory[ePt+1 : ePt+9])
			memory[addr] = memory[sPt]
			return 9, 1
		},
	},
	{ // Int16 pop
		Pat:     ">o",
		Op:      0x15,
		ArgSize: 16,
	},
	{ // Int32 pop
		Pat:     ">0",
		Op:      0x16,
		ArgSize: 32,
	},
	{ // Int64 pop
		Pat:     ">",
		Op:      0x17,
		ArgSize: 64,
	},
}
