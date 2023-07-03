package main

import (
	"fmt"

	"github.com/jpedro/request"
)

func main() {
	req := request.Get("https://example.com/")
	req.SetHeader("X-Foo", "bar").SetParam("baz", "qux")

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status code: %d\n", res.StatusCode)
	fmt.Printf("Content length: %d bytes\n", len(res.Body))
}
