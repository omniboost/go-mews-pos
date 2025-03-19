package go_mews_pos

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestProductGetAllRequest(t *testing.T) {
	products, err := testClient.NewProductGetAllRequest(
		ProductsWithPageSize(20),
		ProductsWithIncludeProductType(),
		ProductsWithIncludeVariants(),
		ProductsWithIncludeModifierSets(),
	).Do(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	jsonBody, _ := jsoniter.MarshalIndent(products, "", "  ")
	t.Log(string(jsonBody))
}
