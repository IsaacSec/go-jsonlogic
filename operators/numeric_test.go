package operators

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser"
)

/***********************
 *		LESS THAN
 ***********************/

// LESS THAN evaluation, same type equal value -> 40 < 50
func TestLessThanOnSameTypeExpectedTrue(t *testing.T) {
	expression := buildSimpleExp("<", 40, 50)

	assertExpression(t, expression, lessThanOperator, parser.True)
}

// LESS THAN evaluation, same type equal value -> 50 < 40
func TestLessThanOnSameTypeExpectedFalse(t *testing.T) {
	expression := buildSimpleExp("<", 50, 40)

	assertExpression(t, expression, lessThanOperator, parser.False)
}
