package assert

import (
	"regexp"
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestStringContainsAndNotContains(t *testing.T) {
	mockT := new(testing.T)
	assert := New(mockT)

	testStringContainsAndNotContains(t, assert, "", "", true)
	testStringContainsAndNotContains(t, assert, "Hello world", "", true)
	testStringContainsAndNotContains(t, assert, "Hello world", "Hello", true)
	testStringContainsAndNotContains(t, assert, "Hello world", "hello", false)
	testStringContainsAndNotContains(t, assert, "", "Hello", false)
	testStringContainsAndNotContains(t, assert, "", "world", false)
	testStringContainsAndNotContains(t, assert, "Hello world", "world", true)
	testStringContainsAndNotContains(t, assert, "Hello world", "o w", true)
}

func testStringContainsAndNotContains(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isContains bool,
) {
	testContainsString(t, assertion, str, substr, isContains)

	testNotContainsString(t, assertion, str, substr, isContains)

	testContainsStringNow(t, assertion, str, substr, isContains)

	testNotContainsStringNow(t, assertion, str, substr, isContains)
}

func testContainsString(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isContains bool,
) {
	err := assertion.ContainsString(str, substr)
	if isContains && err != nil {
		t.Errorf("ContainsString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	} else if !isContains && err == nil {
		t.Errorf("ContainsString(\"%s\", \"%s\") = nil, want error", str, substr)
	}

	err = ContainsString(assertion.T, str, substr)
	if isContains && err != nil {
		t.Errorf("ContainsString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	} else if !isContains && err == nil {
		t.Errorf("ContainsString(\"%s\", \"%s\") = nil, want error", str, substr)
	}
}

func testNotContainsString(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isContains bool,
) {
	err := assertion.NotContainsString(str, substr)
	if isContains && err == nil {
		t.Errorf("NotContainsString(\"%s\", \"%s\") = nil, want error", str, substr)
	} else if !isContains && err != nil {
		t.Errorf("NotContainsString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	}

	err = NotContainsString(assertion.T, str, substr)
	if isContains && err == nil {
		t.Errorf("NotContainsString(\"%s\", \"%s\") = nil, want error", str, substr)
	} else if !isContains && err != nil {
		t.Errorf("NotContainsString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	}
}

func testContainsStringNow(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isContains bool,
) {
	isTerminated := internal.CheckTermination(func() {
		assertion.ContainsStringNow(str, substr)
	})
	if isContains && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isContains && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		ContainsStringNow(assertion.T, str, substr)
	})
	if isContains && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isContains && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotContainsStringNow(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isContains bool,
) {
	isTerminated := internal.CheckTermination(func() {
		assertion.NotContainsStringNow(str, substr)
	})
	if !isContains && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isContains && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotContainsStringNow(assertion.T, str, substr)
	})
	if !isContains && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isContains && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func TestStringHasPrefixAndNotHasPrefix(t *testing.T) {
	mockT := new(testing.T)
	assert := New(mockT)

	testStringHasPrefixAndNotHasPrefix(t, assert, "", "", true)
	testStringHasPrefixAndNotHasPrefix(t, assert, "Hello world", "", true)
	testStringHasPrefixAndNotHasPrefix(t, assert, "Hello world", "Hello", true)
	testStringHasPrefixAndNotHasPrefix(t, assert, "Hello world", "hello", false)
	testStringHasPrefixAndNotHasPrefix(t, assert, "", "Hello", false)
	testStringHasPrefixAndNotHasPrefix(t, assert, "", "world", false)
	testStringHasPrefixAndNotHasPrefix(t, assert, "Hello world", "world", false)
}

func testStringHasPrefixAndNotHasPrefix(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasPrefix bool,
) {
	testHasPrefixString(t, assertion, str, substr, isHasPrefix)

	testNotHasPrefixString(t, assertion, str, substr, isHasPrefix)

	testHasPrefixStringNow(t, assertion, str, substr, isHasPrefix)

	testNotHasPrefixStringNow(t, assertion, str, substr, isHasPrefix)
}

func testHasPrefixString(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasPrefix bool,
) {
	err := assertion.HasPrefixString(str, substr)
	if isHasPrefix && err != nil {
		t.Errorf("HasPrefixString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	} else if !isHasPrefix && err == nil {
		t.Errorf("HasPrefixString(\"%s\", \"%s\") = nil, want error", str, substr)
	}

	err = HasPrefixString(assertion.T, str, substr)
	if isHasPrefix && err != nil {
		t.Errorf("HasPrefixString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	} else if !isHasPrefix && err == nil {
		t.Errorf("HasPrefixString(\"%s\", \"%s\") = nil, want error", str, substr)
	}
}

func testNotHasPrefixString(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasPrefix bool,
) {
	err := assertion.NotHasPrefixString(str, substr)
	if isHasPrefix && err == nil {
		t.Errorf("NotHasPrefixString(\"%s\", \"%s\") = nil, want error", str, substr)
	} else if !isHasPrefix && err != nil {
		t.Errorf("NotHasPrefixString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	}

	err = NotHasPrefixString(assertion.T, str, substr)
	if isHasPrefix && err == nil {
		t.Errorf("NotHasPrefixString(\"%s\", \"%s\") = nil, want error", str, substr)
	} else if !isHasPrefix && err != nil {
		t.Errorf("NotHasPrefixString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	}
}

func testHasPrefixStringNow(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasPrefix bool,
) {
	isTerminated := internal.CheckTermination(func() {
		assertion.HasPrefixStringNow(str, substr)
	})
	if isHasPrefix && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isHasPrefix && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		HasPrefixStringNow(assertion.T, str, substr)
	})
	if isHasPrefix && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isHasPrefix && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotHasPrefixStringNow(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasPrefix bool,
) {
	isTerminated := internal.CheckTermination(func() {
		assertion.NotHasPrefixStringNow(str, substr)
	})
	if !isHasPrefix && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isHasPrefix && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotHasPrefixStringNow(assertion.T, str, substr)
	})
	if !isHasPrefix && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isHasPrefix && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func TestStringHasSuffixAndNotHasSuffix(t *testing.T) {
	mockT := new(testing.T)
	assert := New(mockT)

	testStringHasSuffixAndNotHasSuffix(t, assert, "", "", true)
	testStringHasSuffixAndNotHasSuffix(t, assert, "Hello world", "", true)
	testStringHasSuffixAndNotHasSuffix(t, assert, "Hello world", "Hello", false)
	testStringHasSuffixAndNotHasSuffix(t, assert, "Hello world", "hello", false)
	testStringHasSuffixAndNotHasSuffix(t, assert, "", "Hello", false)
	testStringHasSuffixAndNotHasSuffix(t, assert, "", "world", false)
	testStringHasSuffixAndNotHasSuffix(t, assert, "Hello world", "world", true)
}

func testStringHasSuffixAndNotHasSuffix(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasSuffix bool,
) {
	testHasSuffixString(t, assertion, str, substr, isHasSuffix)

	testNotHasSuffixString(t, assertion, str, substr, isHasSuffix)

	testHasSuffixStringNow(t, assertion, str, substr, isHasSuffix)

	testNotHasSuffixStringNow(t, assertion, str, substr, isHasSuffix)
}

func testHasSuffixString(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasSuffix bool,
) {
	err := assertion.HasSuffixString(str, substr)
	if isHasSuffix && err != nil {
		t.Errorf("HasSuffixString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	} else if !isHasSuffix && err == nil {
		t.Errorf("HasSuffixString(\"%s\", \"%s\") = nil, want error", str, substr)
	}

	err = HasSuffixString(assertion.T, str, substr)
	if isHasSuffix && err != nil {
		t.Errorf("HasSuffixString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	} else if !isHasSuffix && err == nil {
		t.Errorf("HasSuffixString(\"%s\", \"%s\") = nil, want error", str, substr)
	}
}

func testNotHasSuffixString(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasSuffix bool,
) {
	err := assertion.NotHasSuffixString(str, substr)
	if isHasSuffix && err == nil {
		t.Errorf("NotHasSuffixString(\"%s\", \"%s\") = nil, want error", str, substr)
	} else if !isHasSuffix && err != nil {
		t.Errorf("NotHasSuffixString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	}

	err = NotHasSuffixString(assertion.T, str, substr)
	if isHasSuffix && err == nil {
		t.Errorf("NotHasSuffixString(\"%s\", \"%s\") = nil, want error", str, substr)
	} else if !isHasSuffix && err != nil {
		t.Errorf("NotHasSuffixString(\"%s\", \"%s\") = %v, want nil", str, substr, err)
	}
}

func testHasSuffixStringNow(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasSuffix bool,
) {
	isTerminated := internal.CheckTermination(func() {
		assertion.HasSuffixStringNow(str, substr)
	})
	if isHasSuffix && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isHasSuffix && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		HasSuffixStringNow(assertion.T, str, substr)
	})
	if isHasSuffix && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isHasSuffix && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotHasSuffixStringNow(
	t *testing.T,
	assertion *Assertion,
	str, substr string,
	isHasSuffix bool,
) {
	isTerminated := internal.CheckTermination(func() {
		assertion.NotHasSuffixStringNow(str, substr)
	})
	if !isHasSuffix && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isHasSuffix && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotHasSuffixStringNow(assertion.T, str, substr)
	})
	if !isHasSuffix && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isHasSuffix && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

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
