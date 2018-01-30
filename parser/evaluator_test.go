package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThat_ExpressionIsEvaluatedCorrectly(t *testing.T) {
	// Given, When
	result, _ := Evaluate("3+1-4+1")
	// Then
	assert.Equal(t, 1, result)
}

func TestThat_ExpressionWithMissingRightParenthesisReturnsError(t *testing.T) {
	// Given, When
	_, error := Evaluate("(1+1")
	// Then
	assert.NotNil(t, error, "Expected an error but got none.")
	assert.Equal(t, "Missing right parenthesis in expression.", error.Error())
}

func TestThat_ExpressionWithMissingLeftParenthesisReturnsError(t *testing.T) {
	// Given, When
	_, error := Evaluate("1+1)")
	// Then
	assert.NotNil(t, error, "Expected an error but got none.")
	assert.Equal(t, "Missing left parenthesis in expression.", error.Error())
}

func TestThat_ExpressionWithUnknownTokenReturnsError(t *testing.T) {
	// Given, When
	_, error := Evaluate("1*1)")
	// Then
	assert.NotNil(t, error, "Expected an error but got none.")
	assert.Equal(t, "Unknown token '*'.", error.Error())
}

func TestThat_ExpressionWithParenthesesIsEvaluatedCorrectly(t *testing.T) {
	// Given, When
	result, _ := Evaluate("3+(1-4)+1")
	// Then
	assert.Equal(t, 1, result)
}
