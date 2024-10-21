package jsonlogic

import (
	"github.com/IsaacSec/go-jsonlogic/parser"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/parser/tree"
)

func Apply(rules any, data any) (bool, error) {
	logicTree := tree.Tree{Root: parser.ParseMap(rules)}

	return logicTree.Eval(), nil
}

func ApplyTree(rules any, data any) (*token.EvalNode, error) {

	logicTree := tree.Tree{Root: parser.ParseMap(rules)}

	logicTree.EvaluateTree()

	return logicTree.EvaluateTree(), nil
}
