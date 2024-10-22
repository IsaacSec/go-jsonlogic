package jsonlogic

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnEvaluationTree(t *testing.T) {

	var rules map[string]any

	json.Unmarshal([]byte(`
		{
			"and" : [
				{ 
					"or": [
						{">": [1, 999]},
						{"==": [1, true]}
					]
				},
				{ "==": ["string","string"] },
				{ "!=": [1,0] },
				{ "<": [0,1] },
				{ "<=": [0,0] },
				{ ">": [1,0] },
				{ ">=": [1,1] }
			]
		}
	`), &rules)

	assert.NotNil(t, rules, "cannot parse json")

	res, err1 := ApplyTree(rules, nil)
	assert.NoError(t, err1)

	result, err2 := json.Marshal(res)
	assert.NoError(t, err2)

	fmt.Printf("Tree: %+v\n", string(result))
}
