package main

import (
	"fmt"
	"reflect"
)

type Kind int
type EvalResult int
type Operator string
type Token interface{}

const (
	And       Operator = "and"
	Or        Operator = "or"
	Equals    Operator = "=="
	NotEquals Operator = "!="
)

const (
	Undefined EvalResult = iota
	True
	False
)

const (
	PrimitiveVal Kind = iota + 1
	ReferenceVal
	ValueOperator
	Array
	Expression
	BoolOperator
	Object
)

type Node struct {
	Token           Token
	Kind            Kind
	CommulativeEval EvalResult
	Childrens       *[]Node
}

func ToOperator(t Token) Operator {
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

func (t Tree) eval() bool {
	return t.Root.eval() == True
}

func (n *Node) SetEval(er EvalResult) {
	(*n).CommulativeEval = er
}

func RunBoolOperation(n *Node) EvalResult {
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
	default:
		n.CommulativeEval = False
	}

	return n.CommulativeEval
}

func RunValueOperation(n *Node) EvalResult {
	childs := *n.Childrens

	switch operator := ToOperator(n.Token); operator {
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
	case BoolOperator:
		n.CommulativeEval = RunBoolOperation(n)

	case ValueOperator:
		n.CommulativeEval = RunValueOperation(n)

	default:
		n.CommulativeEval = False
	}

	fmt.Printf("Token [%v] result: %v\n", n.Token, n.CommulativeEval.ToString())

	return n.CommulativeEval
}
