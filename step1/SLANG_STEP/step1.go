package SLANG_STEP_1

//Operators these are the valid operators
type Operators string

const (
	//PLUS + operator
	PLUS Operators = "+"
	//MINUS - operator
	MINUS Operators = "-"
	//DIV / operator
	DIV Operators = "/"
	//MUL * operator
	MUL Operators = "*"
	//INVALID invalid operator
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
	ex Exp
	op Operators
}

//NumericConstant returns a NumericConstantType type with given value
func NumericConstant(value float64) *NumericConstantType {
	nc := NumericConstantType{value}
	return &nc
}

//BinaryExp returns a BinaryExpType type with given values
func BinaryExp(ex1 Exp, ex2 Exp, op Operators) *BinaryExpType {
	be := BinaryExpType{ex1, ex2, op}
	return &be
}

//UnaryExp returns a UnaryExpType type with given values
func UnaryExp(ex Exp, op Operators) *UnaryExpType {
	ue := UnaryExpType{ex, op}
	return &ue
}

//EvaluateNumericConstant returns the (float64)value in the given NumericConstantType
func EvaluateNumericConstant(nc NumericConstantType) (float64, bool) {
	return nc.value, false
}

//EvaluateBinaryExp evaluates the given binary expression
func EvaluateBinaryExp(be BinaryExpType) (float64, bool) {
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

//EvaluateUnaryExp evaluates the given unary expression
func EvaluateUnaryExp(ue UnaryExpType) (float64, bool) {
	switch ue.op {
	case PLUS:
		return ue.ex.Evaluate(), false
	case MINUS:
		return -ue.ex.Evaluate(), false
	default:
		return -1, true
	}
}

//Evaluate returns the (float64)value in the given NumericConstantType, BinaryExpType, UnaryExpType
func Evaluate(exp interface{}) (float64, bool) {
	switch exp.(type) {
	case NumericConstantType:
		if nc, ok :=
			exp.(NumericConstantType); ok {
			return EvaluateNumericConstant(nc)
		}

	case BinaryExpType:
		if be, ok :=
			exp.(BinaryExpType); ok {
			return EvaluateBinaryExp(be)
		}

	case UnaryExpType:
		if ue, ok :=
			exp.(UnaryExpType); ok {
			return EvaluateUnaryExp(ue)
		}

	}

	return 0, true
}
