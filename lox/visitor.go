package lox

type Visitor interface {
	visitBinaryExpr(expr *Binary) Binary
	visitGroupingExpr(expr *Grouping) Grouping
	visitLiteralExpr(expr *Literal) Literal
	visitUnaryExpr(expr *Unary) Unary
}
