package assert

import (
	"testing"
)

func TestMapHasKey(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	tesMapHasKeyAndNotMapHasKey(a, mockA, map[string]int{"a": 1, "b": 2}, "a", true)
	tesMapHasKeyAndNotMapHasKey(a, mockA, map[string]int{"a": 1, "b": 2}, "b", true)
	tesMapHasKeyAndNotMapHasKey(a, mockA, map[string]int{"a": 1, "b": 2}, "c", false)
	tesMapHasKeyAndNotMapHasKey(a, mockA, map[string]int{"a": 1, "b": 2}, 1, false)
	tesMapHasKeyAndNotMapHasKey(a, mockA, map[any]any{"a": "b", 2: 1}, "a", true)
	tesMapHasKeyAndNotMapHasKey(a, mockA, map[any]any{"a": "b", 2: 1}, 2, true)
	tesMapHasKeyAndNotMapHasKey(a, mockA, map[any]any{"a": "b", 2: 1}, true, false)
}

func tesMapHasKeyAndNotMapHasKey(
	a, mockA *Assertion,
	m, key any,
	isHasKey bool,
) {
	// MapHasKey
	testAssertionFunction(a, "MapHasKey", func() error {
		return MapHasKey(mockA.T, m, key)
	}, isHasKey)
	testAssertionFunction(a, "Assertion.MapHasKey", func() error {
		return mockA.MapHasKey(m, key)
	}, isHasKey)

	// NotMapHasKey
	testAssertionFunction(a, "NotMapHasKey", func() error {
		return NotMapHasKey(mockA.T, m, key)
	}, !isHasKey)
	testAssertionFunction(a, "Assertion.NotMapHasKey", func() error {
		return mockA.NotMapHasKey(m, key)
	}, !isHasKey)

	// MapHasKeyNow
	testAssertionNowFunction(a, "MapHasKeyNow", func() {
		MapHasKeyNow(mockA.T, m, key)
	}, !isHasKey)
	testAssertionNowFunction(a, "Assertion.MapHasKeyNow", func() {
		mockA.MapHasKeyNow(m, key)
	}, !isHasKey)

	// NotMapHasKeyNow
	testAssertionNowFunction(a, "NotMapHasKeyNow", func() {
		NotMapHasKeyNow(mockA.T, m, key)
	}, isHasKey)
	testAssertionNowFunction(a, "Assertion.NotMapHasKeyNow", func() {
		mockA.NotMapHasKeyNow(m, key)
	}, isHasKey)
}

func TestMapHasValue(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	tesMapHasValueAndNotMapHasValue(a, mockA, map[string]int{"a": 1, "b": 2}, 1, true)
	tesMapHasValueAndNotMapHasValue(a, mockA, map[string]int{"a": 1, "b": 2}, 2, true)
	tesMapHasValueAndNotMapHasValue(a, mockA, map[string]int{"a": 1, "b": 2}, 3, false)
	tesMapHasValueAndNotMapHasValue(a, mockA, map[string]int{"a": 1, "b": 2}, "a", false)
	tesMapHasValueAndNotMapHasValue(a, mockA, map[any]any{"a": "b", 2: 1}, "b", true)
	tesMapHasValueAndNotMapHasValue(a, mockA, map[any]any{"a": "b", 2: 1}, 1, true)
	tesMapHasValueAndNotMapHasValue(a, mockA, map[any]any{"a": "b", 2: 1}, true, false)
}

func tesMapHasValueAndNotMapHasValue(
	a, mockA *Assertion,
	m, key any,
	isHasValue bool,
) {
	// MapHasValue
	testAssertionFunction(a, "MapHasValue", func() error {
		return MapHasValue(mockA.T, m, key)
	}, isHasValue)
	testAssertionFunction(a, "Assertion.MapHasValue", func() error {
		return mockA.MapHasValue(m, key)
	}, isHasValue)

	// NotMapHasValue
	testAssertionFunction(a, "NotMapHasValue", func() error {
		return NotMapHasValue(mockA.T, m, key)
	}, !isHasValue)
	testAssertionFunction(a, "Assertion.NotMapHasValue", func() error {
		return mockA.NotMapHasValue(m, key)
	}, !isHasValue)

	// MapHasValueNow
	testAssertionNowFunction(a, "MapHasValueNow", func() {
		MapHasValueNow(mockA.T, m, key)
	}, !isHasValue)
	testAssertionNowFunction(a, "Assertion.MapHasValueNow", func() {
		mockA.MapHasValueNow(m, key)
	}, !isHasValue)

	// NotMapHasValueNow
	testAssertionNowFunction(a, "NotMapHasValueNow", func() {
		NotMapHasValueNow(mockA.T, m, key)
	}, isHasValue)
	testAssertionNowFunction(a, "Assertion.NotMapHasValueNow", func() {
		mockA.NotMapHasValueNow(m, key)
	}, isHasValue)
}

func TestIsMapHasKey(t *testing.T) {
	assert := New(t)

	assert.NotTrue(isMapHasKey(nil, nil))
	assert.NotTrue(isMapHasKey(map[string]int{}, "a"))
	assert.True(isMapHasKey(map[string]int{
		"a": 1,
		"b": 2,
	}, "a"))
	assert.NotTrue(isMapHasKey(map[string]int{
		"a": 1,
		"b": 2,
	}, "c"))
	assert.NotTrue(isMapHasKey(map[string]int{
		"a": 1,
		"b": 2,
	}, 1))
	assert.True(isMapHasKey(map[any]int{
		"a": 1,
		1:   2,
	}, 1))
	assert.True(isMapHasKey(map[any]int{
		"a": 1,
		1:   2,
	}, "a"))
	assert.NotTrue(isMapHasKey(map[any]int{
		"a": 1,
		1:   2,
	}, 2))
	assert.NotTrue(isMapHasKey(map[any]int{
		"a": 1,
		1:   2,
	}, "b"))
	assert.NotTrue(isMapHasKey(map[any]int{
		"a": 1,
		1:   2,
	}, 1.1))
}

func TestIsMapHasValue(t *testing.T) {
	assert := New(t)

	assert.NotTrue(isMapHasValue(nil, nil))
	assert.NotTrue(isMapHasValue(map[string]int{}, 3))
	assert.True(isMapHasValue(map[string]int{
		"a": 1,
		"b": 2,
	}, 1))
	assert.NotTrue(isMapHasValue(map[string]int{
		"a": 1,
		"b": 2,
	}, 3))
	assert.NotTrue(isMapHasValue(map[string]int{
		"a": 1,
		"b": 2,
	}, true))
	assert.True(isMapHasValue(map[any]any{
		"a": "b",
		1:   2,
	}, "b"))
	assert.True(isMapHasValue(map[any]any{
		"a": "b",
		1:   2,
	}, 2))
	assert.NotTrue(isMapHasValue(map[any]any{
		"a": "b",
		1:   2,
	}, "a"))
	assert.NotTrue(isMapHasValue(map[any]any{
		"a": "b",
		1:   2,
	}, 1))
	assert.NotTrue(isMapHasValue(map[any]any{
		"a": "b",
		1:   2,
	}, 1.1))
}
