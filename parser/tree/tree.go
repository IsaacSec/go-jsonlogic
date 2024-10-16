package tree

import (
	op "github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

type Tree struct {
	Root *token.Node
}

func (t Tree) Eval() bool {
	return eval(t.Root) == token.True
}

func (t Tree) Flatten() []*token.Node {
	var flattened = []*token.Node{t.Root}
	var Add func(n *token.Node)

	Add = func(n *token.Node) {
		for i := range n.Childrens {
			child := n.Childrens[i]
			flattened = append(flattened, child)
			Add(child)
		}
	}

	Add(t.Root)

	return flattened
}

func eval(n *token.Node) token.EvalResult {
	if n.CommulativeEval != token.Undefined {
		return n.CommulativeEval
	}

	if n.Childrens == nil {
		n.CommulativeEval = token.False
		return n.CommulativeEval
	} else {
		for i := range n.Childrens {
			child := n.Childrens[i]
			eval(child)
		}
	}

	switch kind := n.Kind; kind {
	case token.Operator:
		n.CommulativeEval = op.Run(n)

	default:
		n.CommulativeEval = token.False
	}

	log.Info("Token [%v] result: %v", n.Token, n.CommulativeEval.ToString())

	return n.CommulativeEval
}
