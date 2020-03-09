package step1

//Operators these are the valid operators
type Operators string

const (
	PLUS    Operators = "+"
	MINUS   Operators = "-"
	DIV     Operators = "/"
	MUL     Operators = "*"
	INVALID Operators = "-1"
)

//Exp is an interface with a function Evaluate which returns a float64
type Exp interface {
	Evaluate() float64
}

//NumericConstantType is the type for numericconstant
type NumericConstantType struct {
	value float64
}

//BinaryExpType is the type for binary expressions
type BinaryExpType struct {
	ex1 Exp
	ex2 Exp
	op  Operators
}

//UnaryExpType is the type for unary expressions
type UnaryExpType struct {
	ex1 Exp
	ex2 Exp
	op  Operators
}

//NumericConstant returns a NumericConstantType type with given value
func NumericConstant(value float64) *NumericConstantType {
	nc := NumericConstantType{value}
	return &nc
}

//EvaluateNumericConstant returns the (float64)value in the given NumericConstantType
func EvaluateNumericConstant(nc NumericConstantType) (float64, bool) {
	return nc.value, false
}

//EvaluateBinaryExpType evaluates the given binary expression
func EvaluateBinaryExpType(be BinaryExpType) (float64, bool) {
	switch be.op {
	case PLUS:
		return (be.ex1.Evaluate() + be.ex2.Evaluate()), false
	case MINUS:
		return (be.ex1.Evaluate() - be.ex2.Evaluate()), false
	case DIV:
		return (be.ex1.Evaluate() / be.ex2.Evaluate()), false
	case MUL:
		return (be.ex1.Evaluate() * be.ex2.Evaluate()), false
	default:
		return -1, true
	}
}

//Evaluate returns the (float64)value in the given NumericConstantType, BinaryExpType, UnaryExpType
func Evaluate(exp interface{}) (float64, bool) {
	switch exp.(type) {
	case NumericConstantType:
		nc, ok :=
			exp.(NumericConstantType)
		return 0, EvaluateNumericConstant(nc)
	case BinaryExpType:
		return 0, exp.EvaluateBinaryExpType()
	case UnaryExpType:
		return 0, exp.EvaluateUnaryExpType()
	default:
		return 0, true
	}
}
