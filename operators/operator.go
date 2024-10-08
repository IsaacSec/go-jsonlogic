package operators

import (
	"fmt"
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser"
)

type OperatorToken string

type OperatorRunnable struct {
	Token    OperatorToken
	Evaluate func(n *parser.Node) parser.EvalResult
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
	Evaluate: func(n *parser.Node) parser.EvalResult {
		childs := *n.Childrens

		n.CommulativeEval = parser.True
		for i := range childs {
			child := &(childs)[i]

			if child.CommulativeEval == parser.False {
				n.CommulativeEval = parser.False
				break
			}
		}

		return n.CommulativeEval
	},
}

var equalsOperator = OperatorRunnable{
	Token: Equals,
	Evaluate: func(n *parser.Node) parser.EvalResult {
		childs := *n.Childrens

		if len(childs) != 2 {
			fmt.Printf("Cannot evaluate expression with %d arguments, expected 2 \n", len(childs))
			n.CommulativeEval = parser.False
		} else {

			if childs[0] == childs[1] {
				n.CommulativeEval = parser.True
			} else {
				n.CommulativeEval = parser.False
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

func ToOperator(t parser.Token) OperatorToken {
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

func Run(n *parser.Node) parser.EvalResult {

	operator, ok := operatorMap[ToOperator(n.Token)]

	if ok {
		n.CommulativeEval = operator.Evaluate(n)
	} else {
		n.CommulativeEval = parser.False
	}

	return n.CommulativeEval
}
