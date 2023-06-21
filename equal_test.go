package assert

import "testing"

type testStruct struct {
	v int
}

func TestEqualAndNotEqual(t *testing.T) {
	mockT := new(testing.T)
	assertion := New(mockT)

	testEqualAndNotEqual(t, assertion, 1, 1, true)
	testEqualAndNotEqual(t, assertion, 1, 2, false)
	testEqualAndNotEqual(t, assertion, 1, 1.0, false)
	testEqualAndNotEqual(t, assertion, 1, "1", false)
	testEqualAndNotEqual(t, assertion, 1, '1', false)
	testEqualAndNotEqual(t, assertion, 1, []int{1}, false)
	testEqualAndNotEqual(t, assertion, []int{1}, []int{1}, true)

	obj1 := testStruct{v: 1}
	obj2 := testStruct{v: 1}

	testEqualAndNotEqual(t, assertion, obj1, obj2, true)
	testEqualAndNotEqual(t, assertion, obj1, &obj2, false)
	testEqualAndNotEqual(t, assertion, &obj1, &obj2, true)

	obj2.v = 2
	testEqualAndNotEqual(t, assertion, obj1, obj2, false)
	testEqualAndNotEqual(t, assertion, &obj1, &obj2, false)
}

func testEqualAndNotEqual(t *testing.T, assertion *Assertion, v1, v2 any, isEqual bool) {
	err := assertion.Equal(v1, v2)
	if isEqual && err != nil {
		t.Errorf("Equal(%v, %v) = %v, want = nil", v1, v2, err)
	} else if !isEqual && err == nil {
		t.Errorf("Equal(%v, %v) = nil, want = error", v1, v2)
	}

	err = assertion.NotEqual(v1, v2)
	if isEqual && err == nil {
		t.Errorf("NotEqual(%v, %v) = nil, want = error", v1, v2)
	} else if !isEqual && err != nil {
		t.Errorf("NotEqual(%v, %v) = %v, want = nil", v1, v2, err)
	}
}
