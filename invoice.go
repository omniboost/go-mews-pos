package go_mews_pos

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/omniboost/go-omniboost-http-client/client"
)

type (
	InvoiceGetAllRequest struct {
		client    *MewsPosClient
		PageSize  int    `query:"page[size]"`
		PageAfter string `query:"page[after],omitempty"`
		Include   string `query:"include,omitempty"`

		CreatedAtGt   *Time   `query:"filter[createdAtGt],omitempty"`
		CreatedAtGtEq *Time   `query:"filter[createdAtGteq],omitempty"`
		CreatedAtLt   *Time   `query:"filter[createdAtLt],omitempty"`
		CreatedAtLtEq *Time   `query:"filter[createdAtLteq],omitempty"`
		RegisterIdEq  *string `query:"filter[registerIdEq],omitempty"`
	}

	InvoiceGetAllRequestOption func(*InvoiceGetAllRequest)

	InvoiceGetAllResponse struct {
		Data     []Invoice  `json:"data"`
		Included []Included `json:"included"`
		Links    Links      `json:"links"`
	}

	Invoice struct {
		ID            string               `json:"id"`
		Type          string               `json:"type"`
		Attributes    InvoiceAttributes    `json:"attributes"`
		Relationships InvoiceRelationships `json:"relationships"`
	}
	InvoiceAttributes struct {
		Discount           *string   `json:"discount"`
		Tax                string    `json:"tax"`
		Total              string    `json:"total"`
		Subtotal           string    `json:"subtotal"`
		TipAmount          *string   `json:"tipAmount"`
		CreatedAt          time.Time `json:"createdAt"`
		UpdatedAt          time.Time `json:"updatedAt"`
		Cancelled          bool      `json:"cancelled"`
		CancelReason       *string   `json:"cancelReason"`
		DiscountAmount     *string   `json:"discountAmount"`
		Description        *string   `json:"description"`
		ItemDiscountAmount *string   `json:"itemDiscountAmount"`
	}
	InvoiceRelationships struct {
		User            InvoiceRelationUser            `json:"user"`
		Register        InvoiceRelationRegisters       `json:"register"`
		OriginalInvoice InvoiceRelationOriginalInvoice `json:"originalInvoice"`
		InvoiceItems    InvoiceRelationInvoiceItems    `json:"invoiceItems"`
		Order           InvoiceRelationOrder           `json:"order"`
	}

	InvoiceRelationInvoiceItems struct {
		Data []InvoiceItem `json:"data"`
	}
	InvoiceItem struct {
		ID            string                    `json:"id"`
		Type          string                    `json:"type"`
		Attributes    *InvoiceItemAttributes    `json:"attributes,omitempty"`
		Relationships *InvoiceItemRelationships `json:"relationships,omitempty"`
		Links         *SelfLinks                `json:"links,omitempty"`
	}

	InvoiceItemAttributes struct {
		ProductName          string    `json:"productName"`
		UnitPriceInclTax     string    `json:"unitPriceInclTax"`
		Quantity             string    `json:"quantity"`
		Subtotal             string    `json:"subtotal"`
		Tax                  string    `json:"tax"`
		Total                string    `json:"total"`
		Discount             *string   `json:"discount"`
		Comp                 bool      `json:"comp"`
		Void                 bool      `json:"void"`
		CompVoidReason       *string   `json:"compVoidReason"`
		CompVoidNotes        *string   `json:"compVoidNotes"`
		DiscountAmount       *string   `json:"discountAmount"`
		SubtotalInclDiscount string    `json:"subtotalInclDiscount"`
		TaxInclDiscount      string    `json:"taxInclDiscount"`
		TotalInclDiscount    string    `json:"totalInclDiscount"`
		CreatedAt            time.Time `json:"createdAt"`
		UpdatedAt            time.Time `json:"updatedAt"`
	}

	InvoiceItemRelationships struct {
		Product              InvoiceItemRelationProduct              `json:"product"`
		ProductVariant       InvoiceItemRelationProductVariant       `json:"productVariant"`
		InvoiceItemModifiers InvoiceItemRelationInvoiceItemModifiers `json:"invoiceItemModifiers"`
	}

	InvoiceRelationOrder struct {
		Data *Order `json:"data"`
	}

	Order struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		//Attributes    *OrderAttributes    `json:"attributes,omitempty"`
		//Relationships *OrderRelationships `json:"relationships,omitempty"`
	}
	InvoiceItemRelationProduct struct {
		Data Product `json:"data"`
	}

	InvoiceItemRelationProductVariant struct {
		Data *InvoiceItemProductVariant `json:"data"`
	}

	InvoiceItemRelationInvoiceItemModifiers struct {
		Data []InvoiceItemModifier `json:"data"`
	}

	InvoiceItemModifier struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}

	InvoiceRelationOriginalInvoice struct {
		Data *OriginalInvoice `json:"data"`
	}

	OriginalInvoice struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}
	InvoiceRelationUser struct {
		Data User `json:"data"`
	}

	User struct {
		ID         string          `json:"id"`
		Type       string          `json:"type"`
		Attributes *UserAttributes `json:"attributes,omitempty"`
	}

	UserAttributes struct {
		Name string `json:"name"`
	}

	InvoiceRelationRegisters struct {
		Data Register `json:"data"`
	}

	InvoiceItemProductVariant struct {
		ID         string                               `json:"id"`
		Type       string                               `json:"type"`
		Attributes *InvoiceItemProductVariantAttributes `json:"attributes,omitempty"`
	}
	InvoiceItemProductVariantAttributes struct {
		RetailPriceExclTax string    `json:"retailPriceExclTax"`
		RetailPriceInclTax string    `json:"retailPriceInclTax"`
		RegularRetailPrice string    `json:"regularRetailPrice"`
		SKU                string    `json:"sku"`
		Barcode            string    `json:"barcode"`
		CreatedAt          time.Time `json:"createdAt"`
		UpdatedAt          time.Time `json:"updatedAt"`
	}
)

