package tree

import (
	op "github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/parser/varref"
)

type Tree struct {
	Root *token.Node
}

func (t Tree) Eval(data any) bool {
	var input map[string]any

	if data != nil {
		input = data.(map[string]any)
	} else {
		input = make(map[string]any)
	}

	return eval(t.Root, input).ToBool()
}

// A json object as input
// Todo: implement array input
func (t Tree) EvaluateTree(data any) *token.EvalNode {
	var input map[string]any

	if data != nil {
		input = data.(map[string]any)
	} else {
		input = make(map[string]any)
	}

	return eval(t.Root, input)
}

func (t Tree) Flatten() []*token.Node {
	// Start by adding root node
	flattened := []*token.Node{t.Root}
	var Add func(n *token.Node)

	Add = func(n *token.Node) {
		for i := range n.Childrens {
			child := n.Childrens[i]

			if child != nil {
				flattened = append(flattened, child)
				Add(child)
			}
		}
	}

	Add(t.Root)

	return flattened
}

func eval(n *token.Node, data map[string]any) *token.EvalNode {
	// Todo: check that adresses are different
	new := token.EvalNode{
		Token:  n.Token,
		Kind:   n.Kind,
		Result: false,
	}

	if n.Kind == token.ReferenceVal {
		new.Result = varref.GetValue(data, n.Token)
		new.Kind = token.PrimitiveVal
		return &new
	}

	if n.Kind == token.Null {
		new.Result = nil
		return &new
	}

	if n.Childrens == nil { // This is a leaf node, Todo: what happens of null value tokens
		new.Result = n.Token
		return &new
	} else {
		var args []*token.EvalNode

		for i := range n.Childrens {
			child := n.Childrens[i]
			newChild := eval(child, data)
			args = append(args, newChild)
		}

		new.Args = args
	}

	// Node has children
	switch kind := n.Kind; kind {
	case token.Operator:
		new.Result = op.Run(&new)

	default:
		new.Result = false
	}

	return &new
}
