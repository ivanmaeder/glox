package lox

type Literal struct {
	Value any
}

func (r Literal) Accept(visitor Visitor[string]) string {
	return visitor.VisitLiteralExpr(r)
}

