package lox

import "glox/pkg/tokens"

type Unary struct {
	operator tokens.Token
	right    Expr
}

func NewUnary(operator tokens.Token, right Expr) Unary {
	return Unary{operator, right}
}

func (r *Unary) accept(visitor Visitor) Unary {
	return visitor.visitUnaryExpr(r)
}
