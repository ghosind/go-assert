package assert

import (
	"testing"
)

func TestElementContainsAndNotContains(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testElementContainsAndNotContains(a, mockA, []int{1}, 1, true)
	testElementContainsAndNotContains(a, mockA, []int{1, 2, 3}, 3, true)
	testElementContainsAndNotContains(a, mockA, []int{1, 2, 3}, 4, false)
	testElementContainsAndNotContains(a, mockA, []int{}, 1, false)
	testElementContainsAndNotContains(a, mockA, [1]int{1}, 1, true)
	testElementContainsAndNotContains(a, mockA, [3]int{1, 2, 3}, 3, true)
	testElementContainsAndNotContains(a, mockA, [3]int{1, 2, 3}, 4, false)
	testElementContainsAndNotContains(a, mockA, [0]int{}, 1, false)
}

func testElementContainsAndNotContains(
	a, mockA *Assertion,
	source, expect any,
	isContains bool,
) {
	// ContainsElement
	testAssertionFunction(a, "ContainsElement", func() error {
		return ContainsElement(mockA.T, source, expect)
	}, isContains)
	testAssertionFunction(a, "Assertion.ContainsElement", func() error {
		return mockA.ContainsElement(source, expect)
	}, isContains)

	// NotContainsElement
	testAssertionFunction(a, "NotContainsElement", func() error {
		return NotContainsElement(mockA.T, source, expect)
	}, !isContains)
	testAssertionFunction(a, "Assertion.NotContainsElement", func() error {
		return mockA.NotContainsElement(source, expect)
	}, !isContains)

	// ContainsElementNow
	testAssertionNowFunction(a, "ContainsElementNow", func() {
		ContainsElementNow(mockA.T, source, expect)
	}, !isContains)
	testAssertionNowFunction(a, "Assertion.ContainsElementNow", func() {
		mockA.ContainsElementNow(source, expect)
	}, !isContains)

	// NotContainsElementNow
	testAssertionNowFunction(a, "NotContainsElementNow", func() {
		NotContainsElementNow(mockA.T, source, expect)
	}, isContains)
	testAssertionNowFunction(a, "Assertion.NotContainsElementNow", func() {
		mockA.NotContainsElementNow(source, expect)
	}, isContains)
}
