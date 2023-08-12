# Assertion for Golang

![test](https://github.com/ghosind/go-assert/workflows/test/badge.svg)
[![codecov](https://codecov.io/gh/ghosind/go-assert/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-assert)
![Version Badge](https://img.shields.io/github/v/release/ghosind/go-assert)
![License Badge](https://img.shields.io/github/license/ghosind/go-assert)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/go-assert.svg)](https://pkg.go.dev/github.com/ghosind/go-assert)

A collection of Golang assertion functions for verifying invariants.

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

  // var object

  // assert for nil
  assert.Nil(t, object)

  // assert for not nil
  assert.NotNil(t, object)
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

  // The following lines will never execute if they are not deep equal.
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

- [`DeepEqual`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.DeepEqual) and [`NotDeepEqual`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotDeepEqual): assert the deep equality or inequality.
- [`Equal`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.Equal) and [`NotEqual`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotEqual): assert the equality or inequality.
- [`Match`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.Match) and [`NotMatch`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotMatch): assert whether the string matches the regular expression pattern or not.
- [`MatchString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.MatchString) and [`NotMatchString`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotMatchString): compile the regular expression pattern and assert whether the string matches the pattern or not.
- [`Nil`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.Nil) and [`NotNil`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotNil): assert the value is nil or not.
- [`Panic`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.Panic) and [`NotPanic`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotPanic): assert the function will panic or not.
- [`True`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.True) and [`NotTrue`](https://pkg.go.dev/github.com/ghosind/go-assert#Assertion.NotTrue): assert the truthy of the value.

## License

This project was published under the MIT license, you can see [LICENSE](./LICENSE) file to get more information.
