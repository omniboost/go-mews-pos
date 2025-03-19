package go_mews_pos

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestRegisterGetAllRequest(t *testing.T) {
	registers, err := testClient.NewRegisterGetAllRequest(
		RegisterWithID("c7d3104e-cebf-4fde-8456-69190c1a082b"),
		RegisterWithIncludeOutlet(),
	).Do(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	jsonBody, _ := jsoniter.MarshalIndent(registers, "", "  ")
	t.Log(string(jsonBody))
}
