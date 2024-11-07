package varref

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

func TestGetReferenceValueFromArray(t *testing.T) {
	json := []any{"zero", "one", "two", "three"}

	res := GetValue(json, 1)
	assert.Equal(t, "one", res)
}

func TestGetReferenceValueFromArrayWhenOutOfBound(t *testing.T) {
	json := []any{"zero", "one", "two"}

	lower := GetValue(json, -1)
	upper := GetValue(json, 4)

	assert.Nil(t, lower)
	assert.Nil(t, upper)
}

func TestGetFallbackOnMissingReference(t *testing.T) {
	json := map[string]any{
		"bar": "test",
	}

	res := GetValue(json, []any{"foo", "fallback"})
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

	res := GetValue(json, []any{"user.name.first", "John"})
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
