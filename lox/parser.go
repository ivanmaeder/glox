package lox

import (
	"errors"
	gloxerrors "glox/pkg/errors"
	"glox/pkg/tokens"
)

type Parser struct {
	Tokens    []tokens.Token
	Current   int
	FlagError gloxerrors.TokenErrorHandler
}

func (p *Parser) expression() (Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()

	if err != nil {
		return nil, err
	}

	for p.match(tokens.BANG_EQUAL, tokens.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()

		if err != nil {
			return nil, err
		}

		expr = Binary{expr, operator, right}
	}

	return expr, nil
}

func (p *Parser) match(types ...tokens.TokenType) bool {
	for _, element := range types {
		if p.check(element) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType tokens.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().TokenType == tokenType
}

func (p *Parser) advance() tokens.Token {
	if !p.isAtEnd() {
		p.Current++
	}

	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == tokens.EOF
}

func (p *Parser) peek() tokens.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() tokens.Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()

	if err != nil {
		return nil, err
	}

	for p.match(tokens.GREATER, tokens.GREATER_EQUAL, tokens.LESS, tokens.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()

		if err != nil {
			return nil, err
		}

		expr = Binary{expr, operator, right}
	}

	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()

	if err != nil {
		return nil, err
	}

	for p.match(tokens.MINUS, tokens.PLUS) {
		operator := p.previous()
		right, err := p.factor()

		if err != nil {
			return nil, err
		}

		expr = Binary{expr, operator, right}
	}

	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()

	if err != nil {
		return nil, err
	}

	for p.match(tokens.SLASH, tokens.STAR) {
		operator := p.previous()
		right, _ := p.unary()
		expr = Binary{expr, operator, right}
	}

	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(tokens.BANG, tokens.MINUS) {
		operator := p.previous()
		right, _ := p.unary()
		return Unary{operator, right}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	if p.match(tokens.FALSE) {
		return Literal{false}, nil
	}

	if p.match(tokens.TRUE) {
		return Literal{true}, nil
	}

	if p.match(tokens.NIL) {
		return Literal{nil}, nil
	}

	if p.match(tokens.NUMBER, tokens.STRING) {
		return Literal{p.previous().Literal}, nil
	}

	if p.match(tokens.LEFT_PAREN) {
		expr, _ := p.expression()
		_, err := p.consume(tokens.RIGHT_PAREN, "Expect ')' after expression.")

		if err != nil {
			return nil, err
		}

		return Grouping{expr}, nil
	}

	err := &gloxerrors.TokenError{
		Token: p.peek(),
		Err:   errors.New("expect expression"),
	}

	p.FlagError(p.peek(), err)

	return nil, err
}

func (p *Parser) consume(tokenType tokens.TokenType, message string) (tokens.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}

	token, err := p.peek(), &gloxerrors.TokenError{
		Token: p.peek(),
		Err:   errors.New(message),
	}

	p.FlagError(token, err)

	return token, err
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().TokenType == tokens.SEMICOLON {
			return
		}

		switch p.peek().TokenType {
		case tokens.CLASS:
		case tokens.FUN:
		case tokens.VAR:
		case tokens.FOR:
		case tokens.IF:
		case tokens.WHILE:
		case tokens.PRINT:
		case tokens.RETURN:
			return
		}

		p.advance()
	}
}

func (p *Parser) Parse() Expr {
	expression, err := p.expression()

	if err != nil {
		return nil
	}

	return expression
}
