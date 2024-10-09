package operators

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser"
)

/***********************
 *		EQUALS
 ***********************/

// EQUALS evaluation, same type equal value -> 53 == 53
func TestEqualsOnSameTypeAndValue(t *testing.T) {
	expression := buildSimpleExp("==", 53, 53)

	assertExpression(t, expression, equalsOperator, parser.True)
}

// EQUALS evaluation, different type equal value -> 53 == "53"
func TestEqualsOnDiffType(t *testing.T) {
	expression := buildSimpleExp("==", 53, "fifty three")

	assertExpression(t, expression, equalsOperator, parser.False)
}

// EQUALS evaluation, different value -> 3421 == -123
func TestEqualsOnDiffValue(t *testing.T) {
	expression := buildSimpleExp("==", 3421, -123)

	assertExpression(t, expression, equalsOperator, parser.False)
}

// Todo: Add conversion type when possible i.e. -> 1 == "1" == 1.0

// EQUALS evaluation, different type and equal value -> 42 == 42.0
func TestEqualsOnDiffTypeWithSameValue(t *testing.T) {
	expression := buildSimpleExp("==", 42, 42.0)

	assertExpression(t, expression, equalsOperator, parser.False)
}

/***********************
 *		NOT EQUALS
 ***********************/

// NOT EQUALS evaluation, same type and different value -> 35 != 53
func TestNotEqualsWithDiffValue(t *testing.T) {
	expression := buildSimpleExp("!=", 35, 53)

	assertExpression(t, expression, notEqualsEvaluator, parser.True)
}

// NOT EQUALS evaluation, same type and equal value -> 53 != 53
func TestNotEqualsWithSameValue(t *testing.T) {
	expression := buildSimpleExp("!=", 53, 53)

	assertExpression(t, expression, notEqualsEvaluator, parser.False)
}

/***********************
 *		AND
 ***********************/

// AND evaluation, empty list -> [] = true
func TestAndWithEmptyList(t *testing.T) {
	expression := buildGroupExp(
		"and",
	)

	assertExpression(t, expression, andOperator, parser.True)
}

// AND evaluation, one true and multiple false -> [true, false, false] = false
func TestAndWithMultipleFalseAndAtLeastOneTrue(t *testing.T) {
	expression := buildGroupExp(
		"and",
		parser.True,
		parser.False,
		parser.False,
	)

	assertExpression(t, expression, andOperator, parser.False)
}

// AND evaluation, all true -> [true, true, true] = true
func TestAndWithAllTrue(t *testing.T) {
	expression := buildGroupExp(
		"and",
		parser.True,
		parser.True,
		parser.True,
	)

	assertExpression(t, expression, andOperator, parser.True)
}

// AND evaluation, all false -> [false, false, false] = false
func TestAndWithAllFalse(t *testing.T) {
	expression := buildGroupExp(
		"and",
		parser.False,
		parser.False,
		parser.False,
	)

	assertExpression(t, expression, andOperator, parser.False)
}

/***********************
 *		OR
 ***********************/

// OR evaluation, empty list -> [] = true
func TestOrWithEmptyList(t *testing.T) {
	expression := buildGroupExp(
		"or",
	)

	assertExpression(t, expression, orEvaluator, parser.True)
}

// OR evaluation, one true and multiple false -> [false, true, false] = true
func TestOrWithMultipleFalseAndAtLeastOneTrue(t *testing.T) {
	expression := buildGroupExp(
		"or",
		parser.False,
		parser.True,
		parser.False,
	)

	assertExpression(t, expression, orEvaluator, parser.True)
}

// AND evaluation, all true -> [true, true, true] = true
func TestOrWithAllTrue(t *testing.T) {
	expression := buildGroupExp(
		"or",
		parser.True,
		parser.True,
		parser.True,
	)

	assertExpression(t, expression, orEvaluator, parser.True)
}

// AND evaluation, all false -> [false, false, false] = false
func TestOrWithAllFalse(t *testing.T) {
	expression := buildGroupExp(
		"or",
		parser.False,
		parser.False,
		parser.False,
	)

	assertExpression(t, expression, orEvaluator, parser.False)
}

func buildGroupExp(op parser.Token, results ...parser.EvalResult) parser.Node {
	var expressions = make([]parser.Node, len(results))

	for i, res := range results {
		expressions[i] = parser.Node{
			Token:           res.ToString() == "True",
			CommulativeEval: res,
			Kind:            parser.PrimitiveVal,
			Childrens:       nil,
		}
	}

	var group = parser.Node{Token: op, Kind: parser.Operator, Childrens: &expressions}

	return group
}

func buildSimpleExp(op parser.Token, a parser.Token, b parser.Token) parser.Node {
	return parser.Node{
		Token: op,
		Kind:  parser.Operator,
		Childrens: &[]parser.Node{
			{
				Token: a,
				Kind:  parser.PrimitiveVal,
			},
			{
				Token: b,
				Kind:  parser.PrimitiveVal,
			},
		},
	}
}

func assertExpression(t *testing.T, exp parser.Node, evaluator OperatorRunnable, expected parser.EvalResult) {
	res := evaluator.Evaluate(&exp)

	if res != expected {
		t.Errorf("Expected %v but got '%v'", expected.ToString(), res.ToString())
	}
}
