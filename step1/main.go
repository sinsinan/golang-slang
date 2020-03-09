package main

import (
	"fmt"

	"github.com/sinsinan/golang-slang/step1/SLANG_STEP1"
)

func main() {
	exp := SLANG_STEP1.BinaryExp(SLANG_STEP1.NumericConstant(10), SLANG_STEP1.NumericConstant(20), SLANG_STEP1.PLUS)
	if value, ok := SLANG_STEP1.Evaluate(exp); ok {
		fmt.Println(value)
	}
}
