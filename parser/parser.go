package parser

import (
	"errors"
	"github.com/patriksvensson/gocalc/collections"
	"strconv"
)

func parse(tokens *[]token) (expression, error) {
	stack := &collections.Stack{}
	for _, token := range *tokens {
		if token.tokenType == number {
			value, _ := strconv.Atoi(token.value)
			stack.Push(&integerExpression{value: value})
		} else {
			if !token.isOperator() {
				return nil, errors.New("Invalid expression. Expected operator.")
			}

			right := stack.Pop().(expression)
			left := stack.Pop().(expression)

			stack.Push(&arithmeticExpression{
				operator: token.tokenType,
				left:     left,
				right:    right})
		}
	}
	if stack.Count() != 1 {
		return nil, errors.New("Invalid expression. There should be exactly one expression on the stack.")
	}
	return stack.Pop().(expression), nil
}
