package operators

import (
	"fmt"
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser"
	"github.com/IsaacSec/go-jsonlogic/util"
)

var lessThanOperator = OperatorRunnable{
	Token: LessThan,
	Evaluate: func(n *parser.Node) parser.EvalResult {
		childs := *n.Childrens

		if len(childs) != 2 {
			fmt.Printf("Cannot evaluate expression with %d arguments, expected 2 \n", len(childs))
			n.CommulativeEval = parser.False
		} else {

			first, second, err := util.ConvertToFloat(childs[0].Token, childs[1].Token)

			if err != nil {
				fmt.Printf("Conversion failed: %s\n", err)
				n.CommulativeEval = parser.False
				return n.CommulativeEval
			}

			if first < second {
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
