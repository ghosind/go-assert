package assert

import (
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestElementContainsAndNotContains(t *testing.T) {
	mockT := new(testing.T)
	assert := New(mockT)

	testElementContainsAndNotContains(t, assert, []int{1}, 1, true)
	testElementContainsAndNotContains(t, assert, []int{1, 2, 3}, 3, true)
	testElementContainsAndNotContains(t, assert, []int{1, 2, 3}, 4, false)
	testElementContainsAndNotContains(t, assert, []int{}, 1, false)
	testElementContainsAndNotContains(t, assert, [1]int{1}, 1, true)
	testElementContainsAndNotContains(t, assert, [3]int{1, 2, 3}, 3, true)
	testElementContainsAndNotContains(t, assert, [3]int{1, 2, 3}, 4, false)
	testElementContainsAndNotContains(t, assert, [0]int{}, 1, false)
}

func testElementContainsAndNotContains(
	t *testing.T,
	assertion *Assertion,
	source, expect any,
	isContains bool,
) {
	testContainsElement(t, assertion, source, expect, isContains)

	testNotContainsElement(t, assertion, source, expect, isContains)

	testContainsElementNow(t, assertion, source, expect, isContains)

	testNotContainsElementNow(t, assertion, source, expect, isContains)
}

func testContainsElement(
	t *testing.T,
	assertion *Assertion,
	source, expect any,
	isContains bool,
) {
	err := assertion.ContainsElement(source, expect)
	if isContains && err != nil {
		t.Errorf("ContainsElement(\"%v\", \"%v\") = %v, want nil", source, expect, err)
	} else if !isContains && err == nil {
		t.Errorf("ContainsElement(\"%v\", \"%v\") = nil, want error", source, expect)
	}

	err = ContainsElement(assertion.T, source, expect)
	if isContains && err != nil {
		t.Errorf("ContainsElement(\"%v\", \"%v\") = %v, want nil", source, expect, err)
	} else if !isContains && err == nil {
		t.Errorf("ContainsElement(\"%v\", \"%v\") = nil, want error", source, expect)
	}
}

func testNotContainsElement(
	t *testing.T,
	assertion *Assertion,
	source, expect any,
	isContains bool,
) {
	err := assertion.NotContainsElement(source, expect)
	if isContains && err == nil {
		t.Errorf("NotContainsElement(\"%v\", \"%v\") = nil, want error", source, expect)
	} else if !isContains && err != nil {
		t.Errorf("NotContainsElement(\"%v\", \"%v\") = %v, want nil", source, expect, err)
	}

	err = NotContainsElement(assertion.T, source, expect)
	if isContains && err == nil {
		t.Errorf("NotContainsElement(\"%v\", \"%v\") = nil, want error", source, expect)
	} else if !isContains && err != nil {
		t.Errorf("NotContainsElement(\"%v\", \"%v\") = %v, want nil", source, expect, err)
	}
}

func testContainsElementNow(
	t *testing.T,
	assertion *Assertion,
	source, expect any,
	isContains bool,
) {
	isTerminated := internal.CheckTermination(func() {
		assertion.ContainsElementNow(source, expect)
	})
	if isContains && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isContains && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		ContainsElementNow(assertion.T, source, expect)
	})
	if isContains && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isContains && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotContainsElementNow(
	t *testing.T,
	assertion *Assertion,
	source, expect any,
	isContains bool,
) {
	isTerminated := internal.CheckTermination(func() {
		assertion.NotContainsElementNow(source, expect)
	})
	if !isContains && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isContains && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotContainsElementNow(assertion.T, source, expect)
	})
	if !isContains && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isContains && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}
