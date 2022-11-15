package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func generate() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: glox generate <directory>")
		os.Exit(64)
	}

	defineAst(os.Args[2], "Expr", []string{
		"Binary   : left Expr, operator tokens.Token, right Expr",
		"Grouping : expression Expr",
		"Literal  : value any",
		"Unary    : operator tokens.Token, right Expr",
	})
}

func defineAst(outputDirectory string, baseName string, types []string) {
	defineVisitor(outputDirectory, baseName, types)

	for _, element := range types {
		className := strings.TrimSpace(strings.Split(element, ":")[0])
		fields := strings.Split(strings.TrimSpace(strings.Split(element, ":")[1]), ",")

		defineType(outputDirectory, baseName, className, fields)
	}
}

func defineVisitor(outputDirectory string, baseName string, types []string) {
	f, err := os.Create(outputDirectory + "/visitor.go")

	if err != nil {
		panic(nil)
	}

	defer f.Close()

	f.WriteString("package lox\n\n")
	f.WriteString("type Visitor interface {\n")

	for _, element := range types {
		typeName := strings.TrimSpace(strings.Split(element, ":")[0])

		f.WriteString("\tVisit" + typeName + baseName + "(" + strings.ToLower(baseName) + " " + typeName + ") any\n")
	}

	f.WriteString("}\n")
}

func defineType(outputDirectory string, baseName string, className string, fields []string) {
	path := outputDirectory + "/" + strings.ToLower(className) + ".go"

	f, err := os.Create(path)

	if err != nil {
		panic(nil)
	}

	defer f.Close()

	f.WriteString("package lox\n\n")

	if strings.Contains(strings.Join(fields, ","), "tokens.Token") {
		f.WriteString("import \"glox/pkg/tokens\"\n\n")
	}

	f.WriteString("type " + className + " struct {\n")
	for _, element := range fields {
		r := []rune(strings.TrimSpace(element))
		r[0] = unicode.ToUpper(r[0])
		s := string(r)

		f.WriteString("\t" + s + "\n")
	}
	f.WriteString("}\n\n")

	f.WriteString("func (r " + className + ") Accept(visitor Visitor) any {\n")
	f.WriteString("\treturn visitor.Visit" + className + baseName + "(r)\n")
	f.WriteString("}\n\n")
}
