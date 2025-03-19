package go_mews_pos

import (
	"context"
	"errors"
	"github.com/omniboost/go-omniboost-http-client/client"
	"net/http"
	"strings"
	"time"
)

type (
	ProductGetAllRequest struct {
		client    *MewsPosClient
		PageSize  int    `query:"page[size]"`
		PageAfter string `query:"page[after],omitempty"`
		Include   string `query:"include,omitempty"`
	}

	ProductGetAllRequestOption func(*ProductGetAllRequest)

	ProductGetAllResponse struct {
		Data     []Product  `json:"data"`
		Included []Included `json:"included"`
		Links    struct {
			Prev *string `json:"prev"`
			Next *string `json:"next"`
		} `json:"links"`
	}
	ProductRelationModifierSets struct {
		Data []ModifierSet `json:"data"`
	}

	ModifierSet struct {
		ID         string                 `json:"id"`
		Type       string                 `json:"type"`
		Attributes *ModifierSetAttributes `json:"attributes"`
	}

	ModifierSetAttributes struct {
		Name         string    `json:"name"`
		Selection    string    `json:"selection"`
		MaximumCount int       `json:"maximumCount"`
		MinimumCount int       `json:"minimumCount"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
	}
	ProductRelationProductType struct {
		Data ProductType `json:"data"`
	}
	ProductType struct {
		ID         string                 `json:"id"`
		Type       string                 `json:"type"`
		Attributes *ProductTypeAttributes `json:"attributes,omitempty"`
	}
	ProductTypeAttributes struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
	ProductRelationProductVariants struct {
		Data []ProductVariant `json:"data"`
	}

	ProductVariant struct {
		ID         string                    `json:"id"`
		Type       string                    `json:"type"`
		Attributes *ProductVariantAttributes `json:"attributes,omitempty"`
	}
	ProductVariantAttributes struct {
		RetailPriceExclTax string    `json:"retailPriceExclTax"`
		RetailPriceInclTax string    `json:"retailPriceInclTax"`
		RegularRetailPrice string    `json:"regularRetailPrice"`
		SKU                string    `json:"sku"`
		Barcode            string    `json:"barcode"`
		CreatedAt          time.Time `json:"createdAt"`
		UpdatedAt          time.Time `json:"updatedAt"`
	}

	ProductRelationships struct {
		ProductType     ProductRelationProductType     `json:"productType"`
		ModifierSets    ProductRelationModifierSets    `json:"modifierSets"`
		ProductVariants ProductRelationProductVariants `json:"productVariants"`
	}
	Product struct {
		ID            string                `json:"id"`
		Type          string                `json:"type"`
		Attributes    *ProductAttributes    `json:"attributes,omitempty"`
		Relationships *ProductRelationships `json:"relationships,omitempty"`
	}

	ProductAttributes struct {
		Name               string    `json:"name"`
		Description        *string   `json:"description"`
		SKU                *string   `json:"sku"`
		Barcode            *string   `json:"barcode"`
		Status             string    `json:"status"`
		Tax                string    `json:"tax"`
		RetailPriceExclTax string    `json:"retailPriceExclTax"`
		RetailPriceInclTax string    `json:"retailPriceInclTax"`
		RegularRetailPrice *string   `json:"regularRetailPrice"`
		UnitPrice          *string   `json:"unitPrice"`
		CreatedAt          time.Time `json:"createdAt"`
		UpdatedAt          time.Time `json:"updatedAt"`
	}
)

var _ client.Request = (*ProductGetAllRequest)(nil)

func (r *ProductGetAllRequest) Method() string {
	return http.MethodGet
}

func (r *ProductGetAllRequest) PathTemplate() string {
	return "/api/v2/products"
}

func ProductsWithPageSize(pageSize int) ProductGetAllRequestOption {
	return func(r *ProductGetAllRequest) {
		r.PageSize = pageSize
	}
}

func ProductsWithPageAfter(pageAfter string) ProductGetAllRequestOption {
	return func(r *ProductGetAllRequest) {
		r.PageAfter = pageAfter
	}
}

func ProductsWithIncludeProductType() ProductGetAllRequestOption {
	return func(r *ProductGetAllRequest) {
		r.Include = strings.Trim(r.Include+",productType", ",")
	}
}

func ProductsWithIncludeVariants() ProductGetAllRequestOption {
	return func(r *ProductGetAllRequest) {
		r.Include = strings.Trim(r.Include+",productVariants", ",")
	}
}

func ProductsWithIncludeModifierSets() ProductGetAllRequestOption {
	return func(r *ProductGetAllRequest) {
		r.Include = strings.Trim(r.Include+",modifierSets", ",")
	}
}

func (m *MewsPosClient) NewProductGetAllRequest(options ...ProductGetAllRequestOption) *ProductGetAllRequest {
	r := &ProductGetAllRequest{
		client:   m,
		PageSize: 100,
	}

	for _, o := range options {
		o(r)
	}

	return r
}

func (r *ProductGetAllRequest) Do(ctx context.Context) ([]Product, error) {
	var products []Product

	maxRequests := 1000
	for {
		var responseBody ProductGetAllResponse
		err := r.client.Do(ctx, r, &responseBody)
		if err != nil {
			return nil, err
		}

		if len(responseBody.Data) == 0 {
			break
		}

		includedProductTypes, err := getIncludedByType[ProductType](r.client, responseBody.Included, "productTypes")
		if err != nil {
			return nil, err
		}
		includedProductVariants, err := getIncludedByType[ProductVariant](r.client, responseBody.Included, "productVariants")
		if err != nil {
			return nil, err
		}
		includedModifierSets, err := getIncludedByType[ModifierSet](r.client, responseBody.Included, "modifierSets")
		if err != nil {
			return nil, err
		}

		for i, product := range responseBody.Data {
			productType, ok := includedProductTypes[product.Relationships.ProductType.Data.ID]
			if ok {
				product.Relationships.ProductType.Data = productType
			}
			for j, modifierSet := range product.Relationships.ModifierSets.Data {
				modifierSet, ok := includedModifierSets[modifierSet.ID]
				if ok {
					product.Relationships.ModifierSets.Data[j] = modifierSet
				}
			}
			for j, productVariant := range product.Relationships.ProductVariants.Data {
				productVariant, ok := includedProductVariants[productVariant.ID]
				if ok {
					product.Relationships.ProductVariants.Data[j] = productVariant
				}
			}

			responseBody.Data[i] = product
		}
		products = append(products, responseBody.Data...)

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
	return products, nil

}
