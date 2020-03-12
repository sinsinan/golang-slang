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

const (
	TOK_ILLEGAL         string = "INVALID"
	TOK_PLUS            string = "+"
	TOK_MUL             string = "*"
	TOK_DIV             string = "/"
	TOK_SUB             string = "-"
	TOK_OPAREN          string = "("
	TOK_CPAREN          string = ")"
	TOK_DOUBLE          string = "DOUBLE"
	TOK_NULL            string = "N"
	TOK_PRINT           string = "PRINT"
	TOK_UNQUOTED_STRING string = "U"
	TOK_PRINT_LINE      string = "PRINTLINE"
	TOK_SEMI            string = ";"
)

//Lexer creates and returns a LexerType
func Lexer(exp string) *LexerType {
	lexer := LexerType{exp, len(exp), 0}
	return &lexer
}

//GetToken is used get token from a given lexer
func GetToken(lexer *LexerType) (string, float64, bool) {
	tok := TOK_ILLEGAL
	// fmt.Println(lexer)
	var floatVal float64 = 0
	ok := false
	for lexer.Index < lexer.length && (lexer.iexpr[lexer.Index] == '\t' || lexer.iexpr[lexer.Index] == ' ') {
		lexer.Index++
	}

	if lexer.Index == lexer.length {
		tok = TOK_NULL
		ok = true
		return tok, floatVal, ok
	}

	switch string(lexer.iexpr[lexer.Index]) {
	case TOK_PLUS, TOK_MUL, TOK_DIV, TOK_SUB, TOK_OPAREN, TOK_CPAREN, TOK_SEMI:
		tok = string(lexer.iexpr[lexer.Index])
		lexer.Index++
		ok = true
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
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
	case "P":
		if lexer.Index+9 < lexer.length && lexer.iexpr[lexer.Index:lexer.Index+9] == TOK_PRINT_LINE {
			lexer.Index += 9
			tok = TOK_PRINT_LINE
			ok = true
		} else if lexer.Index+5 < lexer.length && lexer.iexpr[lexer.Index:lexer.Index+5] == TOK_PRINT {
			lexer.Index += 5
			tok = TOK_PRINT
			ok = true
		}

	}

	return tok, floatVal, ok
}
