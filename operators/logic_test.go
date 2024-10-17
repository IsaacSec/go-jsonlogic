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

	assertExpression(t, expression, equalsOperator, true)
}

// EQUALS evaluation, different type equal value -> 53 == "53"
func TestEqualsOnDiffType(t *testing.T) {
	expression := buildSimpleExp("==", 53, "fifty three")

	assertExpression(t, expression, equalsOperator, false)
}

// EQUALS evaluation, different value -> 3421 == -123
func TestEqualsOnDiffValue(t *testing.T) {
	expression := buildSimpleExp("==", 3421, -123)

	assertExpression(t, expression, equalsOperator, false)
}

// Todo: Add conversion type when possible i.e. -> 1 == "1" == 1.0

// EQUALS evaluation, different type and equal value -> 42 == 42.0
func TestEqualsOnDiffTypeWithSameValue(t *testing.T) {
	expression := buildSimpleExp("==", 42, 42.0)

	assertExpression(t, expression, equalsOperator, false)
}

/***********************
 *		NOT EQUALS
 ***********************/

// NOT EQUALS evaluation, same type and different value -> 35 != 53
func TestNotEqualsWithDiffValue(t *testing.T) {
	expression := buildSimpleExp("!=", 35, 53)

	assertExpression(t, expression, notEqualsEvaluator, true)
}

// NOT EQUALS evaluation, same type and equal value -> 53 != 53
func TestNotEqualsWithSameValue(t *testing.T) {
	expression := buildSimpleExp("!=", 53, 53)

	assertExpression(t, expression, notEqualsEvaluator, false)
}

/***********************
 *		AND
 ***********************/

// AND evaluation, empty list -> [] = true
func TestAndWithEmptyList(t *testing.T) {
	expression := buildGroupExp(
		"and",
	)

	assertExpression(t, expression, andOperator, true)
}

// AND evaluation, one true and multiple false -> [true, false, false] = false
func TestAndWithMultipleFalseAndAtLeastOneTrue(t *testing.T) {
	expression := buildGroupExp(
		"and",
		true,
		false,
		false,
	)

	assertExpression(t, expression, andOperator, false)
}

// AND evaluation, all true -> [true, true, true] = true
func TestAndWithAllTrue(t *testing.T) {
	expression := buildGroupExp(
		"and",
		true,
		true,
		true,
	)

	assertExpression(t, expression, andOperator, true)
}

// AND evaluation, all false -> [false, false, false] = false
func TestAndWithAllFalse(t *testing.T) {
	expression := buildGroupExp(
		"and",
		false,
		false,
		false,
	)

	assertExpression(t, expression, andOperator, false)
}

/***********************
 *		OR
 ***********************/

// OR evaluation, empty list -> [] = true
func TestOrWithEmptyList(t *testing.T) {
	expression := buildGroupExp(
		"or",
	)

	assertExpression(t, expression, orEvaluator, true)
}

// OR evaluation, one true and multiple false -> [false, true, false] = true
func TestOrWithMultipleFalseAndAtLeastOneTrue(t *testing.T) {
	expression := buildGroupExp(
		"or",
		false,
		true,
		false,
	)

	assertExpression(t, expression, orEvaluator, true)
}

// AND evaluation, all true -> [true, true, true] = true
func TestOrWithAllTrue(t *testing.T) {
	expression := buildGroupExp(
		"or",
		true,
		true,
		true,
	)

	assertExpression(t, expression, orEvaluator, true)
}

// AND evaluation, all false -> [false, false, false] = false
func TestOrWithAllFalse(t *testing.T) {
	expression := buildGroupExp(
		"or",
		false,
		false,
		false,
	)

	assertExpression(t, expression, orEvaluator, false)
}

func buildGroupExp(op token.Token, results ...token.Result) token.EvalNode {
	var expressions = make([]*token.EvalNode, len(results))

	for i, res := range results {
		expressions[i] = &token.EvalNode{
			Token:     res,
			Result:    res,
			Kind:      token.PrimitiveVal,
			Childrens: nil,
		}
	}

	var group = token.EvalNode{Token: op, Kind: token.Operator, Childrens: expressions}

	return group
}

func buildSimpleExp(op token.Token, a token.Token, b token.Token) token.EvalNode {
	return token.EvalNode{
		Token: op,
		Kind:  token.Operator,
		Childrens: []*token.EvalNode{
			{
				Token:  a,
				Kind:   token.PrimitiveVal,
				Result: a,
			},
			{
				Token:  b,
				Kind:   token.PrimitiveVal,
				Result: b,
			},
		},
	}
}

func assertExpression(t *testing.T, exp token.EvalNode, evaluator OperatorRunnable, expected token.Result) {
	res := evaluator.Evaluate(&exp)

	if res != expected {
		t.Errorf("Expected %v but got '%v'", expected, res)
	}
}
