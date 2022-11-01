package tokens

import "fmt"

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   string //object?
	Line      int
}

func (t Token) String() string {
	return fmt.Sprintf("%d %s %s", t.TokenType, t.Lexeme, t.Literal)
}
