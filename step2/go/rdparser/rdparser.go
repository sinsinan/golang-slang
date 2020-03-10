package rdparser

import (
	"github.com/sinsinan/golang-slang/step2/go/ast"
	"github.com/sinsinan/golang-slang/step2/go/lexer"
)

//LANGOAGE GRAMMER
// E -> T | T{+|-}E
// T -> F | F{*|/}T
// F -> n | (E) | {+|-}F

//ParserType a type which point current parser position
type ParserType struct {
	currentToken byte
	floatVal     float64
	ok           bool
}

//CallExpr used to get a ast expression from a given string
func CallExpr(expression string) (ast.Exp, bool) {
	currentLexer := lexer.Lexer(expression)
	parsor := ParserType{lexer.TOK_ILLEGAL, 0, false}
	return Expr(currentLexer, &parsor)
}

func getTokenAndParse(currentLexer *lexer.LexerType, parser *ParserType) {
	currentToken, floatVal, ok := lexer.GetToken(currentLexer)
	parser.currentToken = currentToken
	parser.floatVal = floatVal
	parser.ok = ok
}

//Expr evaluates expression
func Expr(currentLexer *lexer.LexerType, parser *ParserType) (ast.Exp, bool) {
	T, ok := Term(currentLexer, parser)
	getTokenAndParse(currentLexer, parser)
	if parser.currentToken == ast.PLUS || parser.currentToken == ast.MINUS {
		operator := parser.currentToken
		E, ok := Expr(currentLexer, parser)
		return ast.BinaryExp(T, E, operator), ok
	}
	currentLexer.Index--
	return T, ok
}

//Term evaluates a term
func Term(currentLexer *lexer.LexerType, parser *ParserType) (ast.Exp, bool) {
	F, ok := Factor(currentLexer, parser)
	getTokenAndParse(currentLexer, parser)
	if parser.currentToken == ast.DIV || parser.currentToken == ast.MUL {
		operator := parser.currentToken
		T, ok := Term(currentLexer, parser)
		return ast.BinaryExp(F, T, operator), ok
	}
	currentLexer.Index--
	return F, ok
}

//Factor evaluates a factor
func Factor(currentLexer *lexer.LexerType, parser *ParserType) (ast.Exp, bool) {
	getTokenAndParse(currentLexer, parser)
	switch parser.currentToken {
	case lexer.TOK_DOUBLE:
		return ast.NumericConstant(parser.floatVal), parser.ok
	case lexer.TOK_OPAREN:
		E, ok := Expr(currentLexer, parser)
		getTokenAndParse(currentLexer, parser)
		if parser.currentToken == lexer.TOK_CPAREN {
			return E, ok
		}
	case ast.PLUS, ast.MINUS:
		operator := parser.currentToken
		F, ok := Factor(currentLexer, parser)
		return ast.UnaryExp(F, operator), ok
	}
	return ast.NumericConstant(0), false
}
