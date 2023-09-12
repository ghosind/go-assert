package assert

import (
	"testing"
)

type testStruct struct {
	v int
}

func TestDeepEqualAndNotDeepEqual(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testDeepEqualAndNotDeepEqual(a, mockA, 1, 1, true)
	testDeepEqualAndNotDeepEqual(a, mockA, 1, 2, false)
	testDeepEqualAndNotDeepEqual(a, mockA, 1, 1.0, false)
	testDeepEqualAndNotDeepEqual(a, mockA, 1, "1", false)
	testDeepEqualAndNotDeepEqual(a, mockA, 1, '1', false)
	testDeepEqualAndNotDeepEqual(a, mockA, 1, []int{1}, false)
	testDeepEqualAndNotDeepEqual(a, mockA, []int{1}, []int{1}, true)

	obj1 := testStruct{v: 1}
	obj2 := testStruct{v: 1}

	testDeepEqualAndNotDeepEqual(a, mockA, obj1, obj2, true)
	testDeepEqualAndNotDeepEqual(a, mockA, obj1, &obj2, false)
	testDeepEqualAndNotDeepEqual(a, mockA, &obj1, &obj2, true)

	obj2.v = 2
	testDeepEqualAndNotDeepEqual(a, mockA, obj1, obj2, false)
	testDeepEqualAndNotDeepEqual(a, mockA, &obj1, &obj2, false)
}

func testDeepEqualAndNotDeepEqual(a, mockA *Assertion, v1, v2 any, isEqual bool) {
	// DeepEqual
	testAssertionFunction(a, "DeepEqual", func() error {
		return DeepEqual(mockA.T, v1, v2)
	}, isEqual)
	testAssertionFunction(a, "Assertion.DeepEqual", func() error {
		return mockA.DeepEqual(v1, v2)
	}, isEqual)

	// NotDeepEqual
	testAssertionFunction(a, "NotDeepEqual", func() error {
		return NotDeepEqual(mockA.T, v1, v2)
	}, !isEqual)
	testAssertionFunction(a, "Assertion.NotDeepEqual", func() error {
		return mockA.NotDeepEqual(v1, v2)
	}, !isEqual)

	// DeepEqualNow
	testAssertionNowFunction(a, "DeepEqualNow", func() {
		DeepEqualNow(mockA.T, v1, v2)
	}, !isEqual)
	testAssertionNowFunction(a, "Assertion.DeepEqualNow", func() {
		mockA.DeepEqualNow(v1, v2)
	}, !isEqual)

	// NotDeepEqualNow
	testAssertionNowFunction(a, "NotDeepEqualNow", func() {
		NotDeepEqualNow(mockA.T, v1, v2)
	}, isEqual)
	testAssertionNowFunction(a, "Assertion.NotDeepEqualNow", func() {
		mockA.NotDeepEqualNow(v1, v2)
	}, isEqual)
}

func TestEqualAndNotEqual(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testEqualAndNotEqual(a, mockA, 1, 1, true)
	testEqualAndNotEqual(a, mockA, 1, 2, false)
	testEqualAndNotEqual(a, mockA, 1, int64(1), true)
	testEqualAndNotEqual(a, mockA, 1, uint64(1), true)
	testEqualAndNotEqual(a, mockA, 1, 1.0, false)
	testEqualAndNotEqual(a, mockA, 1, "1", false)
	testEqualAndNotEqual(a, mockA, 1, '1', false)
	testEqualAndNotEqual(a, mockA, 1, []int{1}, false)
	testEqualAndNotEqual(a, mockA, []int{1}, []int{1}, true)

	str1 := "Hello"
	testEqualAndNotEqual(a, mockA, str1, "Hello", true)
	str1p := &str1
	testEqualAndNotEqual(a, mockA, str1p, "Hello", false)
	testEqualAndNotEqual(a, mockA, *str1p, "Hello", true)

	obj1 := testStruct{v: 1}
	obj2 := testStruct{v: 1}

	testEqualAndNotEqual(a, mockA, obj1, obj2, true)
	testEqualAndNotEqual(a, mockA, obj1, &obj2, false)
	testEqualAndNotEqual(a, mockA, &obj1, &obj2, false)

	obj2.v = 2
	testEqualAndNotEqual(a, mockA, obj1, obj2, false)
}

