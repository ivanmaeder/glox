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

/*
  GO
  - in interfaces we define "method sets"
	- is-a relationship:

		type Person struct {
			Name string
		}

		func (p *Person) Talk() {
			fmt.Println("Hi, my name is", p.Name)
		}

		type Android struct {
		  Person // this gives all Androids access to Talk()
		  Model string
		}

		a := new(Android)
		a.Talk() //an Android is a person

	- another option is for a struct to have the same function(s) as an interface
	  (like duck typing but they call it structural typing because it happens at
	  compile-time and duck typing is usually at runtime),

		type Shape interface {
			Area() float32
		}

		type Rectangle struct {
			length float32
			width  float32
		}

		func (r *Rectangle) Area() float32 {
			return r.length * r.width
		}

		var shape Shape = &Rectangle{
			length: 10,
			width:  20,
		}
  - unexported (private) is really package-level
	- packages are just 1-word (not paths)
	- naked returns
	- short var declarations (:=) are only available inside functions
	- const <name> (<type>) = <value>
	- ifs can include a short statement: if <statement>; condition {...}
		- vars in this statement are in scope in the if/else block only
	- case statements can include logic:

		switch {
		case a <= 12:
			//...
		case a > 12:
			//...
		}

	- a deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns
	- pointers
		- var p *int
		- p        //pointer to p
		- &p       //value
		- *p = 123 //set the value of p
		- with structs accessing fields "should" be: (*s).field
			- but the s.field shorthand is equivalent
	- this is a struct literal: Vertex{val, val}
	- an array's length is part of its type!
		- "when you assign or pass around an array value you will make a copy of its contents. (To avoid the copy you could pass a pointer to the array, but then that’s a pointer to an array, not an array.) One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value."
		- init: [3]string{"a", "b", "c"} (or [...]string{"a", "b", "c"})
		- slices are a dynamically-sized view into the elements of an array
			- create a slice with arr[n:m]
			- other slices with the same underlying array will see changes
			- this creates a slice (and array), []string{"a", "b", "c"}
			- slice capacity = number of elements in the underlying array counting from the first element in the slice
				- len(s), cap(s)
		- use make([]string, size[, capacity]) to create slices with dynamically-sized arrays
		- use `slice = append(slice, ...values)` to add to a slice (this will increase the capacity of the array)
*/
