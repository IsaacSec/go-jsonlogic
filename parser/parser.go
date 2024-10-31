package parser

import (
	"encoding/json"

	"github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/util/logger"
)

func ParseJson(raw []byte) (*token.Node, error) {
	var rules interface{}

	err := json.Unmarshal(raw, &rules)

	if err != nil {
		return nil, err
	}

	return parse(rules), nil
}

func ParseMap(rules map[string]interface{}) *token.Node {

	return parse(rules)
}

func parse(rules interface{}) *token.Node {

	switch val := rules.(type) {
	case map[string]interface{}:
		// Determine if it is operator or object
		if len(val) == 1 {

			// Identify variable reference
			if ref, found := val["var"]; found {
				return buildValueRefNode(ref)
			}

			// Get operator key and children
			for op, children := range val {
				// Check if operator exits
				if _, ok := operators.Operators[op]; ok {
					// Todo: Add cast error handler
					logger.Info("Building operation node. Token: %v, Children: %v", op, children)

					switch args := children.(type) {
					case []interface{}:
						return buildOperationNode(op, args)

					default:
						logger.Error("Error parsing: Value: %v Operator: %v -> Children: %v", args, op, children)

						return buildOperationNode(op, []interface{}{args})
					}
				}
			}
		}

		return buildObjectNode(val)

	case []interface{}:
		return buildArrayNode(val)

	case string, float64, bool, int:
		return buildValueNode(val)

	default:
		return buildNullNode()
	}
}

func buildValueRefNode(ref interface{}) *token.Node {
	node := &token.Node{
		Kind:  token.ReferenceVal,
		Token: ref,
	}

	return node
}

func buildOperationNode(op token.Token, args []interface{}) *token.Node {
	node := &token.Node{
		Kind:  token.Operator,
		Token: op,
	}

	for _, arg := range args {
		// Recursion (DFS) to parse every node
		node.Childrens = append(node.Childrens, parse(arg))
	}

	return node
}

func buildArrayNode(array []interface{}) *token.Node {
	node := &token.Node{
		Kind:      token.Array,
		Childrens: make([]*token.Node, len(array)),
		Token:     array,
	}

	for _, item := range array {
		// Recursion (DFS) to parse every node
		node.Childrens = append(node.Childrens, parse(item))
	}

	return node
}

func buildObjectNode(obj interface{}) *token.Node {

	return &token.Node{
		Kind:  token.Object,
		Token: obj,
	}
}

func buildValueNode(value interface{}) *token.Node {

	return &token.Node{
		Kind:  token.PrimitiveVal,
		Token: value,
	}
}

func buildNullNode() *token.Node {

	return &token.Node{
		Kind:  token.Null,
		Token: nil,
	}
}
