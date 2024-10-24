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
	a.T.Helper()

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
	testEqualAndNotEqual(a, mockA, 1, uint64(1), false)
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
	a.T.Helper()

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

func TestFloatEqualAndFloatNotEqual(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testFloatEqualAndFloatNotEqual(a, mockA, 1, 1, 1e-7, true)
	testFloatEqualAndFloatNotEqual(a, mockA, 1, 1.1, 1e-7, false)
	testFloatEqualAndFloatNotEqual(a, mockA, 1, 2, 1e-7, false)
	testFloatEqualAndFloatNotEqual(a, mockA, 0.999999999999, 1, 1e-7, true)
	testFloatEqualAndFloatNotEqual(a, mockA, 1.00000000001, 1, 1e-7, true)
	testFloatEqualAndFloatNotEqual(a, mockA, 1.00001, 1, 1e-7, false)
	testFloatEqualAndFloatNotEqual(a, mockA, 0.9999, 1, 1e-7, false)
}

func testFloatEqualAndFloatNotEqual(a, mockA *Assertion, v1, v2, epsilon any, isEqual bool) {
	a.T.Helper()

	// FloatEqual
	testAssertionFunction(a, "FloatEqual", func() error {
		return FloatEqual(mockA.T, v1, v2, epsilon)
	}, isEqual)
	testAssertionFunction(a, "Assertion.FloatEqual", func() error {
		return mockA.FloatEqual(v1, v2, epsilon)
	}, isEqual)

	// FloatNotEqual
	testAssertionFunction(a, "FloatNotEqual", func() error {
		return FloatNotEqual(mockA.T, v1, v2, epsilon)
	}, !isEqual)
	testAssertionFunction(a, "Assertion.FloatNotEqual", func() error {
		return mockA.FloatNotEqual(v1, v2, epsilon)
	}, !isEqual)

	// FloatEqualNow
	testAssertionNowFunction(a, "FloatEqualNow", func() {
		FloatEqualNow(mockA.T, v1, v2, epsilon)
	}, !isEqual)
	testAssertionNowFunction(a, "Assertion.FloatEqualNow", func() {
		mockA.FloatEqualNow(v1, v2, epsilon)
	}, !isEqual)

	// FloatNotEqualNow
	testAssertionNowFunction(a, "FloatNotEqualNow", func() {
		FloatNotEqualNow(mockA.T, v1, v2, epsilon)
	}, isEqual)
	testAssertionNowFunction(a, "Assertion.FloatNotEqualNow", func() {
		mockA.FloatNotEqualNow(v1, v2, epsilon)
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
	a.T.Helper()

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
	a.T.Helper()

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

func TestIsEqual(t *testing.T) {
	assert := New(t)

	type testStruct1 struct {
		A int
	}
	type testStruct2 struct {
		A int
	}

	var s1 *testStruct1

	assert.True(isEqual(nil, nil))
	assert.NotTrue(isEqual(nil, s1)) // s1 is nil
	assert.NotTrue(isEqual(true, false))
	assert.True(isEqual(1, 1))
	assert.NotTrue(isEqual(1, 2))
	assert.True(isEqual(1, int64(1)))
	assert.NotTrue(isEqual(1, int64(2)))
	assert.True(isEqual(uint(1), uint(1)))
	assert.NotTrue(isEqual(uint(1), uint(2)))
	assert.True(isEqual(uint(1), uint64(1)))
	assert.NotTrue(isEqual(uint(1), uint64(2)))
	assert.True(isEqual(uint(1), uintptr(1)))
	assert.True(isEqual(1.0, 1.0))
	assert.NotTrue(isEqual(1.0, 2.0))
	assert.True(isEqual(1.0, float32(1.0)))
	assert.NotTrue(isEqual(1.0, float32(2.0)))
	assert.True(isEqual(complex(1, 1), complex(1, 1)))
	assert.NotTrue(isEqual(complex(1, 1), complex(2, 2)))
	assert.True(isEqual(complex(1, 1), complex64(complex(1, 1))))
	assert.NotTrue(isEqual(complex(1, 1), complex64(complex(2, 2))))
	assert.True(isEqual([1]int{0}, [1]int{0}))
	assert.NotTrue(isEqual([1]int{0}, [1]int{1}))
	assert.NotTrue(isEqual([1]int{0}, [2]int{0, 0}))
	assert.NotTrue(isEqual([1]int{0}, [1]float64{0.0}))
	assert.True(isEqual("hello", "hello"))
	assert.NotTrue(isEqual("hello", "world"))

	slice1 := []int{0}
	slice2 := []int{0}
	slice3 := []int{0, 0}
	slice4 := []int{1}
	slice5 := []float64{0.0}
	assert.True(isEqual(slice1, slice1))
	assert.True(isEqual(slice1, slice2))
	assert.NotTrue(isEqual(slice1, slice3))
	assert.NotTrue(isEqual(slice1, slice4))
	assert.NotTrue(isEqual(slice1, slice5))

	assert.True(isEqual([][]any{{1}, {2, 3}}, [][]any{{1}, {2, 3}}))
	assert.NotTrue(isEqual([][]any{{1}, {2, 3}}, [][]any{{1.0}, {2.0, 3.0}}))
	assert.NotTrue(isEqual([][]any{{1}, {2, 3}}, [][]any{{"1"}, {"2", "3"}}))
	assert.True(isEqual([][][]any{{{1}, {2}}, {{2, 3}}}, [][][]any{{{1}, {2}}, {{2, 3}}}))
	assert.NotTrue(isEqual([][][]any{{{1}, {2}}, {{2, 3}}}, [][][]any{{{1}, {2}}, {{2, "3"}}}))

	assert.True(isEqual(testStruct1{A: 0}, testStruct1{A: 0}))
	assert.NotTrue(isEqual(testStruct1{A: 0}, testStruct1{A: 1}))
	assert.True(isEqual(s1, s1))
	assert.NotTrue(isEqual(&testStruct1{A: 0}, &testStruct1{A: 1}))
	assert.NotTrue(isEqual(testStruct1{A: 0}, testStruct2{A: 0}))
}

func TestIsNil(t *testing.T) {
	assert := New(t)

	assert.NotTrue(isNil(1))  // int
	assert.NotTrue(isNil("")) // string
	assert.True(isNil(nil))
	var testAssert *Assertion
	assert.True(isNil(testAssert))
	assert.NotTrue(isNil(assert))
}

func TestIsTrue(t *testing.T) {
	assert := New(t)

	// reflect.Invalid
	assert.NotTrue(isTrue(nil))

	// reflect.Slice
	assert.True(isTrue([]int{0}))
	assert.NotTrue(isTrue([]int{}))

	// other kinds
	assert.True(isTrue(1))
	assert.NotTrue(isTrue(0))
	assert.True(isTrue(1.0))
	assert.NotTrue(isTrue(0.0))
	assert.True(isTrue("Hello"))
	assert.NotTrue(isTrue(""))
	assert.True(isTrue(func() {}))
}
