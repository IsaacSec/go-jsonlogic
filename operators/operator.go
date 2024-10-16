package operators

import "github.com/IsaacSec/go-jsonlogic/parser/token"

type OperatorRunnable struct {
	Token    string
	Evaluate func(n *token.Node) token.EvalResult
}

var OperatorMap map[string]OperatorRunnable = make(map[string]OperatorRunnable)

func Run(n *token.Node) token.EvalResult {

	operator, ok := OperatorMap[n.Token.(string)]

	if ok {
		n.CommulativeEval = operator.Evaluate(n)
	} else {
		n.CommulativeEval = token.False
	}

	return n.CommulativeEval
}
