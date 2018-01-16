package parsing

type TokenType int
type Token struct {
	Type  TokenType
	Value string
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

const (
	Number             TokenType = 0
	Plus                         = 1
	Minus                        = 2
	OpeningParenthesis           = 3
	ClosingParenthesis           = 4
)
