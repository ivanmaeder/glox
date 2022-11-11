package main

import (
	"fmt"
	"glox/lox"
	"glox/pkg/tokens"
)

type Printer struct {
	//
}

func (p *Printer) VisitBinaryExpr(expr lox.Binary) string {
	return p.parenthesize(expr.Operator.Lexeme, []lox.Expr{expr.Left, expr.Right})
}

func (p *Printer) VisitGroupingExpr(expr lox.Grouping) string {
	return p.parenthesize("group", []lox.Expr{expr.Expression})
}

func (p *Printer) VisitLiteralExpr(expr lox.Literal) string {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprintf("%v", expr.Value)
}

func (p *Printer) VisitUnaryExpr(expr lox.Unary) string {
	return p.parenthesize(expr.Operator.Lexeme, []lox.Expr{expr.Right})
}

func (p *Printer) print(expr lox.Expr) string {
	return expr.Accept(p)
}

func (p *Printer) parenthesize(name string, exprs []lox.Expr) string {
	str := "(" + name
	for _, expr := range exprs {
		str += " "
		str += expr.Accept(p)
	}
	str += ")"

	return str
}

func main() {
	minus := tokens.Token{
		TokenType: tokens.MINUS,
		Lexeme:    "-",
		Literal:   nil,
		Line:      1,
	}

	unary := lox.NewUnary(minus, lox.Literal{Value: 123})

	star := tokens.Token{
		TokenType: tokens.STAR,
		Lexeme:    "*",
		Literal:   nil,
		Line:      1,
	}

	grouping := lox.NewGrouping(lox.NewLiteral(45.67))

	binary := lox.NewBinary(unary, star, grouping)

	fmt.Println(new(Printer).print(binary))
	//expr := lox.NewBinary(
	//	lox.NewUnary(tokens.Token{tokens.TokenType.MINUS, "-", nil, 1}, lox.NewLiteral((123))),
	//	tokens.Token{tokens.TokenType.STAR, "*", nil, 1},
	//	lox.NewGrouping(lox.NewLiteral(45.67)),
	//)

	//fmt.Println(new(Printer).print())
	/*
		Expr expression = new Expr.Binary(
			new Expr.Unary(
					new Token(TokenType.MINUS, "-", null, 1),
					new Expr.Literal(123)),
			new Token(TokenType.STAR, "*", null, 1),
			new Expr.Grouping(
					new Expr.Literal(45.67)));

		System.out.println(new AstPrinter().print(expression));
	*/
}
