package main

import (
	"fmt"
	"os"
)

func main() {
	argLength := len(os.Args)

	switch argLength {
	case 0:
		//runPrompt()
	case 1:
		runFile(os.Args[0])
	default:
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	}
}

func runFile(path string) {
	//
}
