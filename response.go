package request

import (
	"encoding/json"
	"net/http"
)

// The response struct
type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string]any
	Cookies    []*http.Cookie
}

// Returned the response body parsed as JSON.
func (r *Response) Json() (any, error) {
	var data any
	err := json.Unmarshal(r.Body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Returned the response body parsed as JSON.
func (r *Response) Text() string {
	return string(r.Body)
}
