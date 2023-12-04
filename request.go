package request

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HeaderContentType     = "Content-Type"
	HeaderAccept          = "Accept"
	HeaderContentTypeJson = "application/json"
	HeaderContentTypeText = "text/plain"
	HeaderAuthorization   = "Authorization"

	// AcceptsJson = "json"
)

// var (
// 	Accepts = map[string]map[string]string{
// 		"json": {
// 			"Content-type": "application/json",
// 		},
// 	}
// )

// The request struct
type Request struct {
	url     string
	method  string
	payload []byte
	headers map[string]string
	params  map[string]string
	timeout int64
}

// Creates a new HTTP request
func NewRequest(url, method string) *Request {
	r := &Request{}
	r.url = url
	r.method = method
	r.headers = map[string]string{}
	r.params = map[string]string{}
	r.timeout = 2

	return r
}

// Creates a new GET HTTP request.
func Get(url string) *Request {
	return NewRequest(url, http.MethodGet)
}

// Creates a new POST HTTP request.
func Post(url string) *Request {
	return NewRequest(url, http.MethodPost)
}

// Creates a new PUT HTTP request.
func Put(url string) *Request {
	return NewRequest(url, http.MethodPut)
}

// Creates a new DELETE HTTP request.
func Delete(url string) *Request {
	return NewRequest(url, http.MethodDelete)
}

// Encodes fields.
func EncodeFields(fields map[string]string) string {
	values := []string{}
	for key, val := range fields {
		// values = append(values, key+"="+url.QueryEscape(fmt.Sprintf("%v", val)))
		values = append(values, key+"="+url.QueryEscape(val))
	}

	return strings.Join(values, "&")
}

// // Sets both the "Accept" and the "Content-Type" header to "application/json".
// func (r *Request) UsesJson() *Request {
// 	// r.SendsJson()
// 	r.AcceptsJson()
// 	return r
// }

// // Sets the "Accept" header to "application/json".
// func (r *Request) AcceptsJson() *Request {
// 	r.WithHeader(HeaderAccept, HeaderContentTypeJson)
// 	return r
// }

// Sets the "Content-Type" header to "application/json".
func (r *Request) Json(data map[string]any) *Request {
	r.Header(HeaderContentType, HeaderContentTypeJson)

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	r.payload = bytes

	return r
}

func (r *Request) Text(data string) *Request {
	r.Header(HeaderContentType, HeaderContentTypeText)
	r.payload = []byte(data)
	return r
}

// Sets the HTTP request timeout in seconds.
func (r *Request) Timeout(seconds int64) *Request {
	r.timeout = seconds
	return r
}

// Sets the HTTP request payload.
// func (r *Request) Payload(payload []byte) *Request {
func (r *Request) Payload(data any) *Request {

	switch value := data.(type) {
	case string:
		return r.Text(value)
	case []byte:
		r.payload = value
		return r
	case map[string]any:
		return r.Json(value)
	default:
		log.Fatalf("I don't know how to handle this data: %#v", data)
		panic("And... we stop here")
	}
}

// Set a new query string parameter.
func (r *Request) Param(key, val string) *Request {
	r.params[key] = val
	return r
}

// Set new query string parameters from a map.
func (r *Request) Params(values map[string]string) *Request {
	for key, val := range values {
		r.Param(key, val)
	}

	return r
}

// Set a new HTTP header.
func (r *Request) Header(key, val string) *Request {
	r.headers[key] = val
	return r
}

// Set HTTP headers from a map.
func (r *Request) Headers(values map[string]string) *Request {
	for key, val := range values {
		r.Header(key, val)
	}

	return r
}

// Return the fully assembled URL with the query string appended.
func (r *Request) FullUrl() string {
	if len(r.params) == 0 {
		return r.url
	}

	return r.url + "?" + EncodeFields(r.params)
}

// Runs a request and returns a response.
func (r *Request) Run() (*Response, error) {
	var err error
	var req *http.Request

	url := r.FullUrl()

	switch r.method {
	case http.MethodPost, http.MethodPut:
		reader := bytes.NewBuffer(r.payload)
		req, err = http.NewRequest(r.method, url, reader)
	default:
		req, err = http.NewRequest(r.method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	for key, val := range r.headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: time.Second * time.Duration(r.timeout),
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body := []byte{}
	if res.Body != nil {
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
	}

	headers := map[string]any{}
	for key, val := range res.Header {
		headers[key] = val
	}

	response := &Response{}
	response.StatusCode = res.StatusCode
	response.Headers = headers
	response.Cookies = res.Cookies()
	response.Body = body

	return response, nil
}
