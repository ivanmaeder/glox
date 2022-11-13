package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: glox execute [script]")
		fmt.Println("       glox generate <directory>")
		fmt.Println("       glox print")

		os.Exit(64)
	}

	switch os.Args[1] {
	case "execute":
		execute()
	case "generate":
		generate()
	case "print":
		print()
	}
}
