package main

import (
	"fmt"
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
	PlaceHolder
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

func ToOperator(s string) (Operator, error) {
	switch s {
	case string(And):
		return And, nil
	case string(Or):
		return Or, nil
	case string(Equals):
		return Equals, nil
	case string(NotEquals):
		return NotEquals, nil
	default:
		return "", fmt.Errorf("invalid token '%s' for bool operator", s)
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
	}

	return "Undefined"
}

type Tree struct {
	Root *Node
}

func (t Tree) eval() bool {
	return t.Root.eval() == True
}

func (n Node) eval() EvalResult {
	if n.CommulativeEval != Undefined {
		return n.CommulativeEval
	}

	if n.Childrens == nil {
		n.CommulativeEval = False
		return n.CommulativeEval
	} else {
		for _, child := range *n.Childrens {
			child.eval()
		}
	}

	switch kind := n.Kind; kind {
	case BoolOperator:
		switch operator, _ := ToOperator((n.Token).(string)); operator {
		case And:
			n.CommulativeEval = True
			for _, child := range *n.Childrens {
				if child.CommulativeEval == False {
					n.CommulativeEval = False
					break
				}
			}
		default:
			n.CommulativeEval = False
		}

	case ValueOperator:
		switch operator, _ := ToOperator((n.Token).(string)); operator {
		case Equals:
			if len(*n.Childrens) != 2 {
				fmt.Printf("Cannot evaluate expression with %d arguments, expected 2 \n", len(*n.Childrens))
				n.CommulativeEval = False
			} else {
				fmt.Printf("Evaluating %v == %v \n", (*n.Childrens)[0].Token, (*n.Childrens)[1].Token)
				if (*n.Childrens)[0] == (*n.Childrens)[1] {
					n.CommulativeEval = True
				} else {
					n.CommulativeEval = False
				}
			}
		default:
			n.CommulativeEval = False
		}
	default:
		n.CommulativeEval = False
	}

	fmt.Printf("Token [%v] result: %v\n", n.Token, n.CommulativeEval.ToString())

	return n.CommulativeEval
}
