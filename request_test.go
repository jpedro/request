package request

import (
	"fmt"
	"testing"
)

func TestRequestGet(t *testing.T) {
	url := "https://dummyjson.com/products"
	req := Get(url)
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.GetJson()
	if err != nil {
		t.Errorf("Failed to load products: %s\n", err)
	}

	products := data.(map[string]any)["products"].([]any)
	fmt.Println("data:", len(products))
}

func TestRequestSetParams(t *testing.T) {
	url := "https://dummyjson.com/products/search"
	req := Get(url)
	req.SetParams(map[string]string{
		"q": "phone",
	})
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.GetJson()
	if err != nil {
		t.Errorf("Failed to load products: %s\n", err)
	}

	products := data.(map[string]any)["products"].([]any)
	fmt.Println("data:", len(products))
}

func TestRequestSetHeaders(t *testing.T) {
	url := "https://dummyjson.com/products/search"
	req := Get(url)
	req.SetHeaders(map[string]string{
		"X-1": "Something something",
	})
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.GetJson()
	if err != nil {
		t.Errorf("Failed to load products: %s\n", err)
	}

	products := data.(map[string]any)["products"].([]any)
	fmt.Println("data:", len(products))
}

func TestRequestTimeout(t *testing.T) {
	url := "https://dummyjson.com/products"
	req := Get(url).SetTimeout(0)
	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.GetJson()
	if err != nil {
		t.Errorf("Failed to load products: %s\n", err)
	}

	products := data.(map[string]any)["products"].([]any)
	fmt.Println("data:", len(products))
}

func TestRequestPost(t *testing.T) {
	url := "https://dummyjson.com/products/add"
	req := Post(url)
	payload := `
	{
		"title": "BMW Pencil"
	}
	`

	req.UsesJson()
	req.SetPayload(payload)

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.GetJson()
	if err != nil {
		t.Errorf("Failed to post product: %s\n", err)
	}
	fmt.Println("Created product:", data)
}

func TestRequestPut(t *testing.T) {
	url := "https://dummyjson.com/products/1"
	req := Put(url)
	payload := `
	{
		"title": "UPDATED TITLE"
	}
	`

	req.UsesJson()
	req.SetPayload(payload)

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	data, err := res.GetJson()
	if err != nil {
		t.Errorf("Failed to update product: %s\n", err)
	}
	fmt.Println("Updated product:", data)
}

func TestRequestDelete(t *testing.T) {
	url := "https://dummyjson.com/products/1"
	req := Delete(url)

	res, err := req.Run()
	if err != nil {
		panic(err)
	}

	statusCode := res.StatusCode
	if statusCode != 200 {
		t.Errorf("Failed to delete product statusCode: %d\n", statusCode)
	}

	data, err := res.GetJson()
	if err != nil {
		t.Errorf("Failed to delete product payload: %s\n", err)
	}
	fmt.Println("Deleted payload:", data)
}
