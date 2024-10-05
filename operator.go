package main

import (
	"fmt"
	"reflect"
)

type OperatorToken string

type OperatorRunnable struct {
	Token    OperatorToken
	Evaluate func(n *Node) EvalResult
}

const (
	And       OperatorToken = "and"
	Or        OperatorToken = "or"
	Equals    OperatorToken = "=="
	NotEquals OperatorToken = "!="
)

var operatorMap map[OperatorToken]OperatorRunnable = make(map[OperatorToken]OperatorRunnable)

var andOperator = OperatorRunnable{
	Token: And,
	Evaluate: func(n *Node) EvalResult {
		childs := *n.Childrens

		n.CommulativeEval = True
		for i := range childs {
			child := &(childs)[i]

			if child.CommulativeEval == False {
				n.CommulativeEval = False
				break
			}
		}

		return n.CommulativeEval
	},
}

var equalsOperator = OperatorRunnable{
	Token: Equals,
	Evaluate: func(n *Node) EvalResult {
		childs := *n.Childrens

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

		return n.CommulativeEval
	},
}

func init() {
	operatorMap[And] = andOperator
	operatorMap[Equals] = equalsOperator
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

func RunOperation(n *Node) EvalResult {

	operator, ok := operatorMap[ToOperator(n.Token)]

	if ok {
		n.CommulativeEval = operator.Evaluate(n)
	} else {
		n.CommulativeEval = False
	}

	return n.CommulativeEval
}
