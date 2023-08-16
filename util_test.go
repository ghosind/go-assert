package assert

import (
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestFailedHandler(t *testing.T) {
	mockT := new(testing.T)
	assert := New(t)

	failed(mockT, nil, false)
	assert.DeepEqual(mockT.Failed(), false)

	failed(mockT, newAssertionError("Test error"), false)
	assert.DeepEqual(mockT.Failed(), true)

	isTerminated := internal.CheckTermination(func() {
		failed(mockT, newAssertionError("Test error"), true)
	})
	assert.DeepEqual(isTerminated, true)
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

	assert.DeepEqual(isNil(1), false)  // int
	assert.DeepEqual(isNil(""), false) // string
	assert.DeepEqual(isNil(nil), true)
	var testAssert *Assertion
	assert.DeepEqual(isNil(testAssert), true)
	assert.DeepEqual(isNil(assert), false)
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
	assert.DeepEqual(isTrue(nil), false)

	// reflect.Slice
	assert.DeepEqual(isTrue([]int{0}), true)
	assert.DeepEqual(isTrue([]int{}), false)

	// other kinds
	assert.DeepEqual(isTrue(1), true)
	assert.DeepEqual(isTrue(0), false)
	assert.DeepEqual(isTrue(1.0), true)
	assert.DeepEqual(isTrue(0.0), false)
	assert.DeepEqual(isTrue("Hello"), true)
	assert.DeepEqual(isTrue(""), false)
	assert.DeepEqual(isTrue(func() {}), true)
}
