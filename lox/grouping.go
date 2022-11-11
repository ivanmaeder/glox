package lox

type Grouping struct {
	Expression Expr
}

func NewGrouping(expression Expr) Grouping {
	return Grouping{expression}
}

func (r Grouping) Accept(visitor Visitor[string]) string {
	return visitor.VisitGroupingExpr(r)
}

