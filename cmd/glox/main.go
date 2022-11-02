package main

import (
	"bufio"
	"fmt"
	"glox/scanner"
	"io"
	"os"
)

/*
   new Scanner(code)
   s.scanTokens()
     s.scanToken() "finds" the right token
       s.addToken()
         s.tokens.append()

*/
var hadError = false

func main() {
	argLength := len(os.Args)

	switch argLength {
	case 1:
		runPrompt()
	case 2:
		runFile(os.Args[0])
	default:
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	}
}

func runFile(path string) {
	content, err := os.ReadFile(os.Args[1])

	if err != nil {
		panic(nil)
	}

	run(string(content))

	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			panic(nil)
		}

		if line == "\n" {
			break
		}

		run(line)
		hadError = false
	}
}

func run(code string) {
	s := scanner.NewScanner(code, flagError)
	tokens := s.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func flagError(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d ] Error%s: %s\n", line, where, message)

	hadError = true
}
