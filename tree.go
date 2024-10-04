package main

import (
	"fmt"
	"reflect"
)

type Kind int
type EvalResult int
type OperatorToken string
type Token interface{}

const (
	And       OperatorToken = "and"
	Or        OperatorToken = "or"
	Equals    OperatorToken = "=="
	NotEquals OperatorToken = "!="
)

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

func ToOperator(t Token) OperatorToken {
	switch t {
	case string(And):
		return And
	case string(Or):
		return Or
	case string(Equals):
		return Equals
	case string(NotEquals):
		return NotEquals
	default:
		fmt.Printf("invalid token '%v' type '%s' for operator \n", t, reflect.TypeOf(t))
		return ""
	}
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

func runOperation(n *Node) EvalResult {
	childs := *n.Childrens

	switch operator := ToOperator(n.Token); operator {
	case And:
		n.CommulativeEval = True
		for i := range *n.Childrens {
			child := &(*n.Childrens)[i]

			if child.CommulativeEval == False {
				n.CommulativeEval = False
				break
			}
		}
	case Equals:
		if len(childs) != 2 {
			fmt.Printf("Cannot evaluate expression with %d arguments, expected 2 \n", len(childs))
			n.CommulativeEval = False
		} else {

			if childs[0] == childs[1] {
				n.CommulativeEval = True
			} else {
				n.CommulativeEval = False
			}

			fmt.Printf("Evaluating [%v == %v] -> %v \n", childs[0].Token, childs[1].Token, n.CommulativeEval.ToString())
		}
	default:
		n.CommulativeEval = False
	}

	return n.CommulativeEval
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
		n.CommulativeEval = runOperation(n)

	default:
		n.CommulativeEval = False
	}

	fmt.Printf("Token [%v] result: %v\n", n.Token, n.CommulativeEval.ToString())

	return n.CommulativeEval
}
