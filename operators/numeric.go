package operators

import (
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/util"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

var lessThanOperator = OperatorRunnable{
	Token: "<",
	Evaluate: func(n *token.Node) token.EvalResult {
		childs := n.Childrens

		if len(childs) != 2 {
			log.Info("Cannot evaluate expression with %d arguments, expected 2", len(childs))
			n.CommulativeEval = token.False
		} else {

			first, second, err := util.ConvertToFloat(childs[0].Token, childs[1].Token)

			if err != nil {
				log.Warn("Conversion failed: %s", err)
				n.CommulativeEval = token.False
				return n.CommulativeEval
			}

			if first < second {
				n.CommulativeEval = token.True
			} else {
				n.CommulativeEval = token.False
			}

			log.Info(
				"Evaluating [ (%s) %v %s (%s) %v ] -> %v",
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
	OperatorMap["<"] = lessThanOperator
}
