package builder

import (
	"github.com/sinsinan/golang-slang/step3/go/aststatement"
	"github.com/sinsinan/golang-slang/step3/go/rdparser"
)

//Run runs code given by the grammer
func Run(code string) {
	stmtList := rdparser.Parse(code)
	for _, stmt := range stmtList {
		aststatement.Execute(stmt)
	}
}
