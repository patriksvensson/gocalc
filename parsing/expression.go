package parsing

type Expression interface {
	Accept(*Visitor)
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

func (exp *IntegerExpression) Visit(visitor *Visitor) {
	(*visitor).VisitInteger(exp)
}

/////////////////////////////////////
// Arithmetic expression
/////////////////////////////////////

type ArithmeticExpression struct {
	operator TokenType
}

func (exp *ArithmeticExpression) Visit(visitor *Visitor) {
	(*visitor).VisitArithmetic(exp)
}
