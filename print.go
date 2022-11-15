package main

import (
	"fmt"
	"glox/lox"
	"glox/pkg/tokens"
)

type Printer struct {
	//
}

func (p *Printer) VisitBinaryExpr(expr lox.Binary) any {
	return p.parenthesize(expr.Operator.Lexeme, []lox.Expr{expr.Left, expr.Right})
}

func (p *Printer) VisitGroupingExpr(expr lox.Grouping) any {
	return p.parenthesize("group", []lox.Expr{expr.Expression})
}

func (p *Printer) VisitLiteralExpr(expr lox.Literal) any {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprintf("%v", expr.Value)
}

func (p *Printer) VisitUnaryExpr(expr lox.Unary) any {
	return p.parenthesize(expr.Operator.Lexeme, []lox.Expr{expr.Right})
}

func (p *Printer) Print(expr lox.Expr) any {
	return expr.Accept(p)
}

func (p *Printer) parenthesize(name string, exprs []lox.Expr) string {
	str := "(" + name
	for _, expr := range exprs {
		str += " "
		str += fmt.Sprintf("%v", expr.Accept(p))
	}
	str += ")"

	return str
}

func print() {
	minus := tokens.Token{
		TokenType: tokens.MINUS,
		Lexeme:    "-",
		Literal:   nil,
		Line:      1,
	}

	unary := lox.Unary{
		Operator: minus,
		Right:    lox.Literal{Value: 123},
	}

	star := tokens.Token{
		TokenType: tokens.STAR,
		Lexeme:    "*",
		Literal:   nil,
		Line:      1,
	}

	grouping := lox.Grouping{Expression: lox.Literal{Value: 45.67}}

	binary := lox.Binary{
		Left:     unary,
		Operator: star,
		Right:    grouping,
	}

	fmt.Println(new(Printer).Print(binary))
}
