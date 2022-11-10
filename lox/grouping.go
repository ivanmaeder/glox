package lox

type Grouping struct {
	expression Expr
}

func NewGrouping(expression Expr) Grouping {
	return Grouping{expression}
}

func (r *Grouping) accept(visitor Visitor) Grouping {
	return visitor.visitGroupingExpr(r)
}
