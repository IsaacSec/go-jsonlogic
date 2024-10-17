package operators

import (
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/util"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

var lessThanOperator = OperatorRunnable{
	Token: "<",
	Evaluate: func(n *token.EvalNode) (res token.Result) {
		childs := n.Childrens

		if len(childs) != 2 {
			log.Info("Cannot evaluate expression with %d arguments, expected 2", len(childs))
			res = false
		} else {

			first, second, err := util.ConvertToFloat(childs[0].Result, childs[1].Result)

			if err != nil {
				log.Warn("Conversion failed: %s", err)
				res = false
				return res
			}

			res = first < second

			log.Info(
				"Evaluating [ (%s) %v %s (%s) %v ] -> %v",
				reflect.TypeOf(first),
				first,
				n.Token,
				reflect.TypeOf(second),
				second,
				res,
			)
		}

		return res
	},
}

func init() {
	OperatorMap["<"] = lessThanOperator
}
