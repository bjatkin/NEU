package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// a := intToBytes(555555555555)
	// fmt.Println(a)
	// b := bytesToInt(a)
	// fmt.Println(b)
	// return
	asm := asembler{}

	asm.addMatcher(newStringMatcher("<", byte(0x01)))
	asm.addMatcher(newStringMatcher("+", byte(0x02)))
	asm.addMatcher(newStringMatcher("-", byte(0x03)))
	asm.addMatcher(newStringMatcher(">", byte(0x04)))
	asm.addMatcher(newStringMatcher("[]", byte(0x05)))
	asm.addMatcher(&decimalMatcher{})

	data, err := asm.execute("example.jy")
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("../interpreter/example.n", data, 0755)
	if err != nil {
		log.Fatalln(err)
	}
}
