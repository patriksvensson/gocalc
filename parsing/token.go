package parsing

type TokenType int
type Token struct {
	Type  TokenType
	Value string
}

const (
	Number TokenType = iota
	Plus
	Minus
	OpeningParenthesis
	ClosingParenthesis
)

func (t *TokenType) getString() string {
	switch ty := t; *ty {
	case Plus:
		return "+"
	case Minus:
		return "-"
	}
	return ""
}

func (t *Token) isOperand() bool {
	return t.Type == Number
}

func (t *Token) isOperator() bool {
	switch ty := t.Type; ty {
	case Plus:
		return true
	case Minus:
		return true
	}
	return false
}

func (t *Token) IsLeftAssociative() bool {
	return false
}

func (t *Token) GetPrecedence() int {
	switch ty := t.Type; ty {
	case Plus:
		return 1
	case Minus:
		return 1
	}
	return 0
}
