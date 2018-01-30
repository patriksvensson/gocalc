package parser

import "testing"

func TestThatExpressionIsEvaluatedCorrectly(t *testing.T) {
	// Given, When
	result, _ := Evaluate("3+1-4+1")
	// Then
	if result != 1 {
		t.Error("Expected 1, got", result)
	}
}

func TestThatExpressionWithParenthesesIsEvaluatedCorrectly(t *testing.T) {
	// Given, When
	result, _ := Evaluate("3+(1-4)+1)")
	// Then
	if result != 1 {
		t.Error("Expected 1, got", result)
	}
}
