package lox

type Literal struct {
	value any
}

func NewLiteral(value any) Literal {
	return Literal{value}
}

func (r *Literal) accept(visitor Visitor) Literal {
	return visitor.visitLiteralExpr(r)
}
