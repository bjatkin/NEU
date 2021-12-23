package main

func intToBytes(i int) []byte {
	ret := []byte{
		byte(i),
		byte(i >> 8),
		byte(i >> 16),
		byte(i >> 24),
		byte(i >> 32),
		byte(i >> 40),
		byte(i >> 48),
		byte(i >> 56),
	}
	return ret
}
