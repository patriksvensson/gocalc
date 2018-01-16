package parser

import (
	"fmt"
)

type expression interface {
	accept(visitor) (int, error)
}
type visitor interface {
	visitInteger(*integerExpression) (int, error)
	visitArithmetic(*arithmeticExpression) (int, error)
}

/////////////////////////////////////
// Integer expression
/////////////////////////////////////

type integerExpression struct {
	value int
}

func (exp *integerExpression) accept(visitor visitor) (int, error) {
	return visitor.visitInteger(exp)
}

func (exp *integerExpression) String() string {
	return fmt.Sprintf("%d", exp.value)
}

/////////////////////////////////////
// Arithmetic expression
/////////////////////////////////////

type arithmeticExpression struct {
	operator tokenType
	left     expression
	right    expression
}

func (exp *arithmeticExpression) accept(visitor visitor) (int, error) {
	return visitor.visitArithmetic(exp)
}

func (exp *arithmeticExpression) String() string {
	return fmt.Sprintf("(%v %s %v)", exp.left, exp.operator.getString(), exp.right)
}
