package operators

import (
	"testing"
)

/***********************
 *		LESS THAN
 ***********************/

// LESS THAN evaluation:  40 < 50 -> true
func TestLessThanExpectedTrue(t *testing.T) {
	expression := buildSimpleExp("<", 40, 50)

	assertExpression(t, expression, true)
}

// LESS THAN evaluation : 50 < 40 -> false
func TestLessThanExpectedFalse(t *testing.T) {
	expression := buildSimpleExp("<", 50, 40)

	assertExpression(t, expression, false)
}

/***********************
 *		LESS THAN EQUALS
 ***********************/

// LESS THAN EQUALS evaluation: 1 <= 2 && 2 <= 2 -> true
func TestLessOrEqualsThanExpectedTrue(t *testing.T) {
	exp1 := buildSimpleExp("<=", 1, 2)
	exp2 := buildSimpleExp("<=", 2, 2)

	assertExpression(t, exp1, true)
	assertExpression(t, exp2, true)
}

// LESS THAN EQUALS evaluation: 5 <= 1 --> false
func TestLessOrEqualsThanExpectedFalse(t *testing.T) {
	exp := buildSimpleExp("<=", 5, 1)

	assertExpression(t, exp, false)
}

/***********************
 *		GREATER THAN
 ***********************/

// GREATER THAN evaluation: 60 > 50 -> true
func TestGreaterThanExpectedTrue(t *testing.T) {
	expression := buildSimpleExp(">", 60, 50)

	assertExpression(t, expression, true)
}

// GREATER THAN evaluation: 50 > 60 -> false
func TestGreaterThanExpectedFalse(t *testing.T) {
	expression := buildSimpleExp(">", 50, 60)

	assertExpression(t, expression, false)
}

/***********************
 *		GREATER THAN EQUALS
 ***********************/

// GREATER THAN EQUALS evaluation: 2 >= 1 && 2 >= 2 -> true
func TestGreaterOrEqualsThanExpectedTrue(t *testing.T) {
	exp1 := buildSimpleExp(">=", 2, 1)
	exp2 := buildSimpleExp(">=", 2, 2)

	assertExpression(t, exp1, true)
	assertExpression(t, exp2, true)
}

// GREATER THAN EQUALS evaluation: 1 >= 5 --> false
func TestGreaterOrEqualsThanExpectedFalse(t *testing.T) {
	exp := buildSimpleExp(">=", 1, 5)

	assertExpression(t, exp, false)
}
