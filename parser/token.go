package parser

type tokenType int
type token struct {
	tokenType tokenType
	value     string
}

const (
	number tokenType = iota
	plus
	minus
	openingParenthesis
	closingParenthesis
)

func (t *tokenType) getString() string {
	switch ty := t; *ty {
	case plus:
		return "+"
	case minus:
		return "-"
	}
	return ""
}

func (t *token) isOperand() bool {
	return t.tokenType == number
}

func (t *token) isOperator() bool {
	switch ty := t.tokenType; ty {
	case plus:
		return true
	case minus:
		return true
	}
	return false
}

func (t *token) IsLeftAssociative() bool {
	return false
}

func (t *token) GetPrecedence() int {
	switch ty := t.tokenType; ty {
	case plus:
		return 1
	case minus:
		return 1
	}
	return 0
}
