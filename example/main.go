package main

import (
	"fmt"

	"github.com/jpedro/request"
)

func main() {
    req := request.Get("https://www.google.com")
    res, err := req.Run()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Response: %#v\n", res)
}
