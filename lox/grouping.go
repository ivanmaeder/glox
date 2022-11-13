package lox

type Grouping struct {
	Expression Expr
}

func (r Grouping) Accept(visitor Visitor[string]) string {
	return visitor.VisitGroupingExpr(r)
}

