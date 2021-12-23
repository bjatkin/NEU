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
		Fn:      ByteAdd,
	},
	{ // Int16 Add
		Pat:     "+o",
		Op:      0x01,
		ArgSize: 16,
		Fn:      Int16Add,
	},
	{ // Int32 Add
		Pat:     "+0",
		Op:      0x02,
		ArgSize: 32,
		Fn:      Int32Add,
	},
	{ // Int64 Add
		Pat:     "+",
		Op:      0x03,
		ArgSize: 64,
		Fn:      Int64Add,
	},
	{ // Byte Minus
		Pat:     "-.",
		Op:      0x04,
		ArgSize: 8,
		Fn:      ByteMinus,
	},
	{ // Int16 Minus
		Pat:     "-o",
		Op:      0x05,
		ArgSize: 16,
		Fn:      Int16Minus,
	},
	{ // Int32 Minus
		Pat:     "-0",
		Op:      0x06,
		ArgSize: 32,
		Fn:      Int32Minus,
	},
	{ // Int64 Minus
		Pat:     "-",
		Op:      0x07,
		ArgSize: 64,
		Fn:      Int64Minus,
	},
	{ // Byte Multiply
		Pat:     "*.",
		Op:      0x08,
		ArgSize: 8,
	},
	{ // Int16 Multiply
		Pat:     "*o",
		Op:      0x09,
		ArgSize: 16,
	},
	{ // Int32 Multiply
		Pat:     "*0",
		Op:      0x0a,
		ArgSize: 32,
	},
	{ // Int64 Multiply
		Pat:     "*",
		Op:      0x0b,
		ArgSize: 64,
	},
	{ // Byte divide
		Pat:     "/.",
		Op:      0x0c,
		ArgSize: 8,
	},
	{ // Int16 divide
		Pat:     "/o",
		Op:      0x0d,
		ArgSize: 16,
	},
	{ // Int32 divide
		Pat:     "/0",
		Op:      0x0e,
		ArgSize: 32,
	},
	{ // Int64 divide
		Pat:     "/",
		Op:      0x0f,
		ArgSize: 64,
	},
	{ // Byte push
		Pat:     "<.",
		Op:      0x10,
		ArgSize: 8,
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
		ArgSize: 8,
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

func ByteAdd(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
	sPt += 1
	memory[sPt] = memory[sPt] + memory[sPt-1]
	return 1, 1
}

func Int16Add(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
	add := I16tob(Btoi16(memory[sPt:sPt+2]) + Btoi16(memory[sPt+2:sPt+4]))
	memory[sPt+2] = add[0]
	memory[sPt+3] = add[1]
	return 1, 2
}

func Int32Add(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
	add := I32tob(Btoi32(memory[sPt:sPt+4]) + Btoi32(memory[sPt+4:sPt+8]))
	memory[sPt+4] = add[0]
	memory[sPt+5] = add[1]
	memory[sPt+6] = add[2]
	memory[sPt+7] = add[3]
	return 1, 4
}

func Int64Add(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
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
}

func ByteMinus(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
	sPt += 1
	memory[sPt] = memory[sPt-1] - memory[sPt]
	return 1, 1
}

func Int16Minus(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
	add := I16tob(Btoi16(memory[sPt:sPt+2]) - Btoi16(memory[sPt+2:sPt+4]))
	memory[sPt+2] = add[0]
	memory[sPt+3] = add[1]
	return 1, 2
}

func Int32Minus(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
	add := I32tob(Btoi32(memory[sPt:sPt+4]) - Btoi32(memory[sPt+4:sPt+8]))
	memory[sPt+4] = add[0]
	memory[sPt+5] = add[1]
	memory[sPt+6] = add[2]
	memory[sPt+7] = add[3]
	return 1, 4
}

func Int64Minus(memory []byte, _, sPt int) (ePtDelta int, sPtDelta int) {
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
}
