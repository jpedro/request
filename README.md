# Request

[![Tests](https://github.com/jpedro/request/actions/workflows/test.yaml/badge.svg)](https://github.com/jpedro/request/actions/workflows/test.yaml)

Wrapper for [`net/http`](https://pkg.go.dev/net/http).


## Example

```go
package main

import (
    "github.com/jpedro/request"
)

func main() {
    req, _ := request.Get("https://www.google.com")
    res := req.Run()
    fmt.Println("Response", res)
}
```
