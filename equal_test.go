package assert

import "testing"

type testStruct struct {
	v int
}

func TestDeepEqualAndNotDeepEqual(t *testing.T) {
	mockT := new(testing.T)
	assertion := New(mockT)

	testDeepEqualAndNotDeepEqual(t, assertion, 1, 1, true)
	testDeepEqualAndNotDeepEqual(t, assertion, 1, 2, false)
	testDeepEqualAndNotDeepEqual(t, assertion, 1, 1.0, false)
	testDeepEqualAndNotDeepEqual(t, assertion, 1, "1", false)
	testDeepEqualAndNotDeepEqual(t, assertion, 1, '1', false)
	testDeepEqualAndNotDeepEqual(t, assertion, 1, []int{1}, false)
	testDeepEqualAndNotDeepEqual(t, assertion, []int{1}, []int{1}, true)

	obj1 := testStruct{v: 1}
	obj2 := testStruct{v: 1}

	testDeepEqualAndNotDeepEqual(t, assertion, obj1, obj2, true)
	testDeepEqualAndNotDeepEqual(t, assertion, obj1, &obj2, false)
	testDeepEqualAndNotDeepEqual(t, assertion, &obj1, &obj2, true)

	obj2.v = 2
	testDeepEqualAndNotDeepEqual(t, assertion, obj1, obj2, false)
	testDeepEqualAndNotDeepEqual(t, assertion, &obj1, &obj2, false)
}

func testDeepEqualAndNotDeepEqual(t *testing.T, assertion *Assertion, v1, v2 any, isEqual bool) {
	err := assertion.DeepEqual(v1, v2)
	if isEqual && err != nil {
		t.Errorf("Equal(%v, %v) = %v, want = nil", v1, v2, err)
	} else if !isEqual && err == nil {
		t.Errorf("Equal(%v, %v) = nil, want = error", v1, v2)
	}

	err = assertion.NotDeepEqual(v1, v2)
	if isEqual && err == nil {
		t.Errorf("NotEqual(%v, %v) = nil, want = error", v1, v2)
	} else if !isEqual && err != nil {
		t.Errorf("NotEqual(%v, %v) = %v, want = nil", v1, v2, err)
	}
}

func TestNilAndNotNil(t *testing.T) {
	mockT := new(testing.T)
	assert := New(mockT)

	testNilAndNotNil(t, assert, 1, false)
	testNilAndNotNil(t, assert, "", false)
	testNilAndNotNil(t, assert, nil, true)
	var testAssert *Assertion
	testNilAndNotNil(t, assert, testAssert, true)
	testNilAndNotNil(t, assert, assert, false)
}

func testNilAndNotNil(t *testing.T, assertion *Assertion, v any, isNil bool) {
	err := assertion.Nil(v)
	if isNil && err != nil {
		t.Errorf("Nil(%v) = %v, want nil", v, err)
	} else if !isNil && err == nil {
		t.Errorf("Nil(%v) = nil, want error", v)
	}

	err = assertion.NotNil(v)
	if isNil && err == nil {
		t.Errorf("Nil(%v) = nil, want error", v)
	} else if !isNil && err != nil {
		t.Errorf("Nil(%v) = %v, want nil", v, err)
	}
}
