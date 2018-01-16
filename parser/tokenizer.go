package parser

import (
	"errors"
	"fmt"
	"github.com/patriksvensson/gocalc/collections"
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

	output := &collections.Queue{}
	stack := &collections.Stack{}

	for _, t := range *tokens {
		if t.isOperand() {
			output.Enqueue(t)
		} else if t.isOperator() {
			for stack.Count() > 0 {
				peek := stack.Peek().(token)
				cond1 := peek.isOperator() && t.IsLeftAssociative() && t.GetPrecedence() <= peek.GetPrecedence()
				cond2 := peek.isOperator() && !t.IsLeftAssociative() && t.GetPrecedence() < peek.GetPrecedence()
				if cond1 || cond2 {
					// Current operator is left asociative and has a precedence less or equal than the operator on the stack,
					// OR the current operator is right associative and has a precedence less than the operator on the stack.
					output.Enqueue(stack.Pop())
				} else {
					break
				}
			}
			// Push the current operator onto the stack.
			stack.Push(t)
		} else if t.tokenType == openingParenthesis {
			// Push the current operator onto the stack.
			stack.Push(t)
		} else if t.tokenType == closingParenthesis {
			foundLeftParenthesis := false
			for stack.Count() > 0 {
				if stack.Peek().(token).tokenType == openingParenthesis {
					foundLeftParenthesis = true
					break
				}
				output.Enqueue(stack.Pop())
			}
			if !foundLeftParenthesis {
				return nil, errors.New("Missing left parenthesis in expression.")
			}
			// Pop the left parenthesis from the stack.
			stack.Pop()
		}
	}

	// Pop all operators from the stack
	for stack.Count() > 0 {
		if stack.Peek().(token).tokenType == openingParenthesis {
			return nil, errors.New("Missing right parenthesis in expression.")
		}
		output.Enqueue(stack.Pop())
	}

	// Convert the output queue an output slice.
	result := make([]token, output.Count())
	for index, t := range output.DequeueAll() {
		result[index] = t.(token)
	}

	return result, nil
}
