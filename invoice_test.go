package go_mews_pos

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestInvoiceGetAllRequest(t *testing.T) {
	invoices, err := testClient.NewInvoiceGetAllRequest(
		InvoicesWithPageSize(100),
		InvoicesWithRegisterIdEq("c7d3104e-cebf-4fde-8456-69190c1a082b"),
		//InvoicesWithIncludeInvoiceItems(),
		//InvoicesWithIncludeUser(),
		//InvoicesWithIncludeRegister(),
		//InvoicesWithIncludeOriginalInvoice(),
	).Do(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	jsonBody, _ := jsoniter.MarshalIndent(invoices, "", "  ")
	t.Log(string(jsonBody))
}