func testEqualAndNotEqual(a, mockA *Assertion, v1, v2 any, isEqual bool) {
	// Equal
	testAssertionFunction(a, "Equal", func() error {
		return Equal(mockA.T, v1, v2)
	}, isEqual)
	testAssertionFunction(a, "Assertion.Equal", func() error {
		return mockA.Equal(v1, v2)
	}, isEqual)

	// NotEqual
	testAssertionFunction(a, "NotEqual", func() error {
		return NotEqual(mockA.T, v1, v2)
	}, !isEqual)
	testAssertionFunction(a, "Assertion.NotEqual", func() error {
		return mockA.NotEqual(v1, v2)
	}, !isEqual)

	// EqualNow
	testAssertionNowFunction(a, "EqualNow", func() {
		EqualNow(mockA.T, v1, v2)
	}, !isEqual)
	testAssertionNowFunction(a, "Assertion.EqualNow", func() {
		mockA.EqualNow(v1, v2)
	}, !isEqual)

	// NotEqualNow
	testAssertionNowFunction(a, "NotEqualNow", func() {
		NotEqualNow(mockA.T, v1, v2)
	}, isEqual)
	testAssertionNowFunction(a, "Assertion.NotEqualNow", func() {
		mockA.NotEqualNow(v1, v2)
	}, isEqual)
}

func TestNilAndNotNil(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testNilAndNotNil(a, mockA, 1, false)
	testNilAndNotNil(a, mockA, "", false)
	testNilAndNotNil(a, mockA, nil, true)
	var testAssert *Assertion
	testNilAndNotNil(a, mockA, testAssert, true)
	testNilAndNotNil(a, mockA, mockA, false)
}

func testNilAndNotNil(a, mockA *Assertion, v any, isNil bool) {
	// Nil
	testAssertionFunction(a, "Nil", func() error {
		return Nil(mockA.T, v)
	}, isNil)
	testAssertionFunction(a, "Assertion.Nil", func() error {
		return mockA.Nil(v)
	}, isNil)

	// NotNil
	testAssertionFunction(a, "NotNil", func() error {
		return NotNil(mockA.T, v)
	}, !isNil)
	testAssertionFunction(a, "Assertion.NotNil", func() error {
		return mockA.NotNil(v)
	}, !isNil)

	// NilNow
	testAssertionNowFunction(a, "NilNow", func() {
		NilNow(mockA.T, v)
	}, !isNil)
	testAssertionNowFunction(a, "Assertion.NilNow", func() {
		mockA.NilNow(v)
	}, !isNil)

	// NotNilNow
	testAssertionNowFunction(a, "NotNilNow", func() {
		NotNilNow(mockA.T, v)
	}, isNil)
	testAssertionNowFunction(a, "Assertion.NotNilNow", func() {
		mockA.NotNilNow(v)
	}, isNil)
}

func TestTrueAndNotTrue(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testTrueAndNotTrue(a, mockA, nil, false)
	testTrueAndNotTrue(a, mockA, []int{}, false)
	testTrueAndNotTrue(a, mockA, []int{0}, true)
	testTrueAndNotTrue(a, mockA, 0, false)
	testTrueAndNotTrue(a, mockA, 1, true)
	testTrueAndNotTrue(a, mockA, 0.0, false)
	testTrueAndNotTrue(a, mockA, 1.0, true)
	testTrueAndNotTrue(a, mockA, "", false)
	testTrueAndNotTrue(a, mockA, "test", true)
	testTrueAndNotTrue(a, mockA, func() {}, true)
}

func testTrueAndNotTrue(a, mockA *Assertion, v any, isTruthy bool) {
	// True
	testAssertionFunction(a, "True", func() error {
		return True(mockA.T, v)
	}, isTruthy)
	testAssertionFunction(a, "Assertion.True", func() error {
		return mockA.True(v)
	}, isTruthy)

	// NotTrue
	testAssertionFunction(a, "NotTrue", func() error {
		return NotTrue(mockA.T, v)
	}, !isTruthy)
	testAssertionFunction(a, "Assertion.NotTrue", func() error {
		return mockA.NotTrue(v)
	}, !isTruthy)

	// TrueNow
	testAssertionNowFunction(a, "TrueNow", func() {
		TrueNow(mockA.T, v)
	}, !isTruthy)
	testAssertionNowFunction(a, "Assertion.TrueNow", func() {
		mockA.TrueNow(v)
	}, !isTruthy)

	// NotTrueNow
	testAssertionNowFunction(a, "NotTrueNow", func() {
		NotTrueNow(mockA.T, v)
	}, isTruthy)
	testAssertionNowFunction(a, "Assertion.NotTrueNow", func() {
		mockA.NotTrueNow(v)
	}, isTruthy)
}
