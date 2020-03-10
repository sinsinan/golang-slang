package lexer

import (
	"strconv"
)

//LexerType type definition for a lexer
type LexerType struct {
	iexpr  string
	length int
	Index  int
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
func GetToken(lexer *LexerType) (byte, float64, bool) {
	tok := TOK_ILLEGAL
	// fmt.Println(lexer)
	var floatVal float64 = 0
	ok := false
	for lexer.Index < lexer.length && lexer.iexpr[lexer.Index] == '\t' {
		lexer.Index++
	}

	if lexer.Index == lexer.length {
		tok = TOK_NULL
		ok = true
		return tok, floatVal, ok
	}

	switch lexer.iexpr[lexer.Index] {
	case TOK_PLUS, TOK_MUL, TOK_DIV, TOK_SUB, TOK_OPAREN, TOK_CPAREN:
		tok = lexer.iexpr[lexer.Index]
		lexer.Index++
		ok = true
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		var numberString string = string(lexer.iexpr[lexer.Index])
		lexer.Index++
		for lexer.Index < lexer.length && (lexer.iexpr[lexer.Index] == '0' || lexer.iexpr[lexer.Index] == '1' || lexer.iexpr[lexer.Index] == '2' || lexer.iexpr[lexer.Index] == '3' || lexer.iexpr[lexer.Index] == '4' || lexer.iexpr[lexer.Index] == '5' || lexer.iexpr[lexer.Index] == '6' || lexer.iexpr[lexer.Index] == '7' || lexer.iexpr[lexer.Index] == '8' || lexer.iexpr[lexer.Index] == '9') {
			numberString += string(lexer.iexpr[lexer.Index])
			lexer.Index++
		}
		if s, err := strconv.ParseFloat(numberString, 64); err == nil {
			floatVal = s
			ok = true
			tok = TOK_DOUBLE
		}
	}

	return tok, floatVal, ok
}
