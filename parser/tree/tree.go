package tree

import (
	op "github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

type Tree struct {
	Root *token.Node
}

func (t Tree) Eval() bool {
	return eval(t.Root).ToBool()
}

func (t Tree) Flatten() []*token.Node {
	// Start by adding root node
	var flattened = []*token.Node{t.Root}
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

func eval(n *token.Node) token.EvalNode {

	// Todo: check that adresses are different
	new := token.EvalNode{
		Token:  n.Token,
		Kind:   n.Kind,
		Result: false,
	}

	if n.Kind == token.Null {
		new.Result = nil
		return new
	}

	if n.Childrens == nil { // This is a leaf node, Todo: what happens of null value tokens
		new.Result = n.Token
		return new
	} else {
		var children []*token.EvalNode

		for i := range n.Childrens {
			child := n.Childrens[i]
			newChild := eval(child)
			children = append(children, &newChild)
		}

		new.Args = children
	}

	// Node has children
	switch kind := n.Kind; kind {
	case token.Operator:
		new.Result = op.Run(&new)

	default:
		new.Result = false
	}

	log.Info("Token [%v] result: %v", new.Token, new.Result)

	return new
}
