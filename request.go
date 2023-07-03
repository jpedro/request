package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	HeaderContentType     = "Content-Type"
	HeaderContentTypeJson = "application/json"
	HeaderAuthorization   = "Authorization"
)

type Request struct {
	Url     string
	Method  string
	Payload []byte
	Headers map[string]string
	Params  map[string]string
	Timeout int64
}

type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string]any
	Cookies    []*http.Cookie
}

func NewRequest(url, method string) *Request {
	r := &Request{}
	r.Url = url
	r.Method = method
	r.Headers = map[string]string{}
	r.Params = map[string]string{}
	r.Timeout = 2

	return r
}

func Get(url string) *Request {
	return NewRequest(url, http.MethodGet)
}

func Post(url string) *Request {
	return NewRequest(url, http.MethodPost)
}

func Put(url string) *Request {
	return NewRequest(url, http.MethodPut)
}

func Delete(url string) *Request {
	return NewRequest(url, http.MethodDelete)
}

func (r *Request) UsesJson() *Request {
	r.SetHeader(HeaderContentType, HeaderContentTypeJson)
	return r
}

func (r *Request) SetTimeout(seconds int64) *Request {
	r.Timeout = seconds
	return r
}

func (r *Request) SetPayload(payload string) *Request {
	r.Payload = []byte(payload)
	return r
}

func (r *Request) buildQuery() string {
	values := []string{}
	for key, val := range r.Params {
		// values = append(values, key+"="+url.QueryEscape(val))
		values = append(values, key+"="+val)
	}

	return strings.Join(values, "&")
}

func (r *Request) SetParam(key, val string) *Request {
	r.Params[key] = val
	return r
}

func (r *Request) SetParams(values map[string]string) *Request {
	for key, val := range values {
		r.SetParam(key, val)
	}

	return r
}

func (r *Request) SetHeader(key, val string) *Request {
	r.Headers[key] = val
	return r
}

func (r *Request) SetHeaders(values map[string]string) *Request {
	for key, val := range values {
		r.Headers[key] = val
	}

	return r
}

func (r *Request) GetFullUrl() string {
	if len(r.Params) == 0 {
		return r.Url
	}

	return r.Url + "?" + r.buildQuery()
}

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

func (r *Response) GetJson() (any, error) {
	var data any
	err := json.Unmarshal(r.Body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
