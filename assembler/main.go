package main

import (
	"fmt"

	"github.com/bjatkin/neu_interpreter/assembler/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
