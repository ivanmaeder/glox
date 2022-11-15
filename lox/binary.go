package lox

import "glox/pkg/tokens"

type Binary struct {
	Left Expr
	Operator tokens.Token
	Right Expr
}

func (r Binary) Accept(visitor Visitor) any {
	return visitor.VisitBinaryExpr(r)
}

