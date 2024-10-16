package main

import (
	"fmt"

	op "github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser"
)

type Tree struct {
	Root *parser.Node
}

func (t Tree) Eval() bool {
	return eval(t.Root) == parser.True
}

func eval(n *parser.Node) parser.EvalResult {
	if n.CommulativeEval != parser.Undefined {
		return n.CommulativeEval
	}

	if n.Childrens == nil {
		n.CommulativeEval = parser.False
		return n.CommulativeEval
	} else {
		for i := range *n.Childrens {
			child := &(*n.Childrens)[i]
			eval(child)
		}
	}

	switch kind := n.Kind; kind {
	case parser.Operator:
		n.CommulativeEval = op.Run(n)

	default:
		n.CommulativeEval = parser.False
	}

	log.Info("Token [%v] result: %v", n.Token, n.CommulativeEval.ToString())

	return n.CommulativeEval
}
