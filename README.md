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
    req := request.Get("https://www.google.com")
    res, err := req.Run()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Response: %#v\n", res)
}
```
