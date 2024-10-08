package main

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/parser"
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
	var expression = []parser.Node{
		{
			Token: "==",
			Kind:  parser.Operator,
			Childrens: &[]parser.Node{
				{
					Token: 3,
					Kind:  parser.PrimitiveVal,
				},
				{
					Token: 3,
					Kind:  parser.PrimitiveVal,
				},
			},
		},
	}

	tree.Root = &parser.Node{Token: "and", Kind: parser.Operator, Childrens: &expression}

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
	var expression = []parser.Node{
		{
			Token: "==",
			Kind:  parser.Operator,
			Childrens: &[]parser.Node{
				{
					Token: "pass",
					Kind:  parser.PrimitiveVal,
				},
				{
					Token: "pass",
					Kind:  parser.PrimitiveVal,
				},
			},
		},
		{
			Token: "==",
			Kind:  parser.Operator,
			Childrens: &[]parser.Node{
				{
					Token: 4,
					Kind:  parser.PrimitiveVal,
				},
				{
					Token: 3,
					Kind:  parser.PrimitiveVal,
				},
			},
		},
	}

	tree.Root = &parser.Node{Token: "and", Kind: parser.Operator, Childrens: &expression}

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
	var expression = []parser.Node{
		{
			Token:     2345,
			Kind:      parser.Operator,
			Childrens: &[]parser.Node{},
		},
	}

	tree.Root = &parser.Node{Token: "and", Kind: parser.Operator, Childrens: &expression}

	if tree.Eval() == true {
		t.Errorf("Expected False %+v", tree.Root)
	}
}
