package assert

import (
	"regexp"
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestMatchAndNotMatch(t *testing.T) {
	mockT := new(testing.T)
	a := New(mockT)

	testMatchAndNotMatch(t, a, "Hello", `.+`, true)
	testMatchAndNotMatch(t, a, "", `.+`, false)
	testMatchAndNotMatch(t, a, "Hello", `^H`, true)
	testMatchAndNotMatch(t, a, "hello", `^H`, false)
}

func testMatchAndNotMatch(t *testing.T, a *Assertion, val string, pattern string, isMatch bool) {
	regPattern := regexp.MustCompile(pattern)

	testMatch(t, a, val, regPattern, isMatch)

	testNotMatch(t, a, val, regPattern, isMatch)

	testMatchNow(t, a, val, regPattern, isMatch)

	testNotMatchNow(t, a, val, regPattern, isMatch)

	testMatchString(t, a, val, pattern, isMatch)

	testNotMatchString(t, a, val, pattern, isMatch)

	testMatchStringNow(t, a, val, pattern, isMatch)

	testNotMatchStringNow(t, a, val, pattern, isMatch)
}

func testMatch(t *testing.T, a *Assertion, val string, pattern *regexp.Regexp, isMatch bool) {
	err := a.Match(val, pattern)
	if isMatch && err != nil {
		t.Errorf("Match(%s) = %v, want = nil", val, err)
	} else if !isMatch && err == nil {
		t.Errorf("Match(%s) = nil, want = error", val)
	}

	err = Match(a.T, val, pattern)
	if isMatch && err != nil {
		t.Errorf("Match(%s) = %v, want = nil", val, err)
	} else if !isMatch && err == nil {
		t.Errorf("Match(%s) = nil, want = error", val)
	}
}

func testNotMatch(t *testing.T, a *Assertion, val string, pattern *regexp.Regexp, isMatch bool) {
	err := a.NotMatch(val, pattern)
	if isMatch && err == nil {
		t.Errorf("NotMatch(%s) = nil, want = error", val)
	} else if !isMatch && err != nil {
		t.Errorf("NotMatch(%s) = %v, want = nil", val, err)
	}

	err = NotMatch(a.T, val, pattern)
	if isMatch && err == nil {
		t.Errorf("NotMatch(%s) = nil, want = error", val)
	} else if !isMatch && err != nil {
		t.Errorf("NotMatch(%s) = %v, want = nil", val, err)
	}
}

func testMatchNow(t *testing.T, a *Assertion, val string, pattern *regexp.Regexp, isMatch bool) {
	isTerminated := internal.CheckTermination(func() {
		a.MatchNow(val, pattern)
	})
	if isMatch && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isMatch && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		MatchNow(a.T, val, pattern)
	})
	if isMatch && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isMatch && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotMatchNow(t *testing.T, a *Assertion, val string, pattern *regexp.Regexp, isMatch bool) {
	isTerminated := internal.CheckTermination(func() {
		a.NotMatchNow(val, pattern)
	})
	if !isMatch && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isMatch && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotMatchNow(a.T, val, pattern)
	})
	if !isMatch && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isMatch && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testMatchString(t *testing.T, a *Assertion, val, pattern string, isMatch bool) {
	err := a.MatchString(val, pattern)
	if isMatch && err != nil {
		t.Errorf("MatchString(%s) = %v, want = nil", val, err)
	} else if !isMatch && err == nil {
		t.Errorf("MatchString(%s) = nil, want = error", val)
	}

	err = MatchString(a.T, val, pattern)
	if isMatch && err != nil {
		t.Errorf("MatchString(%s) = %v, want = nil", val, err)
	} else if !isMatch && err == nil {
		t.Errorf("MatchString(%s) = nil, want = error", val)
	}
}

func testNotMatchString(t *testing.T, a *Assertion, val, pattern string, isMatch bool) {
	err := a.NotMatchString(val, pattern)
	if isMatch && err == nil {
		t.Errorf("NotMatchString(%s) = nil, want = error", val)
	} else if !isMatch && err != nil {
		t.Errorf("NotMatchString(%s) = %v, want = nil", val, err)
	}

	err = NotMatchString(a.T, val, pattern)
	if isMatch && err == nil {
		t.Errorf("NotMatchString(%s) = nil, want = error", val)
	} else if !isMatch && err != nil {
		t.Errorf("NotMatchString(%s) = %v, want = nil", val, err)
	}
}

func testMatchStringNow(t *testing.T, a *Assertion, val, pattern string, isMatch bool) {
	isTerminated := internal.CheckTermination(func() {
		a.MatchStringNow(val, pattern)
	})
	if isMatch && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isMatch && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		MatchStringNow(a.T, val, pattern)
	})
	if isMatch && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isMatch && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotMatchStringNow(t *testing.T, a *Assertion, val, pattern string, isMatch bool) {
	isTerminated := internal.CheckTermination(func() {
		a.NotMatchStringNow(val, pattern)
	})
	if !isMatch && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isMatch && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotMatchStringNow(a.T, val, pattern)
	})
	if !isMatch && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isMatch && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}
