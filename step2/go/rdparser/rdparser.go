package rdparser

import (
	"github.com/sinsinan/golang-slang/step2/go/ast"
	"github.com/sinsinan/golang-slang/step2/go/lexer"
)

//LANGOAGE GRAMMER
// E -> T | T{+|-}E
// T -> F | F{*|/}T
// F -> n | (E) | {+|-}F

type parserType struct {
	currentToken byte
	floatVal     float64
	ok           bool
}

func getTokenAndParse(currentLexer lexer.LexerType) *parserType {
	currentToken, floatVal, ok := lexer.GetToken(currentLexer)
	return &parserType{currentToken, floatVal, ok}
}

//Expr evaluates expression
func Expr(currentLexer lexer.LexerType, parser parserType) (ast.Exp, bool) {
	T, ok := Term(currentLexer, parser)
	parser = *getTokenAndParse(currentLexer)
	if parser.currentToken == ast.PLUS || parser.currentToken == ast.MINUS {
		operator := parser.currentToken
		E, ok := Expr(currentLexer, parser)
		return ast.BinaryExp(T, E, operator), ok
	}
	return T, ok
}

//Term evaluates a term
func Term(currentLexer lexer.LexerType, parser parserType) (ast.Exp, bool) {
	F, ok := Factor(currentLexer, parser)
	parser = *getTokenAndParse(currentLexer)
	if parser.currentToken == ast.DIV || parser.currentToken == ast.MUL {
		operator := parser.currentToken
		T, ok := Term(currentLexer, parser)
		return ast.BinaryExp(F, T, operator), ok
	}
	return F, ok
}

//Factor evaluates a factor
func Factor(currentLexer lexer.LexerType, parser parserType) (ast.Exp, bool) {
	parser = *getTokenAndParse(currentLexer)
	switch parser.currentToken {
	case lexer.TOK_DOUBLE:
		return ast.NumericConstant(parser.floatVal), parser.ok
	case lexer.TOK_OPAREN:
		E, ok := Expr(currentLexer, parser)
		parser = *getTokenAndParse(currentLexer)
		if parser.currentToken == lexer.TOK_CPAREN {
			return E, ok
		}
	case ast.PLUS, ast.MINUS:
		F, ok := Factor(currentLexer, parser)
		return F, ok
	}
	return ast.NumericConstant(0), false
}
