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

This library provided assertion functions to verify the equality of values:

```go
func TestExample(t *testing.T) {
  // assert equality
  assert.DeepEqual(t, actual, expect)

  // assert inequality
  assert.NotDeepEqual(t, actual, expect)
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

Every assertion will not terminate the testing workflow. However, they'll return an error if the verification failed, and you can check the return value to get the verification result.

```go
func TestExample(t *testing.T) {
  if err := assert.DeepEqual(t, actual, expect); err != nil {
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
  assertion.DeepEqual(actual, expect)

  // Test inequality
  assertion.NotDeepEqual(actual, expect)
}
```

## License

This project was published under the MIT license, you can see [LICENSE](./LICENSE) file to get more information.
