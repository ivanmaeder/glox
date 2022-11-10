package lox

import "glox/pkg/tokens"

type Binary struct {
	left     Expr
	operator tokens.Token
	right    Expr
}

func NewBinary(left Expr, operator tokens.Token, right Expr) Binary {
	return Binary{left, operator, right}
}

func (r *Binary) accept(visitor Visitor) Binary {
	return visitor.visitBinaryExpr(r)
}
