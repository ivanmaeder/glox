package lox

import "glox/pkg/tokens"

type Binary struct {
	Left Expr
	Operator tokens.Token
	Right Expr
}

func NewBinary(left Expr, operator tokens.Token, right Expr) Binary {
	return Binary{left, operator, right}
}

func (r Binary) Accept(visitor Visitor[string]) string {
	return visitor.VisitBinaryExpr(r)
}

