package main

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
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

	if tree.Eval() == false {
		t.Errorf("Expected true %+v", tree.Root)
	}
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

	if tree.Eval() == true {
		t.Errorf("Expected true %+v", tree.Root)
	}
}

/*
Test evaluating invalid operator (int are not valid for json key)

	{
		2345:[]
	}
*/
func TestInvalidTokenInOperator(t *testing.T) {

	var tree Tree = Tree{}
	var expression = []*token.Node{
		{
			Token: 2345,
			Kind:  token.Operator,
		},
	}

	tree.Root = &token.Node{Token: "and", Kind: token.Operator, Childrens: expression}

	if tree.Eval() == true {
		t.Errorf("Expected False %+v", tree.Root)
	}
}
