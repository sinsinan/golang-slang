package lexer

import (
	"strconv"
)

//LexerType type definition for a lexer
type LexerType struct {
	iexpr  string
	length int
	index  int
}

type Token byte

const (
	TOK_ILLEGAL byte = 'I'
	TOK_PLUS    byte = '+'
	TOK_MUL     byte = '*'
	TOK_DIV     byte = '/'
	TOK_SUB     byte = '-'
	TOK_OPAREN  byte = '('
	TOK_CPAREN  byte = ')'
	TOK_DOUBLE  byte = 'D'
	TOK_NULL    byte = 'N'
)

//Lexer creates and returns a LexerType
func Lexer(exp string) *LexerType {
	lexer := LexerType{exp, len(exp), 0}
	return &lexer
}

//GetToken is used get token from a given lexer
func GetToken(lexer LexerType) (string, float64, bool) {
	tok := TOK_ILLEGAL
	var floatVal float64 = 0
	ok := false
	for lexer.index < lexer.length && lexer.iexpr[lexer.index] == '\t' {
		lexer.index++
	}

	if lexer.index == lexer.length {
		tok = TOK_NULL
		ok = true
	}

	switch lexer.iexpr[lexer.index] {
	case TOK_PLUS, TOK_MUL, TOK_DIV, TOK_SUB, TOK_OPAREN, TOK_CPAREN:
		lexer.index++
		return string(lexer.iexpr[lexer.index]), 0, true
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		lexer.index++
		var numberString string = string(lexer.iexpr[lexer.index])
		for lexer.index < lexer.length && (lexer.iexpr[lexer.index] == '0' || lexer.iexpr[lexer.index] == '1' || lexer.iexpr[lexer.index] == '2' || lexer.iexpr[lexer.index] == '3' || lexer.iexpr[lexer.index] == '4' || lexer.iexpr[lexer.index] == '5' || lexer.iexpr[lexer.index] == '6' || lexer.iexpr[lexer.index] == '7' || lexer.iexpr[lexer.index] == '8' || lexer.iexpr[lexer.index] == '9') {
			numberString += string(lexer.iexpr[lexer.index])
			lexer.index++
		}
		if s, err := strconv.ParseFloat(numberString, 64); err == nil {
			floatVal = s
			ok = true
			tok = TOK_DOUBLE
		}
	}

	return string(tok), floatVal, ok
}
