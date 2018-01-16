package parser

import (
	"errors"
	"strconv"
)

func parse(tokens *[]token) (expression, error) {
	var stack []expression
	for _, token := range *tokens {
		if token.tokenType == number {
			value, _ := strconv.Atoi(token.value)
			stack = append(stack, &integerExpression{value: value})
		} else {
			if !token.isOperator() {
				return nil, errors.New("Invalid expression. Expected operator.")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop
			stack = append(stack, &arithmeticExpression{operator: token.tokenType, left: left, right: right})
		}
	}
	if len(stack) != 1 {
		return nil, errors.New("Invalid expression. There should be exactly one expression on the stack.")
	}
	return stack[len(stack)-1], nil
}
