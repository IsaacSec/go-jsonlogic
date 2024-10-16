package operators

import (
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

var andOperator = OperatorRunnable{
	Token: "and",
	Evaluate: func(n *token.Node) token.EvalResult {
		childs := n.Childrens

		n.CommulativeEval = token.True
		for i := range childs {
			child := childs[i]

			log.Info(
				"Evaluating exp [%v][ %v ]",
				i,
				child.CommulativeEval.ToString(),
			)

			if child.CommulativeEval == token.False {
				n.CommulativeEval = token.False
				break
			}
		}

		return n.CommulativeEval
	},
}

var orEvaluator = OperatorRunnable{
	Token: "or",
	Evaluate: func(n *token.Node) token.EvalResult {
		childs := n.Childrens

		if len(childs) > 0 {
			n.CommulativeEval = token.False
		} else {
			n.CommulativeEval = token.True
		}

		for i := range childs {
			child := childs[i]

			if child.CommulativeEval == token.True {
				n.CommulativeEval = token.True
				break
			}
		}

		return n.CommulativeEval
	},
}

var equalsOperator = OperatorRunnable{
	Token: "==",
	Evaluate: func(n *token.Node) token.EvalResult {
		childs := n.Childrens

		if len(childs) != 2 {
			log.Info("Cannot evaluate expression with %d arguments, expected 2", len(childs))
			n.CommulativeEval = token.False
		} else {

			first := childs[0].Token
			second := childs[1].Token

			if first == second {
				n.CommulativeEval = token.True
			} else {
				n.CommulativeEval = token.False
			}

			log.Info(
				"Evaluating [ (%s) %v == (%s) %v ] -> %v",
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
	Token: "!=",
	Evaluate: func(n *token.Node) token.EvalResult {
		childs := n.Childrens

		if len(childs) != 2 {
			log.Info("Cannot evaluate expression with %d arguments, expected 2", len(childs))
			n.CommulativeEval = token.False
		} else {

			first := childs[0].Token
			second := childs[1].Token

			if first != second {
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
	OperatorMap["and"] = andOperator
	OperatorMap["or"] = orEvaluator
	OperatorMap["=="] = equalsOperator
	OperatorMap["!="] = notEqualsEvaluator
}
