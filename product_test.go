package go_mews_pos

import (
	"context"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestProductGetAllRequest(t *testing.T) {
	products, err := testClient.NewProductGetRequest(
		ProductsWithID("123"),
	).Do(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	jsonBody, _ := jsoniter.MarshalIndent(products, "", "  ")
	t.Log(string(jsonBody))
}
