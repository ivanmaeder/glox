package main

import (
	"bufio"
	"fmt"
	"glox/lox"
	"glox/pkg/tokens"
	"glox/scanner"
	"io"
	"os"
)

var hadError = false

func execute() {
	argLength := len(os.Args)

	switch argLength {
	case 2:
		runPrompt()
	case 3:
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

	parser := lox.Parser{
		Tokens:    tokens,
		Current:   0,
		FlagError: flagTokenError,
	}

	expression := parser.Parse()

	if hadError {
		return
	}

	fmt.Println(expression)

	printer := Printer{}
	fmt.Println(printer.Print(expression))
}

func flagError(line int, message string) {
	report(line, "", message)
}

func flagTokenError(token tokens.Token, err error) {
	if token.TokenType == tokens.EOF {
		report(token.Line, " at end", err.Error())
	} else {
		report(token.Line, " at '"+token.Lexeme+"'", err.Error())
	}
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d ] Error%s: %s\n", line, where, message)

	hadError = true
}
