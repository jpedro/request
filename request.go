package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HeaderContentType     = "Content-Type"
	HeaderContentTypeJson = "application/json"
	HeaderAuthorization   = "Authorization"
)

// The request struct
type Request struct {
	Url     string
	Method  string
	Payload []byte
	Headers map[string]string
	Params  map[string]string
	Timeout int64
}

// The response struct
type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string]any
	Cookies    []*http.Cookie
}

// Creates a new HTTP request
func NewRequest(url, method string) *Request {
	r := &Request{}
	r.Url = url
	r.Method = method
	r.Headers = map[string]string{}
	r.Params = map[string]string{}
	r.Timeout = 2

	return r
}

// Creates a new GET HTTP request
func Get(url string) *Request {
	return NewRequest(url, http.MethodGet)
}

// Creates a new POST HTTP request
func Post(url string) *Request {
	return NewRequest(url, http.MethodPost)
}

// Creates a new PUT HTTP request
func Put(url string) *Request {
	return NewRequest(url, http.MethodPut)
}

// Creates a new DELETE HTTP request
func Delete(url string) *Request {
	return NewRequest(url, http.MethodDelete)
}

// Sets the "Content-Type" header to "application/json"
func (r *Request) UsesJson() *Request {
	r.SetHeader(HeaderContentType, HeaderContentTypeJson)
	return r
}

// Sets the HTTP request timeout in seconds
func (r *Request) SetTimeout(seconds int64) *Request {
	r.Timeout = seconds
	return r
}

// Sets the HTTP request payload
func (r *Request) SetPayload(payload string) *Request {
	r.Payload = []byte(payload)
	return r
}

// Builds the fully HTTP URL with the query string appended
func (r *Request) buildQuery() string {
	values := []string{}
	for key, val := range r.Params {
		values = append(values, key+"="+url.QueryEscape(val))
		// values = append(values, key+"="+val)
	}

	return strings.Join(values, "&")
}

// Set a new query string parameter
func (r *Request) SetParam(key, val string) *Request {
	r.Params[key] = val
	return r
}

// Set new query string parameters from a map
func (r *Request) SetParams(values map[string]string) *Request {
	for key, val := range values {
		r.SetParam(key, val)
	}

	return r
}

// Set a new HTTP header
func (r *Request) SetHeader(key, val string) *Request {
	r.Headers[key] = val
	return r
}

// Set HTTP headers from a map
func (r *Request) SetHeaders(values map[string]string) *Request {
	for key, val := range values {
		r.SetHeader(key, val)
	}

	return r
}

// Return the fully assembled URL with the query string appended
func (r *Request) GetFullUrl() string {
	if len(r.Params) == 0 {
		return r.Url
	}

	return r.Url + "?" + r.buildQuery()
}

// Runs a request and returns a response
func (r *Request) Run() (*Response, error) {
	var err error
	var req *http.Request

	switch r.Method {
	case http.MethodPost, http.MethodPut:
		reader := bytes.NewBuffer(r.Payload)
		req, err = http.NewRequest(r.Method, r.GetFullUrl(), reader)
	default:
		req, err = http.NewRequest(r.Method, r.GetFullUrl(), nil)
	}

	if err != nil {
		return nil, err
	}

	for key, val := range r.Headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: time.Second * time.Duration(r.Timeout),
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

// Returned the response body parsed as JSON
func (r *Response) GetJson() (any, error) {
	var data any
	err := json.Unmarshal(r.Body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
