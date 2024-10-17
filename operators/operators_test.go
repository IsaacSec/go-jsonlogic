package operators

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
)

/*
Test evaluating invalid operator (int are not valid for json key)

	{
		2345:[]
	}
*/
func TestInvalidTokenInOperator(t *testing.T) {

	var expression = token.EvalNode{
		Token:     2345,
		Kind:      token.Operator,
		Childrens: make([]*token.EvalNode, 0),
	}

	if Run(&expression) == true {
		t.Errorf("Expected False %+v", expression)
	}
}
