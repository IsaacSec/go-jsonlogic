package parser

import (
	"encoding/json"
	"log/slog"

	"github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
)

func ParseJson(data []byte) (*token.Node, error) {
	var v interface{}

	err := json.Unmarshal(data, &v)

	if err != nil {
		return nil, err
	}

	return parseValue(v), nil
}

func parseValue(value interface{}) *token.Node {

	switch val := value.(type) {
	case map[string]interface{}:
		// Determine if it is operator or object
		if len(val) == 1 {
			// Get operator key and children
			for op, children := range val {
				// Check if operator exits
				if _, ok := operators.OperatorMap[op]; ok {
					// Todo: Add cast error handler
					slog.Info("Token: %v, Childs: %v", op, children)

					switch childType := children.(type) {
					case []interface{}:
						// Create operator
						node := &token.Node{
							Kind:  token.Operator,
							Token: op,
						}

						for _, child := range children.([]interface{}) {
							// Recursion (DFS) to parse every node
							node.Childrens = append(node.Childrens, parseValue(child))
						}

						return node

					default:
						slog.Error("Error parsing: [%v] %v -> %v", childType, op, children)

						return &token.Node{
							Kind:  token.Object,
							Token: val,
						}
					}
				}
			}
		}

		return &token.Node{
			Kind:  token.Object,
			Token: val,
		}

	case []interface{}:
		node := &token.Node{
			Kind:      token.Array,
			Childrens: make([]*token.Node, len(val)),
			Token:     val,
		}
		for _, v := range val {
			// Recursion (DFS) to parse every node
			node.Childrens = append(node.Childrens, parseValue(v))
		}
		return node
	case string, float64, bool, int:
		return &token.Node{Kind: token.PrimitiveVal, Token: val}
	default:
		return &token.Node{Kind: token.Null, Token: nil}
	}
}
