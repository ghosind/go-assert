package assert

import (
	"testing"

	"github.com/ghosind/go-assert/internal"
)

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
	testDeepEqual(t, assertion, v1, v2, isEqual)

	testNotDeepEqual(t, assertion, v1, v2, isEqual)

	testDeepEqualNow(t, assertion, v1, v2, isEqual)

	testNotDeepEqualNow(t, assertion, v1, v2, isEqual)
}

func testDeepEqual(t *testing.T, assertion *Assertion, v1, v2 any, isEqual bool) {
	err := assertion.DeepEqual(v1, v2)
	if isEqual && err != nil {
		t.Errorf("Equal(%v, %v) = %v, want = nil", v1, v2, err)
	} else if !isEqual && err == nil {
		t.Errorf("Equal(%v, %v) = nil, want = error", v1, v2)
	}

	err = DeepEqual(assertion.T, v1, v2)
	if isEqual && err != nil {
		t.Errorf("Equal(%v, %v) = %v, want = nil", v1, v2, err)
	} else if !isEqual && err == nil {
		t.Errorf("Equal(%v, %v) = nil, want = error", v1, v2)
	}
}

func testNotDeepEqual(t *testing.T, assertion *Assertion, v1, v2 any, isEqual bool) {
	err := assertion.NotDeepEqual(v1, v2)
	if isEqual && err == nil {
		t.Errorf("NotEqual(%v, %v) = nil, want = error", v1, v2)
	} else if !isEqual && err != nil {
		t.Errorf("NotEqual(%v, %v) = %v, want = nil", v1, v2, err)
	}

	err = NotDeepEqual(assertion.T, v1, v2)
	if isEqual && err == nil {
		t.Errorf("NotEqual(%v, %v) = nil, want = error", v1, v2)
	} else if !isEqual && err != nil {
		t.Errorf("NotEqual(%v, %v) = %v, want = nil", v1, v2, err)
	}
}

func testDeepEqualNow(t *testing.T, assertion *Assertion, v1, v2 any, isEqual bool) {
	isTerminated := internal.CheckTermination(func() {
		assertion.DeepEqualNow(v1, v2)
	})
	if isEqual && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isEqual && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		DeepEqualNow(assertion.T, v1, v2)
	})
	if isEqual && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isEqual && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotDeepEqualNow(t *testing.T, assertion *Assertion, v1, v2 any, isEqual bool) {
	isTerminated := internal.CheckTermination(func() {
		assertion.NotDeepEqualNow(v1, v2)
	})
	if !isEqual && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isEqual && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotDeepEqualNow(assertion.T, v1, v2)
	})
	if !isEqual && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isEqual && !isTerminated {
		t.Error("execution do not stopped, want stop")
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
	testNil(t, assertion, v, isNil)

	testNotNil(t, assertion, v, isNil)

	testNilNow(t, assertion, v, isNil)

	testNotNilNow(t, assertion, v, isNil)
}

func testNil(t *testing.T, assertion *Assertion, v any, isNil bool) {
	err := assertion.Nil(v)
	if isNil && err != nil {
		t.Errorf("Nil(%v) = %v, want nil", v, err)
	} else if !isNil && err == nil {
		t.Errorf("Nil(%v) = nil, want error", v)
	}

	err = Nil(assertion.T, v)
	if isNil && err != nil {
		t.Errorf("Nil(%v) = %v, want nil", v, err)
	} else if !isNil && err == nil {
		t.Errorf("Nil(%v) = nil, want error", v)
	}
}

func testNotNil(t *testing.T, assertion *Assertion, v any, isNil bool) {
	err := assertion.NotNil(v)
	if isNil && err == nil {
		t.Errorf("Nil(%v) = nil, want error", v)
	} else if !isNil && err != nil {
		t.Errorf("Nil(%v) = %v, want nil", v, err)
	}

	err = NotNil(assertion.T, v)
	if isNil && err == nil {
		t.Errorf("Nil(%v) = nil, want error", v)
	} else if !isNil && err != nil {
		t.Errorf("Nil(%v) = %v, want nil", v, err)
	}
}

func testNilNow(t *testing.T, assertion *Assertion, v any, isNil bool) {
	isTerminated := internal.CheckTermination(func() {
		assertion.NilNow(v)
	})
	if isNil && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isNil && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NilNow(assertion.T, v)
	})
	if isNil && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isNil && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotNilNow(t *testing.T, assertion *Assertion, v any, isNil bool) {
	isTerminated := internal.CheckTermination(func() {
		assertion.NotNilNow(v)
	})
	if !isNil && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isNil && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotNilNow(assertion.T, v)
	})
	if !isNil && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isNil && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}
