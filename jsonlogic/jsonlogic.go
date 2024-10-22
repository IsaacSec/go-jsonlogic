package jsonlogic

import (
	"github.com/IsaacSec/go-jsonlogic/parser"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/parser/tree"
)

func Apply(rules map[string]any, data any) (bool, error) {
	logicTree := tree.Tree{Root: parser.ParseMap(rules)}

	return logicTree.Eval(data), nil
}

func ApplyTree(rules map[string]any, data any) (*token.EvalNode, error) {

	logicTree := tree.Tree{Root: parser.ParseMap(rules)}

	return logicTree.EvaluateTree(data), nil
}
