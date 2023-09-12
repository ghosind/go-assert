package assert

import (
	"regexp"
	"testing"
)

func TestStringContainsAndNotContains(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testStringContainsAndNotContains(a, mockA, "", "", true)
	testStringContainsAndNotContains(a, mockA, "Hello world", "", true)
	testStringContainsAndNotContains(a, mockA, "Hello world", "Hello", true)
	testStringContainsAndNotContains(a, mockA, "Hello world", "hello", false)
	testStringContainsAndNotContains(a, mockA, "", "Hello", false)
	testStringContainsAndNotContains(a, mockA, "", "world", false)
	testStringContainsAndNotContains(a, mockA, "Hello world", "world", true)
	testStringContainsAndNotContains(a, mockA, "Hello world", "o w", true)
}

func testStringContainsAndNotContains(
	a, mockA *Assertion,
	str, substr string,
	isContains bool,
) {
	// ContainsString
	testAssertionFunction(a, "ContainsString", func() error {
		return ContainsString(mockA.T, str, substr)
	}, isContains)
	testAssertionFunction(a, "Assertion.ContainsString", func() error {
		return mockA.ContainsString(str, substr)
	}, isContains)

	// NotContainsString
	testAssertionFunction(a, "NotContainsString", func() error {
		return NotContainsString(mockA.T, str, substr)
	}, !isContains)
	testAssertionFunction(a, "Assertion.NotContainsString", func() error {
		return mockA.NotContainsString(str, substr)
	}, !isContains)

	// ContainsStringNow
	testAssertionNowFunction(a, "ContainsStringNow", func() {
		ContainsStringNow(mockA.T, str, substr)
	}, !isContains)
	testAssertionNowFunction(a, "Assertion.ContainsStringNow", func() {
		mockA.ContainsStringNow(str, substr)
	}, !isContains)

	// NotContainsStringNow
	testAssertionNowFunction(a, "NotContainsStringNow", func() {
		NotContainsStringNow(mockA.T, str, substr)
	}, isContains)
	testAssertionNowFunction(a, "Assertion.NotContainsStringNow", func() {
		mockA.NotContainsStringNow(str, substr)
	}, isContains)
}

func TestStringHasPrefixAndNotHasPrefix(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testStringHasPrefixAndNotHasPrefix(a, mockA, "", "", true)
	testStringHasPrefixAndNotHasPrefix(a, mockA, "Hello world", "", true)
	testStringHasPrefixAndNotHasPrefix(a, mockA, "Hello world", "Hello", true)
	testStringHasPrefixAndNotHasPrefix(a, mockA, "Hello world", "hello", false)
	testStringHasPrefixAndNotHasPrefix(a, mockA, "", "Hello", false)
	testStringHasPrefixAndNotHasPrefix(a, mockA, "", "world", false)
	testStringHasPrefixAndNotHasPrefix(a, mockA, "Hello world", "world", false)
}

func testStringHasPrefixAndNotHasPrefix(
	a, mockA *Assertion,
	str, prefix string,
	isHasPrefix bool,
) {
	// HasPrefixString
	testAssertionFunction(a, "HasPrefixString", func() error {
		return HasPrefixString(mockA.T, str, prefix)
	}, isHasPrefix)
	testAssertionFunction(a, "Assertion.HasPrefixString", func() error {
		return mockA.HasPrefixString(str, prefix)
	}, isHasPrefix)

	// NotHasPrefixString
	testAssertionFunction(a, "NotHasPrefixString", func() error {
		return NotHasPrefixString(mockA.T, str, prefix)
	}, !isHasPrefix)
	testAssertionFunction(a, "Assertion.NotHasPrefixString", func() error {
		return mockA.NotHasPrefixString(str, prefix)
	}, !isHasPrefix)

	// HasPrefixStringNow
	testAssertionNowFunction(a, "HasPrefixStringNow", func() {
		HasPrefixStringNow(mockA.T, str, prefix)
	}, !isHasPrefix)
	testAssertionNowFunction(a, "Assertion.HasPrefixStringNow", func() {
		mockA.HasPrefixStringNow(str, prefix)
	}, !isHasPrefix)

	// NotHasPrefixStringNow
	testAssertionNowFunction(a, "NotHasPrefixStringNow", func() {
		NotHasPrefixStringNow(mockA.T, str, prefix)
	}, isHasPrefix)
	testAssertionNowFunction(a, "Assertion.NotHasPrefixStringNow", func() {
		mockA.NotHasPrefixStringNow(str, prefix)
	}, isHasPrefix)
}

func TestStringHasSuffixAndNotHasSuffix(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testStringHasSuffixAndNotHasSuffix(a, mockA, "", "", true)
	testStringHasSuffixAndNotHasSuffix(a, mockA, "Hello world", "", true)
	testStringHasSuffixAndNotHasSuffix(a, mockA, "Hello world", "Hello", false)
	testStringHasSuffixAndNotHasSuffix(a, mockA, "Hello world", "hello", false)
	testStringHasSuffixAndNotHasSuffix(a, mockA, "", "Hello", false)
	testStringHasSuffixAndNotHasSuffix(a, mockA, "", "world", false)
	testStringHasSuffixAndNotHasSuffix(a, mockA, "Hello world", "world", true)
}

