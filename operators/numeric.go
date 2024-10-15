package operators

import (
	"fmt"
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser"
)

var lessThanOperator = OperatorRunnable{
	Token: LessThan,
	Evaluate: func(n *parser.Node) parser.EvalResult {
		childs := *n.Childrens

		if len(childs) != 2 {
			fmt.Printf("Cannot evaluate expression with %d arguments, expected 2 \n", len(childs))
			n.CommulativeEval = parser.False
		} else {

			first := childs[0].Token
			second := childs[1].Token

			if reflect.TypeOf(first) != reflect.TypeOf(second) {
				fmt.Printf(
					"Cannot evaluate expression with arguments with different types arg0 (%v) arg1 (%v) \n",
					reflect.TypeOf(first),
					reflect.TypeOf(second),
				)
				n.CommulativeEval = parser.False
				return n.CommulativeEval
			}

			switch reflect.TypeOf(first).Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				first = first.(int)
				second = second.(int)
			default:
				n.CommulativeEval = parser.False
				return n.CommulativeEval
			}

			if first.(int) < second.(int) {
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
	operatorMap[LessThan] = lessThanOperator
}
