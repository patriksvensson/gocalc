package parsing

import (
	"errors"
	"fmt"
	"strings"
	"text/scanner"
)

func Tokenize(text string) ([]Token, error) {

	var s scanner.Scanner
	s.Init(strings.NewReader(text))

	var tokens []Token
	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		value := s.TokenText()
		if token == scanner.Int {
			tokens = append(tokens, Token{Type: Number, Value: value})
		} else if value == "+" {
			tokens = append(tokens, Token{Type: Plus, Value: "+"})
		} else if value == "-" {
			tokens = append(tokens, Token{Type: Minus, Value: "-"})
		} else if value == "(" {
			tokens = append(tokens, Token{Type: OpeningParenthesis, Value: "("})
		} else if value == ")" {
			tokens = append(tokens, Token{Type: ClosingParenthesis, Value: ")"})
		} else {
			return nil, errors.New(fmt.Sprintf("Unknown token '%s'", value))
		}
	}

	return shuntingYard(&tokens), nil
}

func shuntingYard(tokens *[]Token) []Token {

	output := make([]Token, 0)
	stack := make([]Token, 0)

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
		} else if token.Type == OpeningParenthesis {
			// Push the current operator onto the stack.
			stack = append(stack, token)
		} else if token.Type == ClosingParenthesis {
			foundLeftParenthesis := false
			for len(stack) > 0 {
				if stack[len(stack)-1].Type == OpeningParenthesis {
					foundLeftParenthesis = true
					break
				}
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1] // Pop
			}
			if !foundLeftParenthesis {
				panic("Missing left parenthesis in expression.")
			}
			stack = stack[:len(stack)-1] // Pop the left parenthesis from the stack.
		}
	}

	// Pop all operators from the stack
	for len(stack) > 0 {
		if stack[len(stack)-1].Type == OpeningParenthesis {
			panic("Missing right parenthesis in expression.")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1] // Pop
	}

	return output
}
