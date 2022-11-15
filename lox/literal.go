package lox

type Literal struct {
	Value any
}

func (r Literal) Accept(visitor Visitor) any {
	return visitor.VisitLiteralExpr(r)
}

