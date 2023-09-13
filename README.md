# Assertion for Golang

![test](https://github.com/ghosind/go-assert/workflows/test/badge.svg)
[![codecov](https://codecov.io/gh/ghosind/go-assert/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-assert)
![Version Badge](https://img.shields.io/github/v/release/ghosind/go-assert)
![License Badge](https://img.shields.io/github/license/ghosind/go-assert)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/go-assert.svg)](https://pkg.go.dev/github.com/ghosind/go-assert)

A collection of Golang assertion functions for verifying invariants.

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Available Assertions](#available-assertions)
  - [Equality](#equality)
  - [Value](#value)
  - [String](#string)
  - [Slice or Array](#slice-or-array)
  - [Error Handling](#error-handling)
- [License](#license)

## Installation

To install this library, just use `go get` command like the following line:

```bash
go get -u github.com/ghosind/go-assert
```

## Getting Started

This library provided assertion functions to verify the equality of values, or assert for nil:

```go
func TestExample(t *testing.T) {
  // var actual
  // var expect

  // assert equality
  assert.Equal(t, actual, expect)

  // assert inequality
  assert.NotEqual(t, actual, expect)

  // you can also use DeepEqual to assert the equality that also checks the type between the values
  assert.DeepEqual(t, actual, expect)

  // var object

  // assert for nil
  assert.Nil(t, object)

  // assert for not nil
  assert.NotNil(t, object)
}
```

You can use `True` method to check whether a value is truthy or falsy (is the zero value of the type or not).

```go
func TestExample(t *testing.T) {
  assert.True(t, 1) // success
  assert.True(t, 0) // fail
  assert.True(t, "test") // success
  assert.True(t, "") // fail
}
```

If you want to test the value of a string, you can use `Match` method to test it with a regular expression pattern.

```go
func TestExample(t *testing.T) {
  pattern := regexp.MustCompile(`^https?:\/\/`)
  assert.Match(t, "https://example.com", pattern) // success
  assert.Match(t, "example.com", pattern) // fail

  // you can also use `MatchString` to test it without compiling the regexp pattern
  assert.MatchString(t, "https://example.com", `^https?:\/\/`) // success
}
```

Since v0.2.0, we also provided some assertions for array/slice, for example, you can use `ContainsElement` to check whether an array or a slice contains a specified element.

```go
func TestExample(t *testing.T) {
  arr := []int{1, 2, 3}
  assert.ContainsElement(arr, 1) // success
  assert.ContainsElement(arr, 4) // fail
}
```

It also provided assertion functions to verify a function will panic or not:

```go
func TestPanic(t *testing.T) {
  // assert panic
  assert.Panic(t, func () {
    // do something

    panic()
  })

  // assert no panic
  assert.NotPanic(t, func () {
    // do something

    // panic()
  })
}
```

For every assertion functions, it also provided `XXXNow` functions to stop the execution if the test is failed.

```go
func TestExample(t *testing.T) {
  // var actual
  // var expect

  // The following line will set the test result to fail and stop the execution
  assert.EqualNow(t, actual, expect)

  // The following lines will never execute if they are not equal.
  // ...
}
```

Every assertion will not terminate the testing workflow. However, they'll return an error if the verification failed, and you can check the return value to get the verification result.

```go
func TestExample(t *testing.T) {
  if err := assert.Equal(t, actual, expect); err != nil {
    // terminate test
    t.Fail()
  }
}
```

If you need to assert many times, you can also create an `Assertion` instance:

```go
func TestExample(t *testing.T) {
  assertion := assert.New(t)

  // test equality
  assertion.Equal(actual, expect)

  // Test inequality
  assertion.NotEqual(actual, expect)
}
```

## Available Assertions

### Equality

- [`DeepEqual`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.DeepEqual) and [`NotDeepEqual`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotDeepEqual): assert the deep equality or inequality.

  > Since v0.1.0

- [`Equal`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.Equal) and [`NotEqual`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotEqual): assert the equality or inequality.

  > Since v0.1.5

### Value

- [`Nil`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.Nil) and [`NotNil`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotNil): assert the value is nil or not.

  > Since v0.1.1

- [`True`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.True) and [`NotTrue`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotTrue): assert the truthy of the value.

  > Since v0.1.4

### String

- [`ContainsString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.ContainsString) and [`NotContainsString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotContainsString): assert whether the string contains the substring or not.

  > Since v0.1.7

- [`HasPrefixString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.HasPrefixString) and [`NotHasPrefixString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotHasPrefixString): assert whether the string have the prefix string or not.

  > Since v0.1.7

- [`HasSuffixString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.HasSuffixString) and [`NotHasSuffixString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotHasSuffixString): assert whether the string have the suffix string or not.

  > Since v0.1.7

- [`Match`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.Match) and [`NotMatch`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotMatch): assert whether the string matches the regular expression pattern or not.

  > Since v0.1.5

- [`MatchString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.MatchString) and [`NotMatchString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotMatchString): compile the regular expression pattern and assert whether the string matches the pattern or not.

  > Since v0.1.5

### Slice or Array

- [`ContainsElement`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.ContainsElement) and [`NotContainsElement`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotContainsElement): assert whether the array or slice contains the specified element or not.

  > Since v0.2.0

### Error Handling

- [`Panic`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.Panic) and [`NotPanic`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotPanic): assert the function will panic or not.

  > Since v0.1.0

## License

This project was published under the MIT license, you can see [LICENSE](./LICENSE) file to get more information.
