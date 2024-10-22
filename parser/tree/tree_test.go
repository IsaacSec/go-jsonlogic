package tree

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/stretchr/testify/assert"
)

/*
Test simple tree evaluation equivalent to

	{
		"and":[
			{"==": [3, 3]}
		]
	}
*/
func TestTree(t *testing.T) {

	var tree Tree = Tree{}
	var expression = []*token.Node{
		{
			Token: "==",
			Kind:  token.Operator,
			Childrens: []*token.Node{
				{
					Token: 3,
					Kind:  token.PrimitiveVal,
				},
				{
					Token: 3,
					Kind:  token.PrimitiveVal,
				},
			},
		},
	}

	tree.Root = &token.Node{Token: "and", Kind: token.Operator, Childrens: expression}

	assert.Equal(t, true, tree.Eval(nil))
}

/*
Test simple tree evaluation equivalent to

	{
		"and":[
			{"==": ["pass", "pass"]}
			{"==": [4, 3]},
		]
	}
*/
func TestFalseTree(t *testing.T) {

	var tree Tree = Tree{}
	var expression = []*token.Node{
		{
			Token: "==",
			Kind:  token.Operator,
			Childrens: []*token.Node{
				{
					Token: "pass",
					Kind:  token.PrimitiveVal,
				},
				{
					Token: "pass",
					Kind:  token.PrimitiveVal,
				},
			},
		},
		{
			Token: "==",
			Kind:  token.Operator,
			Childrens: []*token.Node{
				{
					Token: 4,
					Kind:  token.PrimitiveVal,
				},
				{
					Token: 3,
					Kind:  token.PrimitiveVal,
				},
			},
		},
	}

	tree.Root = &token.Node{Token: "and", Kind: token.Operator, Childrens: expression}

	assert.Equal(t, false, tree.Eval(nil))
}
