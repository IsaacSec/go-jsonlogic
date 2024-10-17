package operators

import (
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

var andOperator = OperatorRunnable{
	Token: "and",
	Evaluate: func(n *token.EvalNode) token.Result {
		childs := n.Childrens
		// Default value on empty "and"
		res := true

		for i := range childs {
			child := childs[i]

			log.Info(
				"Evaluating exp [%v][ %v ]",
				i,
				child.Result,
			)

			if !child.ToBool() {
				res = false
			}
		}

		return res
	},
}

var orEvaluator = OperatorRunnable{
	Token: "or",
	Evaluate: func(n *token.EvalNode) token.Result {
		childs := n.Childrens
		// Default value on empty "and"
		var res token.Result

		if len(childs) > 0 {
			res = false
		} else {
			res = true
		}

		for i := range childs {
			child := childs[i]

			if child.ToBool() {
				res = true
				//break
			}
		}

		return res
	},
}

var equalsOperator = OperatorRunnable{
	Token: "==",
	Evaluate: func(n *token.EvalNode) (res token.Result) {
		childs := n.Childrens

		if len(childs) != 2 {
			log.Info("Cannot evaluate expression with %d arguments, expected 2", len(childs))
			res = false
		} else {

			first := childs[0].Result
			second := childs[1].Result

			res = first == second

			log.Info(
				"Evaluating [ (%s) %v == (%s) %v ] -> %v",
				reflect.TypeOf(first),
				first,
				reflect.TypeOf(second),
				second,
				res,
			)
		}

		return res
	},
}

var notEqualsEvaluator = OperatorRunnable{
	Token: "!=",
	Evaluate: func(n *token.EvalNode) (res token.Result) {
		childs := n.Childrens

		if len(childs) != 2 {
			log.Info("Cannot evaluate expression with %d arguments, expected 2", len(childs))
			res = false
		} else {

			first := childs[0].Result
			second := childs[1].Result

			res = first != second

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
	OperatorMap["and"] = andOperator
	OperatorMap["or"] = orEvaluator
	OperatorMap["=="] = equalsOperator
	OperatorMap["!="] = notEqualsEvaluator
}
