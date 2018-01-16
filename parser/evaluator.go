package parser

type evaluator struct {
	result int
}

func Evaluate(text string) (int, error) {
	tokens, err := tokenize(text)
	if err != nil {
		return 0, err
	}

	exp, err := parse(&tokens)
	if err != nil {
		return 0, err
	}

	return exp.accept(new(evaluator))
}

func (visitor *evaluator) visitInteger(exp *integerExpression) (int, error) {
	return exp.value, nil
}

func (visitor *evaluator) visitArithmetic(exp *arithmeticExpression) (int, error) {
	// Visit left side
	left, err := exp.left.accept(visitor)
	if err != nil {
		return 0, err
	}
	// Visit right side
	right, err := exp.right.accept(visitor)
	if err != nil {
		return 0, err
	}

	if exp.operator == plus {
		return left + right, nil
	} else if exp.operator == minus {
		return left - right, nil
	}

	panic("Unknown operator in arithmetic expression.")
}
