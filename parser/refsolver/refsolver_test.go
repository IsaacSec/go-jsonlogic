package refsolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReferenceValueFromMap(t *testing.T) {
	json := map[string]any{
		"foo": 1,
		"bar": "test",
	}

	res := GetValue(json, "foo")
	assert.Equal(t, 1, res)
}

func TestGetFallbackOnMissingReference(t *testing.T) {
	json := map[string]any{
		"bar": "test",
	}

	res := GetValue(json, []string{"foo", "fallback"})
	assert.Equal(t, "fallback", res)
}

func TestGetValueFromNestedReference(t *testing.T) {
	json := map[string]any{
		"foo": 1,
		"user": map[string]any{
			"name": map[string]any{
				"first": "Juan",
			},
		},
	}

	res := GetValue(json, "user.name.first")
	assert.Equal(t, "Juan", res)
}

func TestGetFallbackOnMissingWithNestedReference(t *testing.T) {
	json := map[string]any{
		"foo": 1,
		"user": map[string]any{
			"name": map[string]any{
				"second": "Juan",
			},
		},
	}

	res := GetValue(json, []string{"user.name.first", "John"})
	assert.Equal(t, "John", res)
}

func TestGetValueOnMissingFallback(t *testing.T) {
	json := map[string]any{
		"foo": 1,
	}

	res := GetValue(json, []string{"user.name.first"})
	assert.Nil(t, res)
}

func TestGetValueWithInvalidReference(t *testing.T) {
	json := map[string]any{
		"foo": 1,
	}

	res := GetValue(json, true)
	assert.Nil(t, res)
}
