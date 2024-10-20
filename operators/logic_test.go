package operators

import (
	"testing"
)

/***********************
 *		EQUALS
 ***********************/

// EQUALS evaluation, same type equal value -> 53 == 53
func TestEqualsOnSameTypeAndValue(t *testing.T) {
	expression := buildSimpleExp("==", 53, 53)

	assertExpression(t, expression, true)
}

// EQUALS evaluation, different type equal value -> 53 == "53"
func TestEqualsOnDiffType(t *testing.T) {
	expression := buildSimpleExp("==", 53, "fifty three")

	assertExpression(t, expression, false)
}

// EQUALS evaluation, different value -> 3421 == -123
func TestEqualsOnDiffValue(t *testing.T) {
	expression := buildSimpleExp("==", 3421, -123)

	assertExpression(t, expression, false)
}

/***********************
 *		NOT EQUALS
 ***********************/

// NOT EQUALS evaluation, same type and different value -> 35 != 53
func TestNotEqualsWithDiffValue(t *testing.T) {
	expression := buildSimpleExp("!=", 35, 53)

	assertExpression(t, expression, true)
}

// NOT EQUALS evaluation, same type and equal value -> 53 != 53
func TestNotEqualsWithSameValue(t *testing.T) {
	expression := buildSimpleExp("!=", 53, 53)

	assertExpression(t, expression, false)
}

/***********************
 *		AND
 ***********************/

// AND evaluation, empty list -> [] = true
func TestAndWithEmptyList(t *testing.T) {
	expression := buildGroupExp(
		"and",
	)

	assertExpression(t, expression, true)
}

// AND evaluation, one true and multiple false -> [true, false, false] = false
func TestAndWithMultipleFalseAndAtLeastOneTrue(t *testing.T) {
	expression := buildGroupExp(
		"and",
		true,
		false,
		false,
	)

	assertExpression(t, expression, false)
}

// AND evaluation, all true -> [true, true, true] = true
func TestAndWithAllTrue(t *testing.T) {
	expression := buildGroupExp(
		"and",
		true,
		true,
		true,
	)

	assertExpression(t, expression, true)
}

// AND evaluation, all false -> [false, false, false] = false
func TestAndWithAllFalse(t *testing.T) {
	expression := buildGroupExp(
		"and",
		false,
		false,
		false,
	)

	assertExpression(t, expression, false)
}

/***********************
 *		OR
 ***********************/

// OR evaluation, empty list -> [] = true
func TestOrWithEmptyList(t *testing.T) {
	expression := buildGroupExp(
		"or",
	)

	assertExpression(t, expression, true)
}

// OR evaluation, one true and multiple false -> [false, true, false] = true
func TestOrWithMultipleFalseAndAtLeastOneTrue(t *testing.T) {
	expression := buildGroupExp(
		"or",
		false,
		true,
		false,
	)

	assertExpression(t, expression, true)
}

// AND evaluation, all true -> [true, true, true] = true
func TestOrWithAllTrue(t *testing.T) {
	expression := buildGroupExp(
		"or",
		true,
		true,
		true,
	)

	assertExpression(t, expression, true)
}

// AND evaluation, all false -> [false, false, false] = false
func TestOrWithAllFalse(t *testing.T) {
	expression := buildGroupExp(
		"or",
		false,
		false,
		false,
	)

	assertExpression(t, expression, false)
}
