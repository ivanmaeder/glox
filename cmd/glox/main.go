package main

import (
	"bufio"
	"fmt"
	"glox/scanner"
	"io"
	"os"
)

/*
	TERM            LEXICAL GRAMMAR  SYNTACTIC GRAMMAR
	Alphabet        Characters       Lexeme/token
	String          Lexeme/token     Expression
	Implemented by  Scanner          Parser

	SCANNER
	-	rules for how characters get grouped into tokens—was called a regular language
	-	emits a flat sequence of tokens
	-	not enough for arbitrarily-nested structures

		s.scanTokens()
			s.scanToken() finds the right token, sometimes by peeking ahead or e.g., looking for the closing string terminator
				s.addToken()
					s.tokens.append()

	PARSER
	- which strings are valid and which aren't
	- rules = productions
		- head
		- body = a list of symbols
			- two types of symbols
				- terminal = a "letter" in the grammar (token/lexeme) (no more "moves" in the game)
				- nonterminal = reference to another rule; play that rule and insert whatever it produces here
	- derivations = generate strings that are in the grammar

  GO
	- is-a relationship

		type Person struct {
			Name string
		}

		func (p *Person) Talk() {
			fmt.Println("Hi, my name is", p.Name)
		}

		type Android struct {
		  Person
		  Model string
		}

		a := new(Android)
		a.Talk() //an Android is a person

  - in interfaces we define "method sets"
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
