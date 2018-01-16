package parser

import (
	"errors"
	"fmt"
	"strings"
	"text/scanner"
)

func tokenize(text string) ([]token, error) {

	var s scanner.Scanner
	s.Init(strings.NewReader(text))

	var tokens []token
	for t := s.Scan(); t != scanner.EOF; t = s.Scan() {
		value := s.TokenText()
		if t == scanner.Int {
			tokens = append(tokens, token{tokenType: number, value: value})
		} else if value == "+" {
			tokens = append(tokens, token{tokenType: plus, value: "+"})
		} else if value == "-" {
			tokens = append(tokens, token{tokenType: minus, value: "-"})
		} else if value == "(" {
			tokens = append(tokens, token{tokenType: openingParenthesis, value: "("})
		} else if value == ")" {
			tokens = append(tokens, token{tokenType: closingParenthesis, value: ")"})
		} else {
			return nil, fmt.Errorf("Unknown token '%s'.", value)
		}
	}

	return shuntingYard(&tokens)
}

func shuntingYard(tokens *[]token) ([]token, error) {

	var output []token
	var stack []token

	for _, token := range *tokens {
		if token.isOperand() {
			output = append(output, token)
		} else if token.isOperator() {
			for len(stack) > 0 {
				peek := stack[len(stack)-1]
				cond1 := peek.isOperator() && token.IsLeftAssociative() && token.GetPrecedence() <= peek.GetPrecedence()
				cond2 := peek.isOperator() && !token.IsLeftAssociative() && token.GetPrecedence() < peek.GetPrecedence()
				if cond1 || cond2 {
					// Current operator is left asociative and has a precedence less or equal than the operator on the stack,
					// OR the current operator is right associative and has a precedence less than the operator on the stack.
					output = append(output, peek)
					stack = stack[:len(stack)-1] // Pop
				} else {
					break
				}
			}
			// Push the current operator onto the stack.
			stack = append(stack, token)
		} else if token.tokenType == openingParenthesis {
			// Push the current operator onto the stack.
			stack = append(stack, token)
		} else if token.tokenType == closingParenthesis {
			foundLeftParenthesis := false
			for len(stack) > 0 {
				if stack[len(stack)-1].tokenType == openingParenthesis {
					foundLeftParenthesis = true
					break
				}
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1] // Pop
			}
			if !foundLeftParenthesis {
				return nil, errors.New("Missing left parenthesis in expression.")
			}
			stack = stack[:len(stack)-1] // Pop the left parenthesis from the stack.
		}
	}

	// Pop all operators from the stack
	for len(stack) > 0 {
		if stack[len(stack)-1].tokenType == openingParenthesis {
			return nil, errors.New("Missing right parenthesis in expression.")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1] // Pop
	}

	return output, nil
}
