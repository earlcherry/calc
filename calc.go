package main

import (
	"os"
	"fmt"
	"go/parser"
	"strings"
	"go/ast"
	"go/token"
	"strconv"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("No expression provided.")
		os.Exit(1)
	}
	expression := strings.Join(os.Args[1:], " ")
	expr, err := parser.ParseExpr(expression)
	if err != nil {
		fmt.Println("Error parsing expression")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(eval(expr))
}

func eval(expr ast.Expr) float64 {
	switch exp := expr.(type) {
	case *ast.ParenExpr:
		return eval(exp.X)
	case *ast.BinaryExpr:
		return evalBinary(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return float64(i)
		case token.FLOAT:
			i, _ := strconv.ParseFloat(exp.Value, 64)
			return i
		}
	}
	return 0
}

func evalBinary(expr *ast.BinaryExpr) float64 {
	left := eval(expr.X)
	right := eval(expr.Y)
	switch expr.Op {
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left * right
	case token.QUO:
		return left / right
	case token.REM:
		return float64(int(left) % int(right))
	}
	return 0
}