var _ client.Request = (*InvoiceGetAllRequest)(nil)

func (r *InvoiceGetAllRequest) Method() string {
	return http.MethodGet
}
func (r *InvoiceGetAllRequest) PathTemplate() string {
	return "/v1/invoices"
}

func InvoicesWithPageSize(pageSize int) InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.PageSize = pageSize
	}
}

func InvoicesWithPageAfter(pageAfter string) InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.PageAfter = pageAfter
	}
}

func InvoicesWithIncludeInvoiceItems() InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.Include = strings.Trim(r.Include+",invoiceItems", ",")
	}
}

func InvoicesWithIncludeUser() InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.Include = strings.Trim(r.Include+",user", ",")
	}
}

func InvoicesWithIncludeRegister() InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.Include = strings.Trim(r.Include+",register", ",")
	}
}

func InvoicesWithIncludeOriginalInvoice() InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.Include = strings.Trim(r.Include+",originalInvoice", ",")
	}
}

func InvoicesWithRegisterIdEq(registerId string) InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.RegisterIdEq = &registerId
	}
}

func InvoicesWithCreatedAtGt(createdAtGt time.Time) InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.CreatedAtGt = &Time{createdAtGt}
	}
}

func InvoicesWithCreatedAtGtEq(createdAtGtEq time.Time) InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.CreatedAtGtEq = &Time{createdAtGtEq}
	}
}

func InvoicesWithCreatedAtLt(createdAtLt time.Time) InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.CreatedAtLt = &Time{createdAtLt}
	}
}

func InvoicesWithCreatedAtLtEq(createdAtLtEq time.Time) InvoiceGetAllRequestOption {
	return func(r *InvoiceGetAllRequest) {
		r.CreatedAtLtEq = &Time{createdAtLtEq}
	}
}

func (m *MewsPosClient) NewInvoiceGetAllRequest(opts ...InvoiceGetAllRequestOption) *InvoiceGetAllRequest {
	r := &InvoiceGetAllRequest{
		client:   m,
		PageSize: 1000,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *InvoiceGetAllRequest) Do(ctx context.Context) ([]Invoice, error) {
	invoices := make([]Invoice, 0)

	maxRequests := 1000
	for {
		var responseBody InvoiceGetAllResponse
		err := r.client.Do(ctx, r, &responseBody)
		if err != nil {
			return nil, err
		}

		if len(responseBody.Data) == 0 {
			break
		}

		includedUsers, err := getIncludedByType[User](r.client, responseBody.Included, "users")
		if err != nil {
			return nil, err
		}
		includedInvoiceItems, err := getIncludedByType[InvoiceItem](r.client, responseBody.Included, "invoiceItems")
		if err != nil {
			return nil, err
		}
		includedRegisters, err := getIncludedByType[Register](r.client, responseBody.Included, "registers")
		if err != nil {
			return nil, err
		}

		for i, invoice := range responseBody.Data {
			if user, ok := includedUsers[invoice.Relationships.User.Data.ID]; ok {
				invoice.Relationships.User.Data = user
			}
			for j, invoiceItem := range invoice.Relationships.InvoiceItems.Data {
				if includedInvoiceItem, ok := includedInvoiceItems[invoiceItem.ID]; ok {
					invoice.Relationships.InvoiceItems.Data[j] = includedInvoiceItem
				}
			}
			if register, ok := includedRegisters[invoice.Relationships.Register.Data.ID]; ok {
				invoice.Relationships.Register.Data = register
			}
			responseBody.Data[i] = invoice

		}
		invoices = append(invoices, responseBody.Data...)

		r.PageAfter, err = getPageAfter(responseBody.Links)
		if err != nil {
			return nil, err
		}
		if r.PageAfter == "" {
			break
		}

		maxRequests--
		if maxRequests == 0 {
			return nil, errors.New("max requests exceeded")
		}
	}
	return invoices, nil
}
