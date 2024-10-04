package main

import (
	"testing"
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
	var expression = []Node{
		{
			Token: "==",
			Kind:  Operator,
			Childrens: &[]Node{
				{
					Token: 3,
					Kind:  PrimitiveVal,
				},
				{
					Token: 3,
					Kind:  PrimitiveVal,
				},
			},
		},
	}

	tree.Root = &Node{Token: "and", Kind: Operator, Childrens: &expression}

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
	var expression = []Node{
		{
			Token: "==",
			Kind:  Operator,
			Childrens: &[]Node{
				{
					Token: "pass",
					Kind:  PrimitiveVal,
				},
				{
					Token: "pass",
					Kind:  PrimitiveVal,
				},
			},
		},
		{
			Token: "==",
			Kind:  Operator,
			Childrens: &[]Node{
				{
					Token: 4,
					Kind:  PrimitiveVal,
				},
				{
					Token: 3,
					Kind:  PrimitiveVal,
				},
			},
		},
	}

	tree.Root = &Node{Token: "and", Kind: Operator, Childrens: &expression}

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
	var expression = []Node{
		{
			Token:     2345,
			Kind:      Operator,
			Childrens: &[]Node{},
		},
	}

	tree.Root = &Node{Token: "and", Kind: Operator, Childrens: &expression}

	if tree.Eval() == true {
		t.Errorf("Expected False %+v", tree.Root)
	}
}
