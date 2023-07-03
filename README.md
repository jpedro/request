# Request

[![Tests](https://github.com/jpedro/request/actions/workflows/test.yaml/badge.svg)](https://github.com/jpedro/request/actions/workflows/test.yaml)

Wrapper for [`net/http`](https://pkg.go.dev/net/http).


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
		SetPayload("some stuff").
		SetHeader("X-Foo", "bar").
		SetParam("baz", "qux").
		SetTimeout(1)

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status code: %d\n", res.StatusCode)
	fmt.Printf("Body length: %d bytes\n", len(res.Body))
}
```
