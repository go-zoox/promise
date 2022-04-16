# Promise - JavaScript Promise Like with Goroutines

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/promise)](https://pkg.go.dev/github.com/go-zoox/promise)
[![Build Status](https://github.com/go-zoox/promise/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/promise/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/promise)](https://goreportcard.com/report/github.com/go-zoox/promise)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/promise/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/promise?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/promise.svg)](https://github.com/go-zoox/promise/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/promise.svg?label=Release)](https://github.com/go-zoox/promise/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/promise
```

## Getting Started

```go
func main() {
	p := New(func(resolve Resolver, reject Rejecter) {
		time.Sleep(100 * time.Millisecond)
		resolve(1)
	})

	p.Then(func(v interface{}) interface{} {
		return v.(int) + 1
	}).Then(func(v interface{}) interface{} {
		return v.(int) * 2
	}).Catch(func(err error) {
		promiseErr = err
	})

	res, err := p.Wait()
	if err != nil {
		t.Error(err)
	}
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).