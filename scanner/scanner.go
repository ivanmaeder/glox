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
	case '!':
		if s.match('=') {
			s.addToken(tokens.BANG_EQUAL, "")
		} else {
			s.addToken(tokens.BANG, "")
		}
	case '=':
		if s.match('=') {
			s.addToken(tokens.EQUAL_EQUAL, "")
		} else {
			s.addToken(tokens.EQUAL, "")
		}
	case '<':
		if s.match('=') {
			s.addToken(tokens.LESS_EQUAL, "")
		} else {
			s.addToken(tokens.LESS, "")
		}
	case '>':
		if s.match('=') {
			s.addToken(tokens.GREATER_EQUAL, "")
		} else {
			s.addToken(tokens.GREATER, "")
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(tokens.SLASH, "")
		}
	case ' ':
		break
	case '\r':
		break
	case '\t':
		break
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		var lineFeed byte = 10
		if c == lineFeed {
			break
		}

		if c >= '0' && c <= '9' {
			s.number()
			break
		}

		s.flagError(s.line, "Unexpected character.")
	}
}

func (s *scanner) number() {
	for s.peek() >= '0' && s.peek() <= '9' {
		s.advance()
	}

	if s.peek() == '.' && s.peekNext() >= '0' && s.peekNext() <= '9' {
		s.advance()
	}

	for s.peek() >= '0' && s.peek() <= '9' {
		s.advance()
	}

	value := s.source[s.start:s.current]
	s.addToken(tokens.NUMBER, value)
}

func (s *scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}

		s.advance()
	}

	if s.isAtEnd() {
		s.flagError(s.line, "Unterminated string.")
		return
	}

	s.advance() //the closing "

	//trim surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addToken(tokens.STRING, value)
}

func (s *scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++

	return true
}

func (s *scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
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
