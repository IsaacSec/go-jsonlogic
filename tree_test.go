package main

import "testing"

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
			Kind:  ValueOperator,
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

	tree.Root = &Node{Token: "and", Kind: BoolOperator, Childrens: &expression}

	if tree.eval() == false {
		t.Errorf("Expected true %+v", tree.Root)
	}
}
