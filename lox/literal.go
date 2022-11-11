package lox

type Literal struct {
	Value any
}

func NewLiteral(value any) Literal {
	return Literal{value}
}

func (r Literal) Accept(visitor Visitor[string]) string {
	return visitor.VisitLiteralExpr(r)
}

