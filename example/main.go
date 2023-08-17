package main

import (
	"fmt"
	"strings"

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

	fmt.Printf("\033[32;1mStatus code:\033[0m %d\n", res.StatusCode)
	fmt.Printf("\033[32;1mBody length:\033[0m %d bytes\n", len(res.Body))
	fmt.Printf("\033[32;1mBody sample:\033[0m\n")
	lines := strings.Split(string(res.Body[0:1000]), "\n")
	for i, line := range lines {
		fmt.Printf("\033[2m %5d │\033[0m %s\033[0m\n", i, line)
	}
	// fmt.Printf("\033[32;1m¶\033[0m\n",)
}
