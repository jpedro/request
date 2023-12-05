/*
Fluent API for http requests

A fluent wrapper for https://pkg.go.dev/net/http

Example

	```go
	package main

	import (
		"fmt"

		"github.com/jpedro/request"
	)

	func main() {
		res, err := request.Get("https://example.com/").
			Payload("some stuff").
			Header("X-Foo", "bar").
			Param("baz", "qux").
			Timeout(1).
		Run()

		if err != nil {
			panic(err)
		}

		fmt.Printf("Status code: %d\n", res.StatusCode)
		fmt.Printf("Body length: %d bytes\n", len(res.Body))
	}
	```
*/

package request
