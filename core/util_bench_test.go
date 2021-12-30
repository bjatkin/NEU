package core

import "testing"

func BenchmarkI64tob(b *testing.B) {
	tests := []int{
		1029384,
		-50,
		0x7fffffffffffffff,
		-0x7fffffffffffffff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			Btoi64(I64tob(tt))
		}
	}
}

func BenchmarkAsi64(b *testing.B) {
	tests := []int{
		1029384,
		-50,
		0x7fffffffffffffff,
		-0x7fffffffffffffff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b := I64asb(&tt)
			Asi64(&b[0])
		}
	}
}

func BenchmarkAsu64(b *testing.B) {
	tests := []uint{
		1029384,
		0xffffffffffffffff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b := U64asb(&tt)
			Asu64(&b[0])
		}
	}
}

func BenchmarkI32tob(b *testing.B) {
	tests := []int32{
		10,
		-10,
		0x7fffffff,
		-0x7fffffff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			Btoi32(I32tob(tt))
		}
	}
}

func BenchmarkAsi32(b *testing.B) {
	tests := []int32{
		10,
		-10,
		0x7fffffff,
		-0x7fffffff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b := I32asb(&tt)
			Asi32(&b[0])
		}
	}
}

func BenchmarkAsu32(b *testing.B) {
	tests := []uint32{
		10,
		0xffffffff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b := U32asb(&tt)
			Asu32(&b[0])
		}
	}
}
func BenchmarkI16tob(b *testing.B) {
	tests := []int16{
		10,
		-10,
		0x7fff,
		-0x7fff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			Btoi16(I16tob(tt))
		}
	}
}

func BenchmarkAsi16(b *testing.B) {
	tests := []int16{
		10,
		-10,
		0x7fff,
		-0x7fff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b := I16asb(&tt)
			Asi16(&b[0])
		}
	}
}

func BenchmarkAsu16(b *testing.B) {
	tests := []uint16{
		10,
		0xffff,
		0,
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b := U16asb(&tt)
			Asu16(&b[0])
		}
	}
}
