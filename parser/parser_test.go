package parser

import (
	"testing"

	"github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/parser/tree"
	"github.com/IsaacSec/go-jsonlogic/util/maps"
)

// Test all listed operators by default
func TestParseOperators(t *testing.T) {
	root, err := ParseJson([]byte(`
		{
			"and" : [
				{ "or": [true, false] },
				{ "==": [0,0] },
				{ "!=": [1,0] },
				{ "<": [0,1] },
				{ ">": [1,0] }
			]
		}
	`))

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	nodes := tree.Tree{Root: root}.Flatten()
	found := make(map[string]bool, 0)

	for i := range nodes {
		node := nodes[i]
		if node.Kind == token.Operator {
			found[node.Token.(string)] = true
		}
	}

	for key := range operators.OperatorMap {
		if _, ok := found[key]; !ok {
			t.Errorf("Missing operator '%v', found: %v, expected: %+v\n", key, maps.GetKeys(found), maps.GetKeys(operators.OperatorMap))
		}
	}
}
