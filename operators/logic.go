package operators

import (
	"fmt"
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser"
)

var andOperator = OperatorRunnable{
	Token: And,
	Evaluate: func(n *parser.Node) parser.EvalResult {
		childs := *n.Childrens

		n.CommulativeEval = parser.True
		for i := range childs {
			child := &(childs)[i]

			fmt.Printf(
				"Evaluating exp [%v][ %v ] \n",
				i,
				child.CommulativeEval.ToString(),
			)

			if child.CommulativeEval == parser.False {
				n.CommulativeEval = parser.False
				break
			}
		}

		return n.CommulativeEval
	},
}

var orEvaluator = OperatorRunnable{
	Token: Or,
	Evaluate: func(n *parser.Node) parser.EvalResult {
		childs := *n.Childrens

		if len(childs) > 0 {
			n.CommulativeEval = parser.False
		} else {
			n.CommulativeEval = parser.True
		}

		for i := range childs {
			child := &(childs)[i]

			if child.CommulativeEval == parser.True {
				n.CommulativeEval = parser.True
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

			first := childs[0].Token
			second := childs[1].Token

			if first == second {
				n.CommulativeEval = parser.True
			} else {
				n.CommulativeEval = parser.False
			}

			fmt.Printf(
				"Evaluating [ (%s) %v == (%s) %v ] -> %v \n",
				reflect.TypeOf(first),
				first,
				reflect.TypeOf(second),
				second,
				n.CommulativeEval.ToString(),
			)
		}

		return n.CommulativeEval
	},
}

var notEqualsEvaluator = OperatorRunnable{
	Token: NotEquals,
	Evaluate: func(n *parser.Node) parser.EvalResult {
		childs := *n.Childrens

		if len(childs) != 2 {
			fmt.Printf("Cannot evaluate expression with %d arguments, expected 2 \n", len(childs))
			n.CommulativeEval = parser.False
		} else {

			first := childs[0].Token
			second := childs[1].Token

			if first != second {
				n.CommulativeEval = parser.True
			} else {
				n.CommulativeEval = parser.False
			}

			fmt.Printf(
				"Evaluating [ (%s) %v %s (%s) %v ] -> %v \n",
				reflect.TypeOf(first),
				first,
				n.Token,
				reflect.TypeOf(second),
				second,
				n.CommulativeEval.ToString(),
			)
		}

		return n.CommulativeEval
	},
}

func init() {
	operatorMap[And] = andOperator
	operatorMap[Or] = orEvaluator
	operatorMap[Equals] = equalsOperator
	operatorMap[NotEquals] = notEqualsEvaluator
}
