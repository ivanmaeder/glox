package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: generate <directory>")
		os.Exit(64)
	}

	defineAst(os.Args[1], "Expr", []string{
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

		f.WriteString("\tvisit" + typeName + baseName + "(" + strings.ToLower(baseName) + " *" + typeName + ") " + typeName + "\n")
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
		f.WriteString("\t" + strings.TrimSpace(element) + "\n")
	}
	f.WriteString("}\n\n")

	f.WriteString("func New" + className + "(" + strings.Join(fields, ",") + ") " + className + " {\n")
	f.WriteString("\treturn " + className + "{")
	comma := ""
	for _, element := range fields {
		fieldName := strings.Split(strings.TrimSpace(element), " ")[0]

		f.WriteString(comma + fieldName)
		comma = ", "
	}
	f.WriteString("}\n}\n\n")

	f.WriteString("func (r *" + className + ") accept(visitor Visitor) " + className + " {\n")
	f.WriteString("\treturn visitor.visit" + className + baseName + "(r)\n")
	f.WriteString("}\n\n")
}
