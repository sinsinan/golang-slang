package main

import (
	"fmt"

	"github.com/sinsinan/golang-slang/step2/go/ast"
	"github.com/sinsinan/golang-slang/step2/go/rdparser"
)

func main() {
	if Exp, ok := rdparser.CallExpr("-2*(3+3)"); ok {
		if value, ok := ast.Evaluate(Exp); ok {
			fmt.Println(value)
		}
	}

}
