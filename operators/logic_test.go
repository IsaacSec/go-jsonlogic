package operators

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
)

/***********************
 *		EQUALS
 ***********************/

// EQUALS evaluation, same type equal value -> 53 == 53
func TestEqualsOnSameTypeAndValue(t *testing.T) {
	expression := buildSimpleExp("==", 53, 53)

	assertExpression(t, expression, equalsOperator, token.True)
}

// EQUALS evaluation, different type equal value -> 53 == "53"
func TestEqualsOnDiffType(t *testing.T) {
	expression := buildSimpleExp("==", 53, "fifty three")

	assertExpression(t, expression, equalsOperator, token.False)
}

// EQUALS evaluation, different value -> 3421 == -123
func TestEqualsOnDiffValue(t *testing.T) {
	expression := buildSimpleExp("==", 3421, -123)

	assertExpression(t, expression, equalsOperator, token.False)
}

// Todo: Add conversion type when possible i.e. -> 1 == "1" == 1.0

// EQUALS evaluation, different type and equal value -> 42 == 42.0
func TestEqualsOnDiffTypeWithSameValue(t *testing.T) {
	expression := buildSimpleExp("==", 42, 42.0)

	assertExpression(t, expression, equalsOperator, token.False)
}

/***********************
 *		NOT EQUALS
 ***********************/

// NOT EQUALS evaluation, same type and different value -> 35 != 53
func TestNotEqualsWithDiffValue(t *testing.T) {
	expression := buildSimpleExp("!=", 35, 53)

	assertExpression(t, expression, notEqualsEvaluator, token.True)
}

// NOT EQUALS evaluation, same type and equal value -> 53 != 53
func TestNotEqualsWithSameValue(t *testing.T) {
	expression := buildSimpleExp("!=", 53, 53)

	assertExpression(t, expression, notEqualsEvaluator, token.False)
}

/***********************
 *		AND
 ***********************/

// AND evaluation, empty list -> [] = true
func TestAndWithEmptyList(t *testing.T) {
	expression := buildGroupExp(
		"and",
	)

	assertExpression(t, expression, andOperator, token.True)
}

// AND evaluation, one true and multiple false -> [true, false, false] = false
func TestAndWithMultipleFalseAndAtLeastOneTrue(t *testing.T) {
	expression := buildGroupExp(
		"and",
		token.True,
		token.False,
		token.False,
	)

	assertExpression(t, expression, andOperator, token.False)
}

// AND evaluation, all true -> [true, true, true] = true
func TestAndWithAllTrue(t *testing.T) {
	expression := buildGroupExp(
		"and",
		token.True,
		token.True,
		token.True,
	)

	assertExpression(t, expression, andOperator, token.True)
}

// AND evaluation, all false -> [false, false, false] = false
func TestAndWithAllFalse(t *testing.T) {
	expression := buildGroupExp(
		"and",
		token.False,
		token.False,
		token.False,
	)

	assertExpression(t, expression, andOperator, token.False)
}

/***********************
 *		OR
 ***********************/

// OR evaluation, empty list -> [] = true
func TestOrWithEmptyList(t *testing.T) {
	expression := buildGroupExp(
		"or",
	)

	assertExpression(t, expression, orEvaluator, token.True)
}

// OR evaluation, one true and multiple false -> [false, true, false] = true
func TestOrWithMultipleFalseAndAtLeastOneTrue(t *testing.T) {
	expression := buildGroupExp(
		"or",
		token.False,
		token.True,
		token.False,
	)

	assertExpression(t, expression, orEvaluator, token.True)
}

// AND evaluation, all true -> [true, true, true] = true
func TestOrWithAllTrue(t *testing.T) {
	expression := buildGroupExp(
		"or",
		token.True,
		token.True,
		token.True,
	)

	assertExpression(t, expression, orEvaluator, token.True)
}

// AND evaluation, all false -> [false, false, false] = false
func TestOrWithAllFalse(t *testing.T) {
	expression := buildGroupExp(
		"or",
		token.False,
		token.False,
		token.False,
	)

	assertExpression(t, expression, orEvaluator, token.False)
}

func buildGroupExp(op token.Token, results ...token.EvalResult) token.Node {
	var expressions = make([]*token.Node, len(results))

	for i, res := range results {
		expressions[i] = &token.Node{
			Token:           res.ToString() == "True",
			CommulativeEval: res,
			Kind:            token.PrimitiveVal,
			Childrens:       nil,
		}
	}

	var group = token.Node{Token: op, Kind: token.Operator, Childrens: expressions}

	return group
}

func buildSimpleExp(op token.Token, a token.Token, b token.Token) token.Node {
	return token.Node{
		Token: op,
		Kind:  token.Operator,
		Childrens: []*token.Node{
			{
				Token: a,
				Kind:  token.PrimitiveVal,
			},
			{
				Token: b,
				Kind:  token.PrimitiveVal,
			},
		},
	}
}

func assertExpression(t *testing.T, exp token.Node, evaluator OperatorRunnable, expected token.EvalResult) {
	res := evaluator.Evaluate(&exp)

	if res != expected {
		t.Errorf("Expected %v but got '%v'", expected.ToString(), res.ToString())
	}
}
