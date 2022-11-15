package lox

import "glox/pkg/tokens"

type Unary struct {
	Operator tokens.Token
	Right Expr
}

func (r Unary) Accept(visitor Visitor) any {
	return visitor.VisitUnaryExpr(r)
}

