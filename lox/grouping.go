package lox

type Grouping struct {
	Expression Expr
}

func (r Grouping) Accept(visitor Visitor) any {
	return visitor.VisitGroupingExpr(r)
}

