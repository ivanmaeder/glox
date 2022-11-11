package lox

type Expr interface {
	Accept(visitor Visitor[string]) string
}
