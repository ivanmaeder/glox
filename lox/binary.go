package lox

import "glox/pkg/tokens"

type Binary struct {
	Left Expr
	Operator tokens.Token
	Right Expr
}

func (r Binary) Accept(visitor Visitor[string]) string {
	return visitor.VisitBinaryExpr(r)
}

