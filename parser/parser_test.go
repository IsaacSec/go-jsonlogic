package parser

import (
	"fmt"
	"slices"
	"testing"

	"github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/parser/tree"
	"github.com/IsaacSec/go-jsonlogic/util/maps"
	"github.com/stretchr/testify/assert"
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
				{ "<=": [0,0] },
				{ ">": [1,0] },
				{ ">=": [1,1] },
				{ "!": [true]}
			]
		}
	`))

	assert.NoError(t, err, "pasing json error")

	nodes := tree.Tree{Root: root}.Flatten()
	found := make(map[string]bool, 0)

	for i := range nodes {
		node := nodes[i]
		if node.Kind == token.Operator {
			found[node.Token.(string)] = true
		}
	}

	assert.ElementsMatch(t, maps.GetKeys(found), maps.GetKeys(operators.Operators), "operator missing")
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

	assert.NoError(t, err, "pasing json error")

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

	assert.NoError(t, err, "pasing json error")

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

	assert.NoError(t, err, "pasing json error")

	nodes := tree.Tree{Root: root}.Flatten()

	fmt.Printf("%v\n", getNodeValues(nodes))

	// Cannot compare arrays
	assert.Equal(t, 11, len(nodes), "incorrect number of nodes identified in json")
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

	assert.NoError(t, err, "pasing json error")

	nodes := tree.Tree{Root: root}.Flatten()

	exists := slices.ContainsFunc(nodes, func(n *token.Node) bool {
		return n.Kind == token.Null && !(n.Token != nil)
	})

	assert.True(t, exists, fmt.Sprintf("Expected tokens: %v, node list: %v", nil, getNodeValues(nodes)))
}

func assertExpectedTokens[T comparable](t *testing.T, nodes []*token.Node, kind token.Kind, expected ...T) {
	for _, value := range expected {
		exists := slices.ContainsFunc(nodes, func(n *token.Node) bool {
			return n.Kind == kind && n.Token == value
		})

		assert.True(t, exists, fmt.Sprintf("Expected tokens: %v, node list: %v", nil, getNodeValues(nodes)))
	}
}

func getNodeValues(nodes []*token.Node) []token.Token {
	var result []token.Token

	for _, n := range nodes {
		result = append(result, n.Token)
	}

	return result
}
