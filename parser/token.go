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

func (t tokenType) getString() string {
	if t == plus {
		return "+"
	} else if t == minus {
		return "-"
	}
	return ""
}

func (t *token) isOperand() bool {
	return t.tokenType == number
}

func (t *token) isOperator() bool {
	if t.tokenType == plus {
		return true
	} else if t.tokenType == minus {
		return true
	}
	return false
}

func (t *token) IsLeftAssociative() bool {
	return false
}

func (t *token) GetPrecedence() int {
	if t.tokenType == plus {
		return 1
	} else if t.tokenType == minus {
		return 1
	}
	return 0
}
