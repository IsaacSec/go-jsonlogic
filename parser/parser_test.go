package parser

import (
	"fmt"
	"slices"
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
			break
		}
	}
}

// Test parsing strings
func TestParseStringValues(t *testing.T) {
	root, err := ParseJson([]byte(`
		{
			"and" : [
				{ "==": ["one",""] },
				{ "==": ["12312","-$%&*4"] }
			]
		}
	`))

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	nodes := tree.Tree{Root: root}.Flatten()
	assertExpectedTokens(t, nodes, token.PrimitiveVal, "one", "", "12312", "-$%&*4")
}

// Test parsing numbers
// Golang Unmarshal function parse all numeric values as float64
func TestParseNumericValues(t *testing.T) {
	root, err := ParseJson([]byte(`
		{
			"and" : [
				{ "==": [5624,54.23] },
				{ "==": [-1,0] }
			]
		}
	`))

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	nodes := tree.Tree{Root: root}.Flatten()
	assertExpectedTokens(t, nodes, token.PrimitiveVal, 5624.0, 54.23, -1.0, 0.0)
}

// Test parsing arrays
func TestParseArraysValues(t *testing.T) {
	root, err := ParseJson([]byte(`
		{
			"and" : [
				{ "==": [[1,2],[2,2]] },
				{ "==": [[],[]] }
			]
		}
	`))

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	nodes := tree.Tree{Root: root}.Flatten()

	fmt.Printf("%v\n", getNodeValues(nodes))

	// Cannot compare arrays
	if len(nodes) != 11 {
		t.Errorf("Expected 11 tokens but got %v", len(nodes))
	}
}

// Test parsing null values
func TestParseNullValues(t *testing.T) {
	root, err := ParseJson([]byte(`
		{
			"and" : [
				{ "==": [null,null] }
			]
		}
	`))

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	nodes := tree.Tree{Root: root}.Flatten()

	exists := slices.ContainsFunc(nodes, func(n *token.Node) bool {
		return n.Kind == token.Null && !(n.Token != nil)
	})

	if !exists {
		t.Errorf("Expected tokens: %v, node list: %v", nil, getNodeValues(nodes))
	}
}

func assertExpectedTokens[T comparable](t *testing.T, nodes []*token.Node, kind token.Kind, expected ...T) {
	for _, value := range expected {
		exists := slices.ContainsFunc(nodes, func(n *token.Node) bool {
			//fmt.Printf("Kind: %v - %v, Token: %v - %v -> %v %v\n", n.Kind, kind, n.Token, value, n.Kind == kind, n.Token == value)
			return n.Kind == kind && n.Token == value
		})

		if !exists {
			t.Errorf("Expected tokens: %v, node list: %v", value, getNodeValues(nodes))
			break
		}
	}
}

func getNodeValues(nodes []*token.Node) []token.Token {
	var result []token.Token

	for _, n := range nodes {
		result = append(result, n.Token)
	}

	return result
}
