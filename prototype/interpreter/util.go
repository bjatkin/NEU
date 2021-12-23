package main

func bytesToInt(bytes []byte) int {
	return int(bytes[0]) | int(bytes[1])<<8 | int(bytes[2])<<16 |
		int(bytes[3])<<24 | int(bytes[4])<<32 | int(bytes[5])<<40 |
		int(bytes[6])<<48 | int(bytes[7])<<56
}
