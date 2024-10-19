package operators

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/stretchr/testify/assert"
)

/*
Test evaluating invalid operator (int are not valid for json key)

	{
		2345:[]
	}
*/
func TestInvalidTokenInOperator(t *testing.T) {

	var expression = token.EvalNode{
		Token: 2345,
		Kind:  token.Operator,
		Args:  make([]*token.EvalNode, 0),
	}

	if Run(&expression) == true {
		t.Errorf("Expected False %+v", expression)
	}
}

/****************************************************
 *			Operators test utilities
 ****************************************************/

func buildGroupExp(op token.Token, results ...token.Result) token.EvalNode {
	var expressions = make([]*token.EvalNode, len(results))

	for i, res := range results {
		expressions[i] = &token.EvalNode{
			Token:  res,
			Result: res,
			Kind:   token.PrimitiveVal,
			Args:   nil,
		}
	}

	var group = token.EvalNode{Token: op, Kind: token.Operator, Args: expressions}

	return group
}

func buildSimpleExp(op token.Token, a token.Token, b token.Token) token.EvalNode {
	return token.EvalNode{
		Token: op,
		Kind:  token.Operator,
		Args: []*token.EvalNode{
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

func assertExpression(t *testing.T, exp token.EvalNode, expected token.Result) {
	actual := Run(&exp)

	assert.Equal(t, expected, actual)
}
