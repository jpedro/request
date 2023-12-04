package request

import (
	"fmt"
	"testing"
)

const (
	TEST_URL = "https://dummyjson.com"
)

func TestRequestGet(t *testing.T) {
	url := TEST_URL + "/products"
	req := Get(url)
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to load products: %s\n", err)
	}

	products := data.(map[string]any)["products"].([]any)
	fmt.Println("data:", len(products))
}

func TestRequestParams(t *testing.T) {
	url := TEST_URL + "/products/search"
	req := Get(url)
	req.Params(map[string]string{
		"q": "phone",
	})
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to load products: %s\n", err)
	}

	products := data.(map[string]any)["products"].([]any)
	fmt.Println("data:", len(products))
}

func TestRequestHeaders(t *testing.T) {
	url := TEST_URL + "/products/search"
	req := Get(url)
	req.Headers(map[string]string{
		"X-1": "Something something",
	})
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to load products: %s\n", err)
	}

	products := data.(map[string]any)["products"].([]any)
	fmt.Println("data:", len(products))
}

func TestRequestTimeout(t *testing.T) {
	url := TEST_URL + "/products"
	req := Get(url).Timeout(0)
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to load products: %s\n", err)
	}

	products := data.(map[string]any)["products"].([]any)
	fmt.Println("data:", len(products))
}

func TestRequestPostJson(t *testing.T) {
	url := TEST_URL + "/products/add"
	req := Post(url)

	payload := map[string]any{
		"title": "BMW Pencil as JSON",
	}

	req.Payload(payload)
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to post product: %s\n", err)
	}
	fmt.Println("Created product:", data)
}

func TestRequestPostText(t *testing.T) {
	url := TEST_URL + "/products/add"
	req := Post(url)

	payload := `
	{
		"title": "BMW Pencil as JSON"
	}
	`

	req.Payload(payload)
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to post product: %s\n", err)
	}
	fmt.Println("Created product:", data)
}

func TestRequestPostBytes(t *testing.T) {
	url := TEST_URL + "/products/add"
	req := Post(url)

	payload := []byte(`
	{
		"title": "BMW Pencil as bytes"
	}
	`)

	req.Payload(payload)
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to post product: %s\n", err)
	}
	fmt.Println("Created product:", data)
}

func TestRequestPut(t *testing.T) {
	url := TEST_URL + "/products/1"
	payload := map[string]any{
		"title": "UPDATED TITLE",
	}

	req := Put(url)
	req.Payload(payload)

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to update product: %s\n", err)
	}
	fmt.Println("Updated product:", data)
}

func TestRequestDelete(t *testing.T) {
	url := TEST_URL + "/products/1"
	req := Delete(url)

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	statusCode := res.StatusCode
	if statusCode != 200 {
		t.Errorf("Failed to delete product statusCode: %d\n", statusCode)
	}

	data, err := res.Json()
	if err != nil {
		t.Errorf("Failed to delete product payload: %s\n", err)
	}
	fmt.Println("Deleted payload:", data)
}

func TestRequest404(t *testing.T) {
	url := TEST_URL + "/404"
	req := Get(url)

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	statusCode := res.StatusCode
	if statusCode != 404 {
		t.Errorf("Failed to get a 404: %d\n", statusCode)
	}

	fmt.Println("Got status code:", res.StatusCode)
}
