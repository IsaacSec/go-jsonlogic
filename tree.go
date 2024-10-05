package main

import (
	"fmt"
)

type Kind int
type EvalResult int
type Token interface{}

const (
	Undefined EvalResult = iota
	True
	False
)

const (
	PrimitiveVal Kind = iota + 1
	ReferenceVal
	Operator
	Array
	Expression
	Object
)

type Node struct {
	Token           Token
	Kind            Kind
	CommulativeEval EvalResult
	Childrens       *[]Node
}

func (er EvalResult) ToString() string {
	switch er {
	case Undefined:
		return "Undefined"
	case True:
		return "True"
	case False:
		return "False"
	default:
		return "Undefined"
	}
}

type Tree struct {
	Root *Node
}

func (t Tree) Eval() bool {
	return t.Root.eval() == True
}

func (n *Node) eval() EvalResult {
	if n.CommulativeEval != Undefined {
		return n.CommulativeEval
	}

	if n.Childrens == nil {
		n.CommulativeEval = False
		return n.CommulativeEval
	} else {
		for i := range *n.Childrens {
			child := &(*n.Childrens)[i]
			child.eval()
		}
	}

	switch kind := n.Kind; kind {
	case Operator:
		n.CommulativeEval = RunOperation(n)

	default:
		n.CommulativeEval = False
	}

	fmt.Printf("Token [%v] result: %v\n", n.Token, n.CommulativeEval.ToString())

	return n.CommulativeEval
}
