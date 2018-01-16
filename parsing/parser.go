package parsing

import "strconv"

func Parse(tokens *[]Token) Expression {
	stack := make([]Expression, 0)
	for _, token := range *tokens {
		if token.Type == Number {
			value, _ := strconv.Atoi(token.Value)
			stack = append(stack, &IntegerExpression{value: value})
		} else {
			if !token.isOperator() {
				panic("Invalid expression. Expected operator.")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop
			stack = append(stack, &ArithmeticExpression{operator: token.Type, left: &left, right: &right})
		}
	}
	if len(stack) != 1 {
		panic("Invalid expression. There should be exactly one expression on the stack.")
	}
	return stack[len(stack)-1]
}
