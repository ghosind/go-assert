package assert

import (
	"math"
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestFailedHandler(t *testing.T) {
	mockT := new(testing.T)
	assert := New(t)

	failed(mockT, nil, false)
	assert.NotTrue(mockT.Failed(), false)

	failed(mockT, newAssertionError("Test error"), false)
	assert.True(mockT.Failed())

	isTerminated := internal.CheckTermination(func() {
		failed(mockT, newAssertionError("Test error"), true)
	})
	assert.True(isTerminated)
}

func TestIsContainsElement(t *testing.T) {
	assert := New(t)

	assert.PanicNow(func() {
		isContainsElement("not array or slice", 1)
	})
	assert.PanicNow(func() {
		isContainsElement([]string{"a", "b", "c"}, 1)
	})

	assert.NotTrueNow(isContainsElement([]string{}, "c"))
	assert.TrueNow(isContainsElement([]string{"a", "b", "c"}, "c"))
	assert.NotTrueNow(isContainsElement([]string{"a", "b", "c"}, "d"))
	assert.TrueNow(isContainsElement([]int{1, 2, 3}, 3))
	assert.NotTrueNow(isContainsElement([]int{1, 2, 3}, 4))
	assert.TrueNow(isContainsElement([]int64{1, 2, 3}, 3))
	assert.NotTrueNow(isContainsElement([]int64{1, 2, 3}, 4))
	assert.TrueNow(isContainsElement([]uint64{1, 2, 3}, 3))
	assert.NotTrueNow(isContainsElement([]uint64{1, 2, 3}, 4))
	assert.TrueNow(isContainsElement(&[]int{1, 2, 3}, 3))
	assert.TrueNow(isContainsElement([3]int{1, 2, 3}, 3))
	assert.NotTrueNow(isContainsElement([3]int{1, 2, 3}, 4))
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

	assert.Equal(isEqual(nil, nil), true)
	assert.Equal(isEqual(nil, s1), false) // s1 is nil
	assert.Equal(isEqual(true, false), false)
	assert.Equal(isEqual(1, 1), true)
	assert.Equal(isEqual(1, 2), false)
	assert.Equal(isEqual(1, int64(1)), true)
	assert.Equal(isEqual(1, int64(2)), false)
	assert.Equal(isEqual(uint(1), uint(1)), true)
	assert.Equal(isEqual(uint(1), uint(2)), false)
	assert.Equal(isEqual(uint(1), uint64(1)), true)
	assert.Equal(isEqual(uint(1), uint64(2)), false)
	assert.Equal(isEqual(uint(1), uintptr(1)), true)
	assert.Equal(isEqual(1.0, 1.0), true)
	assert.Equal(isEqual(1.0, 2.0), false)
	assert.Equal(isEqual(1.0, float32(1.0)), true)
	assert.Equal(isEqual(1.0, float32(2.0)), false)
	assert.Equal(isEqual(complex(1, 1), complex(1, 1)), true)
	assert.Equal(isEqual(complex(1, 1), complex(2, 2)), false)
	assert.Equal(isEqual(complex(1, 1), complex64(complex(1, 1))), true)
	assert.Equal(isEqual(complex(1, 1), complex64(complex(2, 2))), false)
	assert.Equal(isEqual([1]int{0}, [1]int{0}), true)
	assert.Equal(isEqual([1]int{0}, [1]int{1}), false)
	assert.Equal(isEqual([1]int{0}, [2]int{0, 0}), false)
	assert.Equal(isEqual([1]int{0}, [1]float64{0.0}), false)
	assert.Equal(isEqual("hello", "hello"), true)
	assert.Equal(isEqual("hello", "world"), false)

	slice1 := []int{0}
	slice2 := []int{0}
	slice3 := []int{0, 0}
	slice4 := []int{1}
	slice5 := []float64{0.0}
	assert.Equal(isEqual(slice1, slice1), true)
	assert.Equal(isEqual(slice1, slice2), true)
	assert.Equal(isEqual(slice1, slice3), false)
	assert.Equal(isEqual(slice1, slice4), false)
	assert.Equal(isEqual(slice1, slice5), false)

	assert.Equal(isEqual(testStruct1{A: 0}, testStruct1{A: 0}), true)
	assert.Equal(isEqual(testStruct1{A: 0}, testStruct1{A: 1}), false)
	assert.Equal(isEqual(s1, s1), true)
	assert.Equal(isEqual(&testStruct1{A: 0}, &testStruct1{A: 1}), false)
	assert.Equal(isEqual(testStruct1{A: 0}, testStruct2{A: 0}), false)
}

func TestIsEqualOfMixSignInt(t *testing.T) {
	assert := New(t)

	assert.True(isEqual(0, uint(0)))
	assert.True(isEqual(1, uint(1)))
	assert.True(isEqual(uint(1), 1))
	assert.True(isEqual(math.MaxInt64, uint64(math.MaxInt64)))
	assert.NotTrue(isEqual(-1, uint64(math.MaxUint64)))
	assert.NotTrue(isEqual(uint64(math.MaxUint64), -1))
	assert.NotTrue(isEqual(uint64(math.MaxUint64), 0))
}

func TestIsComparable(t *testing.T) {
	assert := New(t)

	assert.Equal(isComparable(1), true)
	assert.Equal(isComparable(int64(1)), true)
	assert.Equal(isComparable(uint64(1)), true)
	assert.Equal(isComparable(float32(1.0)), true)
	assert.Equal(isComparable(1.0), true)
	assert.Equal(isComparable("Hello"), true)
	assert.Equal(isComparable([]byte{'H', 'e', 'l', 'l', 'o'}), false)
	assert.Equal(isComparable([]int{1, 2, 3}), false)
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

func TestIsPanic(t *testing.T) {
	Nil(t, isPanic(func() {
		// no panic
	}))
	NotNil(t, isPanic(func() {
		panic("unexpected panic")
	}))
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
