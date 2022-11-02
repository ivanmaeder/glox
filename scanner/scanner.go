package scanner

import (
	"glox/pkg/tokens"
)

type errorHandler func(int, string)

type scanner struct {
	source    string
	tokens    []tokens.Token
	start     int
	current   int
	line      int
	flagError errorHandler
}

func NewScanner(source string, flagError errorHandler) scanner {
	scanner := scanner{}

	scanner.source = source

	scanner.start = 0
	scanner.current = 0
	scanner.line = 1

	scanner.flagError = flagError

	return scanner
}

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) addToken(tokenType tokens.TokenType, literal string) {
	text := s.source[s.start:s.current]

	s.tokens = append(s.tokens, tokens.Token{
		TokenType: tokenType,
		Lexeme:    text,
		Literal:   literal,
		Line:      s.line,
	})
}

func (s *scanner) advance() byte {
	c := s.source[s.current]

	s.current++

	return c
}

func (s *scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(tokens.LEFT_PAREN, "")
	case ')':
		s.addToken(tokens.RIGHT_PAREN, "")
	case '{':
		s.addToken(tokens.LEFT_BRACE, "")
	case '}':
		s.addToken(tokens.RIGHT_BRACE, "")
	case ',':
		s.addToken(tokens.COMMA, "")
	case '.':
		s.addToken(tokens.DOT, "")
	case '-':
		s.addToken(tokens.MINUS, "")
	case '+':
		s.addToken(tokens.PLUS, "")
	case ';':
		s.addToken(tokens.SEMICOLON, "")
	case '*':
		s.addToken(tokens.STAR, "")
	default:
		var lineFeed byte = 10
		if c != lineFeed {
			s.flagError(s.line, "Unexpected character.")
		}
	}
}

func (s *scanner) ScanTokens() []tokens.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, tokens.Token{
		TokenType: tokens.EOF,
		Lexeme:    "",
		Literal:   "",
		Line:      s.line,
	})

	return s.tokens
}
