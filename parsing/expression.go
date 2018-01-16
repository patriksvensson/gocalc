package parsing

import (
	"fmt"
)

type Expression interface {
	Accept(visitor *Visitor)
}
type Visitor interface {
	VisitInteger(*IntegerExpression)
	VisitArithmetic(*ArithmeticExpression)
}

/////////////////////////////////////
// Integer expression
/////////////////////////////////////

type IntegerExpression struct {
	value int
}

func (exp *IntegerExpression) Accept(visitor *Visitor) {
	(*visitor).VisitInteger(exp)
}

func (exp *IntegerExpression) String() string {
	return fmt.Sprintf("%d", exp.value)
}

/////////////////////////////////////
// Arithmetic expression
/////////////////////////////////////

type ArithmeticExpression struct {
	operator TokenType
	left     *Expression
	right    *Expression
}

func (exp *ArithmeticExpression) Accept(visitor *Visitor) {
	(*visitor).VisitArithmetic(exp)
}

func (exp *ArithmeticExpression) String() string {
	return fmt.Sprintf("(%v %s %v)", *exp.left, exp.operator.getString(), *exp.right)
}
