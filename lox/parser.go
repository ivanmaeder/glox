package lox

import "glox/pkg/tokens"

type Parser struct {
	Tokens  []tokens.Token
	Current int
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(tokens.BANG_EQUAL, tokens.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = Binary{expr, operator, right}
	}

	return expr
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
		p.current++
	}

	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == tokens.EOF
}

func (p *Parser) peek() tokens.Token {
	return p.tokens.get(p.current) //
}

func (p *Parser) previous() tokens.Token {
	return p.tokens.get(p.current - 1) //
}

func (p *Parser) comparison() Expr {
	expr := p.term()

	for p.match(tokens.GREATER, tokens.GREATER_EQUAL, tokens.LESS, tokens.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for p.match(tokens.MINUS, tokens.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for p.match(tokens.SLASH, tokens.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) unary() Expr {
	if p.match(tokens.BANG, tokens.MINUS) {
		operator := p.previous()
		right := p.unary()
		return Unary{operator, right}
	}

	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(tokens.FALSE) {
		return Literal{false}
	}

	if p.match(tokens.TRUE) {
		return Literal{true}
	}

	if p.match(tokens.NIL) {
		return Literal{nil}
	}

	if p.match(tokens.NUMBER, tokens.STRING) {
		return Literal(p.previous().Literal)
	}

	if p.match(tokens.LEFT_PAREN) {
		expr := p.expression()
		p.consume(tokens.RIGHT_PAREN, "Expect ')' after expression.")
		return Grouping{expr}
	}

	panic(555)
}

func (p *Parser) consume(tokenType tokens.TokenType, message string) tokens.Token {
	if p.check(tokenType) {
		return p.advance()
	}

	//throw error(peek(), message); // FIXME
	panic(666)
}

//func (p *Parser) error(token tokens.Token, message strings) ParseError {
//	Lox.error(token, message);
//
//	return ParseError{};
//}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().tokenType == tokens.SEMICOLON {
			return
		}

		switch p.peek().tokenType {
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
	return p.expression()
	//try {
	//	return expression();
	//} catch (ParseError error) {
	//	return null;
	//}
}
