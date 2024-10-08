package operators

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser"
)

/*
EQUALS evaluation, same type equal value -> 53 == 53
*/
func TestEqualOnSameTypeAndValue(t *testing.T) {
	expression := buildSimpleExp("==", 53, 53)

	res := equalsOperator.Evaluate(&expression)
	if res != parser.True {
		t.Errorf("Expected True but got '%v'", res.ToString())
	}
}

/*
EQUALS evaluation, different type equal value -> 53 == "53"
*/
func TestEqualOnDiffType(t *testing.T) {
	expression := buildSimpleExp("==", 53, "fifty three")

	res := equalsOperator.Evaluate(&expression)
	if res != parser.False {
		t.Errorf("Expected False but got '%v'", res.ToString())
	}
}

/*
EQUALS evaluation, different value -> 3421 == -123
*/
func TestEqualOnDiffValue(t *testing.T) {
	expression := buildSimpleExp("==", 3421, -123)

	res := equalsOperator.Evaluate(&expression)
	if res != parser.False {
		t.Errorf("Expected False but got '%v'", res.ToString())
	}
}

// Todo: Add conversion type when possible i.e. -> 1 == "1" == 1.0

/*
EQUALS evaluation, different type and equal value -> 42 == 42.0
*/
func TestEqualOnDiffTypeWithSameValue(t *testing.T) {
	expression := buildSimpleExp("==", 42, 42.0)

	res := equalsOperator.Evaluate(&expression)
	if res != parser.False {
		t.Errorf("Expected False but got '%v'", res.ToString())
	}
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
