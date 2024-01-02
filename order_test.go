package assert

import (
	"reflect"
	"testing"
)

func TestGt(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	a.PanicNow(func() {
		mockA.Gt(1, uint(1))
	})
	a.PanicNow(func() {
		mockA.Gt(true, true)
	})

	testGt(a, mockA, 1, 1, false)
	testGt(a, mockA, 2, 1, true)
	testGt(a, mockA, 1, int64(2), false)
	testGt(a, mockA, int64(2), 1, true)
	testGt(a, mockA, uint(1), uint64(2), false)
	testGt(a, mockA, uint(2), uint64(1), true)
	testGt(a, mockA, float32(2.0), 1.0, true)
	testGt(a, mockA, "ABC", "BCD", false)
	testGt(a, mockA, "ABC", "AAA", true)
}

func testGt(a, mockA *Assertion, v1, v2 any, isGt bool) {
	a.T.Helper()

	testAssertionFunction(a, "Gt", func() error {
		return Gt(mockA.T, v1, v2)
	}, isGt)
	testAssertionFunction(a, "Assertion.Gt", func() error {
		return mockA.Gt(v1, v2)
	}, isGt)

	testAssertionNowFunction(a, "GtNow", func() {
		GtNow(mockA.T, v1, v2)
	}, !isGt)
	testAssertionNowFunction(a, "Assertion.GtNow", func() {
		mockA.GtNow(v1, v2)
	}, !isGt)
}

func TestGte(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	a.PanicNow(func() {
		mockA.Gte(1, uint(1))
	})
	a.PanicNow(func() {
		mockA.Gte(true, true)
	})

	testGte(a, mockA, 0, 1, false)
	testGte(a, mockA, 1, 1, true)
	testGte(a, mockA, 2, 1, true)
	testGte(a, mockA, 1, int64(2), false)
	testGte(a, mockA, int64(2), 1, true)
	testGte(a, mockA, int64(2), 2, true)
	testGte(a, mockA, uint(1), uint64(2), false)
	testGte(a, mockA, uint(2), uint64(1), true)
	testGte(a, mockA, float32(2.0), 1.0, true)
	testGte(a, mockA, "ABC", "BCD", false)
	testGte(a, mockA, "ABC", "AAA", true)
	testGte(a, mockA, "ABC", "ABC", true)
}

func testGte(a, mockA *Assertion, v1, v2 any, isGte bool) {
	a.T.Helper()

	testAssertionFunction(a, "Gte", func() error {
		return Gte(mockA.T, v1, v2)
	}, isGte)
	testAssertionFunction(a, "Assertion.Gte", func() error {
		return mockA.Gte(v1, v2)
	}, isGte)

	testAssertionNowFunction(a, "GteNow", func() {
		GteNow(mockA.T, v1, v2)
	}, !isGte)
	testAssertionNowFunction(a, "Assertion.GteNow", func() {
		mockA.GteNow(v1, v2)
	}, !isGte)
}

func TestLt(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	a.PanicNow(func() {
		mockA.Lt(1, uint(1))
	})
	a.PanicNow(func() {
		mockA.Lt(true, true)
	})

	testLt(a, mockA, 1, 1, false)
	testLt(a, mockA, 1, 2, true)
	testLt(a, mockA, 2, int64(1), false)
	testLt(a, mockA, int64(1), 2, true)
	testLt(a, mockA, uint(2), uint64(1), false)
	testLt(a, mockA, uint(1), uint64(2), true)
	testLt(a, mockA, float32(1.0), 2.0, true)
	testLt(a, mockA, "BCD", "ABC", false)
	testLt(a, mockA, "AAA", "ABC", true)
}

func testLt(a, mockA *Assertion, v1, v2 any, isLt bool) {
	a.T.Helper()

	testAssertionFunction(a, "Lt", func() error {
		return Lt(mockA.T, v1, v2)
	}, isLt)
	testAssertionFunction(a, "Assertion.Lt", func() error {
		return mockA.Lt(v1, v2)
	}, isLt)

	testAssertionNowFunction(a, "LtNow", func() {
		LtNow(mockA.T, v1, v2)
	}, !isLt)
	testAssertionNowFunction(a, "Assertion.LtNow", func() {
		mockA.LtNow(v1, v2)
	}, !isLt)
}

func TestLte(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	a.PanicNow(func() {
		mockA.Lte(1, uint(1))
	})
	a.PanicNow(func() {
		mockA.Lte(true, true)
	})

	testLte(a, mockA, 1, 0, false)
	testLte(a, mockA, 1, 1, true)
	testLte(a, mockA, 1, 2, true)
	testLte(a, mockA, 2, int64(1), false)
	testLte(a, mockA, int64(1), 2, true)
	testLte(a, mockA, int64(2), 2, true)
	testLte(a, mockA, uint(2), uint64(1), false)
	testLte(a, mockA, uint(1), uint64(2), true)
	testLte(a, mockA, float32(1.0), 2.0, true)
	testLte(a, mockA, "BCD", "ABC", false)
	testLte(a, mockA, "AAA", "ABC", true)
	testLte(a, mockA, "ABC", "ABC", true)
}

func testLte(a, mockA *Assertion, v1, v2 any, isLte bool) {
	a.T.Helper()

	testAssertionFunction(a, "Lte", func() error {
		return Lte(mockA.T, v1, v2)
	}, isLte)
	testAssertionFunction(a, "Assertion.Lte", func() error {
		return mockA.Lte(v1, v2)
	}, isLte)

	testAssertionNowFunction(a, "LteNow", func() {
		LteNow(mockA.T, v1, v2)
	}, !isLte)
	testAssertionNowFunction(a, "Assertion.LteNow", func() {
		mockA.LteNow(v1, v2)
	}, !isLte)
}

func TestCompareValues(t *testing.T) {
	a := New(t)

	a.True(compareValues(reflect.ValueOf(1), reflect.ValueOf(1), compareTypeEqual))
	a.NotTrue(compareValues(reflect.ValueOf([]int{1}), reflect.ValueOf([]int{1}), compareTypeEqual))
}

func TestIsOrderable(t *testing.T) {
	assert := New(t)

	assert.Equal(isOrderable(1), true)
	assert.Equal(isOrderable(int64(1)), true)
	assert.Equal(isOrderable(uint64(1)), true)
	assert.Equal(isOrderable(float32(1.0)), true)
	assert.Equal(isOrderable(1.0), true)
	assert.Equal(isOrderable("Hello"), true)
	assert.Equal(isOrderable([]byte{'H', 'e', 'l', 'l', 'o'}), false)
	assert.Equal(isOrderable([]int{1, 2, 3}), false)
}
