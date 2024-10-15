package util

import (
	"fmt"
	"testing"
)

func TestFloatConversionWithIncompatibleValues(t *testing.T) {
	a, b, err := ConvertToFloat("one", 45)

	fmt.Printf("Values: %v %v %v\n", a, b, err)

	if err == nil {
		t.Errorf("Expected error on invalid conversion\n")
	}
}

func TestFloatConversionWithSameType(t *testing.T) {
	a, b, err := ConvertToFloat(123, "123.33")

	fmt.Printf("Values: %v %v %v\n", a, b, err)

	if err != nil {
		t.Errorf("Unexpected error: %s\n", err)
	}

	if a != 123.0 || b != 123.33 {
		t.Errorf("Conversion error: expected %v,%v but got %v,%v\n", 123.0, 123.33, a, b)
	}
}