func testStringHasSuffixAndNotHasSuffix(
	a, mockA *Assertion,
	str, suffix string,
	isHasSuffix bool,
) {
	// HasSuffixString
	testAssertionFunction(a, "HasSuffixString", func() error {
		return HasSuffixString(mockA.T, str, suffix)
	}, isHasSuffix)
	testAssertionFunction(a, "Assertion.HasSuffixString", func() error {
		return mockA.HasSuffixString(str, suffix)
	}, isHasSuffix)

	// NotHasSuffixString
	testAssertionFunction(a, "NotHasSuffixString", func() error {
		return NotHasSuffixString(mockA.T, str, suffix)
	}, !isHasSuffix)
	testAssertionFunction(a, "Assertion.NotHasSuffixString", func() error {
		return mockA.NotHasSuffixString(str, suffix)
	}, !isHasSuffix)

	// HasSuffixStringNow
	testAssertionNowFunction(a, "HasSuffixStringNow", func() {
		HasSuffixStringNow(mockA.T, str, suffix)
	}, !isHasSuffix)
	testAssertionNowFunction(a, "Assertion.HasSuffixStringNow", func() {
		mockA.HasSuffixStringNow(str, suffix)
	}, !isHasSuffix)

	// NotHasSuffixStringNow
	testAssertionNowFunction(a, "NotHasSuffixStringNow", func() {
		NotHasSuffixStringNow(mockA.T, str, suffix)
	}, isHasSuffix)
	testAssertionNowFunction(a, "Assertion.NotHasSuffixStringNow", func() {
		mockA.NotHasSuffixStringNow(str, suffix)
	}, isHasSuffix)
}

func TestMatchAndNotMatch(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testMatchAndNotMatch(a, mockA, "Hello", `.+`, true)
	testMatchAndNotMatch(a, mockA, "", `.+`, false)
	testMatchAndNotMatch(a, mockA, "Hello", `^H`, true)
	testMatchAndNotMatch(a, mockA, "hello", `^H`, false)
}

func testMatchAndNotMatch(a, mockA *Assertion, val string, pattern string, isMatch bool) {
	regPattern := regexp.MustCompile(pattern)

	// MatchString
	testAssertionFunction(a, "MatchString", func() error {
		return MatchString(mockA.T, val, pattern)
	}, isMatch)
	testAssertionFunction(a, "Assertion.MatchString", func() error {
		return mockA.MatchString(val, pattern)
	}, isMatch)

	// NotMatchString
	testAssertionFunction(a, "NotMatchString", func() error {
		return NotMatchString(mockA.T, val, pattern)
	}, !isMatch)
	testAssertionFunction(a, "Assertion.NotMatchString", func() error {
		return mockA.NotMatchString(val, pattern)
	}, !isMatch)

	// MatchStringNow
	testAssertionNowFunction(a, "MatchStringNow", func() {
		MatchStringNow(mockA.T, val, pattern)
	}, !isMatch)
	testAssertionNowFunction(a, "Assertion.MatchStringNow", func() {
		mockA.MatchStringNow(val, pattern)
	}, !isMatch)

	// NotMatchStringNow
	testAssertionNowFunction(a, "NotMatchStringNow", func() {
		NotMatchStringNow(mockA.T, val, pattern)
	}, isMatch)
	testAssertionNowFunction(a, "Assertion.NotMatchStringNow", func() {
		mockA.NotMatchStringNow(val, pattern)
	}, isMatch)

	// Match
	testAssertionFunction(a, "Match", func() error {
		return Match(mockA.T, val, regPattern)
	}, isMatch)
	testAssertionFunction(a, "Assertion.Match", func() error {
		return mockA.Match(val, regPattern)
	}, isMatch)

	// NotMatch
	testAssertionFunction(a, "NotMatch", func() error {
		return NotMatch(mockA.T, val, regPattern)
	}, !isMatch)
	testAssertionFunction(a, "Assertion.NotMatch", func() error {
		return mockA.NotMatch(val, regPattern)
	}, !isMatch)

	// MatchNow
	testAssertionNowFunction(a, "MatchNow", func() {
		MatchNow(mockA.T, val, regPattern)
	}, !isMatch)
	testAssertionNowFunction(a, "Assertion.MatchNow", func() {
		mockA.MatchNow(val, regPattern)
	}, !isMatch)

	// NotMatchNow
	testAssertionNowFunction(a, "NotMatchNow", func() {
		NotMatchNow(mockA.T, val, regPattern)
	}, isMatch)
	testAssertionNowFunction(a, "Assertion.NotMatchNow", func() {
		mockA.NotMatchNow(val, regPattern)
	}, isMatch)
}
