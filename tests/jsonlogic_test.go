package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/IsaacSec/go-jsonlogic/jsonlogic"
	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	tests := loadTestCases(t, "equals.json")

	assertTestCases(t, tests)
}

func TestAnd(t *testing.T) {
	tests := loadTestCases(t, "and.json")

	assertTestCases(t, tests)
}

func loadTestCases(t *testing.T, filename string) []any {
	file, err := os.Open("resources/" + filename)

	assert.NoError(t, err)

	defer file.Close()

	var testCases []any

	json.NewDecoder(file).Decode(&testCases)

	return testCases
}

func assertTestCases(t *testing.T, tests []any) {
	for _, test := range tests {
		jsonObject := (test.(map[string]any))

		rules := jsonObject["rules"]
		expected := jsonObject["expected"]

		result, err := jsonlogic.Apply(rules, nil)

		assert.NoError(t, err)
		assert.Equal(t, expected, result, fmt.Sprintf("comparison: %v", rules))
	}
}
