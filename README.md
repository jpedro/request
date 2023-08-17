# Request

[![Tests](https://github.com/jpedro/request/actions/workflows/test.yaml/badge.svg)](https://github.com/jpedro/request/actions/workflows/test.yaml)

A fluent wrapper for [`net/http`](https://pkg.go.dev/net/http).


## Example

```go
package main

import (
	"fmt"

	"github.com/jpedro/request"
)

func main() {
	req := request.Get("https://example.com/").
		UsesJson().
		WithPayload("some stuff").
		WithHeader("X-Foo", "bar").
		WithParam("baz", "qux").
		WithTimeout(1)

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status code: %d\n", res.StatusCode)
	fmt.Printf("Body length: %d bytes\n", len(res.Body))
}
```
