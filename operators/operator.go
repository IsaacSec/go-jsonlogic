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
	And                 OperatorToken = "and"
	Or                  OperatorToken = "or"
	Equals              OperatorToken = "=="
	NotEquals           OperatorToken = "!="
	LessThan            OperatorToken = "<"
	LessOrEqualsThan    OperatorToken = "<="
	GreaterThan         OperatorToken = ">"
	GreaterOrEqualsThan OperatorToken = ">="
)

var operatorMap map[OperatorToken]OperatorRunnable = make(map[OperatorToken]OperatorRunnable)

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
