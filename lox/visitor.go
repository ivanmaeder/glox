package lox

type Visitor[K any] interface {
	VisitBinaryExpr(expr Binary) K
	VisitGroupingExpr(expr Grouping) K
	VisitLiteralExpr(expr Literal) K
	VisitUnaryExpr(expr Unary) K
}
