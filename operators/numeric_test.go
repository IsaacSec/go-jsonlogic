package operators

import (
	"testing"
)

/***********************
 *		LESS THAN
 ***********************/

// LESS THAN evaluation, same type equal value -> 40 < 50
func TestLessThanOnSameTypeExpectedTrue(t *testing.T) {
	expression := buildSimpleExp("<", 40, 50)

	assertExpression(t, expression, true)
}

// LESS THAN evaluation, same type equal value -> 50 < 40
func TestLessThanOnSameTypeExpectedFalse(t *testing.T) {
	expression := buildSimpleExp("<", 50, 40)

	assertExpression(t, expression, false)
}
