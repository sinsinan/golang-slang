package main

import (
	"fmt"

	ast "github.com/sinsinan/golang-slang/step1/SLANG_STEP1"
)

func main() {
	exp := ast.BinaryExp(ast.NumericConstant(10), ast.NumericConstant(20), ast.MINUS)
	if value, ok := ast.Evaluate(exp); ok {
		fmt.Println(value)
	}
}
