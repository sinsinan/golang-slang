package aststatement

import (
	"fmt"

	"github.com/sinsinan/golang-slang/step3/go/ast"
)

//Stmt is an interface for all statemnts
type Stmt interface {
}

//PrintType type used for print statements
type PrintType struct {
	ex ast.Exp
}

//PrintLineType type used for print statements
type PrintLineType struct {
	ex ast.Exp
}

//Print constructor for print
func Print(exp ast.Exp) *PrintType {
	print := PrintType{exp}
	return &print
}

//PrintLine constructor for PrintLine
func PrintLine(exp ast.Exp) *PrintLineType {
	print := PrintLineType{exp}
	return &print
}

//ExecutePrint function to execute print
func ExecutePrint(pt PrintType) bool {
	if value, ok := ast.Evaluate(pt.ex); ok {
		fmt.Print(value)
		return true
	}
	return false
}

//ExecutePrintLine function to execute print
func ExecutePrintLine(plt PrintLineType) bool {
	if value, ok := ast.Evaluate(plt.ex); ok {
		fmt.Print(value, "\n")
		return true
	}
	return false
}

//Execute the function dealing with execution of particular statements
func Execute(stmt Stmt) bool {
	switch stmt.(type) {
	case *PrintType:
		if pt, ok :=
			stmt.(*PrintType); ok {
			return ExecutePrint(*pt)
		}
	case *PrintLineType:
		if plt, ok :=
			stmt.(*PrintLineType); ok {
			return ExecutePrintLine(*plt)
		}
	}
	return false
}
