package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatConversionWithIncompatibleValues(t *testing.T) {
	_, _, err := ConvertToFloat("one", 45)

	assert.Error(t, err, "expected error on invalid conversion")
}

func TestFloatConversionWithSameType(t *testing.T) {
	a, b, err := ConvertToFloat(123, "123.33")

	assert.NoError(t, err)
	assert.Equal(t, 123.0, a)
	assert.Equal(t, 123.33, b)
}
