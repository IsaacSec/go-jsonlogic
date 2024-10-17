package operators

import (
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

type OperatorRunnable struct {
	Token    string
	Evaluate func(n *token.EvalNode) token.Result
}

var OperatorMap map[string]OperatorRunnable = make(map[string]OperatorRunnable)

func Run(n *token.EvalNode) token.Result {

	var res token.Result

	switch n.Token.(type) {
	case string:
		operator, ok := OperatorMap[n.Token.(string)]

		if ok {
			res = operator.Evaluate(n)
		} else {
			log.Error("Undefined operator %s", n.Token)

			// Todo: change with an error handler
			res = false // Default on error value
		}

	default:
		log.Error("Token [%v] with wrong type [%v]", n.Token, reflect.TypeOf(n.Token))
	}

	return res
}
