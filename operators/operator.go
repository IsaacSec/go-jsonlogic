package operators

import (
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

type Operator func(token.Args) token.Result

var Operators map[string]Operator = make(map[string]Operator)

func Run(n *token.EvalNode) token.Result {

	var res token.Result

	switch n.Token.(type) {
	case string:
		evaluate, ok := Operators[n.Token.(string)]

		if ok {
			args := n.Args
			res = evaluate(args)

			log.Info(
				"Evaluation %s %v -> %v",
				n.Token,
				args.GetArgValueAndType(),
				res,
			)
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

func init() {
	Operators["and"] = And
	Operators["or"] = Or
	Operators["=="] = Equals
	Operators["!="] = NotEquals
	Operators["<"] = LessThan
}
