package parser

import (
	"encoding/json"

	"github.com/IsaacSec/go-jsonlogic/operators"
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
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
				node := &token.Node{
					Kind:  token.ReferenceVal,
					Token: ref,
				}

				return node
			}

			// Get operator key and children
			for op, children := range val {
				// Check if operator exits
				if _, ok := operators.Operators[op]; ok {
					// Todo: Add cast error handler
					log.Info("Token: %v, Children: %v", op, children)

					switch childType := children.(type) {
					case []interface{}:
						// Create operator
						node := &token.Node{
							Kind:  token.Operator,
							Token: op,
						}

						for _, child := range children.([]interface{}) {
							// Recursion (DFS) to parse every node
							node.Childrens = append(node.Childrens, parse(child))
						}

						return node

					default:
						log.Error("Error parsing: [%v] %v -> %v", childType, op, children)

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
			node.Childrens = append(node.Childrens, parse(v))
		}
		return node
	case string, float64, bool, int:
		return &token.Node{Kind: token.PrimitiveVal, Token: val}
	default:
		return &token.Node{Kind: token.Null, Token: nil}
	}
}
