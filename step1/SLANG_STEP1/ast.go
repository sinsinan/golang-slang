package ast

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
	return nc.value, true
}

//EvaluateBinaryExp evaluates the given binary expression
func EvaluateBinaryExp(be BinaryExpType) (float64, bool) {
	switch be.op {
	case PLUS:
		value1, ok1 := Evaluate(be.ex1)
		value2, ok2 := Evaluate(be.ex2)
		if ok1 && ok2 {
			return value1 + value2, true
		}
	case MINUS:
		value1, ok1 := Evaluate(be.ex1)
		value2, ok2 := Evaluate(be.ex2)
		if ok1 && ok2 {
			return value1 - value2, true
		}
	case DIV:
		value1, ok1 := Evaluate(be.ex1)
		value2, ok2 := Evaluate(be.ex2)
		if ok1 && ok2 {
			return value1 / value2, true
		}
	case MUL:
		value1, ok1 := Evaluate(be.ex1)
		value2, ok2 := Evaluate(be.ex2)
		if ok1 && ok2 {
			return value1 * value2, true
		}
	}
	return -1, false
}

//EvaluateUnaryExp evaluates the given unary expression
func EvaluateUnaryExp(ue UnaryExpType) (float64, bool) {
	switch ue.op {
	case PLUS:
		if value, ok := Evaluate(ue.ex); ok {
			return value, ok
		}
	case MINUS:
		if value, ok := Evaluate(ue.ex); ok {
			return -value, ok
		}
	}

	return -1, false
}

//Evaluate returns the (float64)value in the given NumericConstantType, BinaryExpType, UnaryExpType
func Evaluate(exp interface{}) (float64, bool) {
	switch exp.(type) {
	case *NumericConstantType:
		if nc, ok :=
			exp.(*NumericConstantType); ok {
			return EvaluateNumericConstant(*nc)
		}

	case *BinaryExpType:
		if be, ok :=
			exp.(*BinaryExpType); ok {
			return EvaluateBinaryExp(*be)
		}

	case *UnaryExpType:

		if ue, ok :=
			exp.(*UnaryExpType); ok {
			return EvaluateUnaryExp(*ue)
		}

	}

	return -1, false
}
