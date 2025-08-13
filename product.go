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
	ProductGetRequest struct {
		client *MewsPosClient
		ID     string `path:"id"`
	}

	ProductGetRequestOption func(*ProductGetRequest)

	ProductGetResponse struct {
		Product Product `json:"product"`
	}

	Product struct {
		ID                             string       `json:"id"`
		CreatedAt                      time.Time    `json:"createdAt"`
		UpdatedAt                      time.Time    `json:"updatedAt"`
		ProductTypeID                  string       `json:"productTypeId"`
		ProductType                    string       `json:"productType"`
		Brand                          *string      `json:"brand"`
		Name                           string       `json:"name"`
		AlternativeName                *string      `json:"alternativeName"`
		Description                    string       `json:"description"`
		Status                         string       `json:"status"`
		Handle                         string       `json:"handle"`
		SKU                            string       `json:"sku"`
		Barcode                        string       `json:"barcode"`
		RegalCode                      *string      `json:"regalCode"`
		TileLabel                      string       `json:"tileLabel"`
		TileColor                      string       `json:"tileColor"`
		Composite                      bool         `json:"composite"`
		PageTitle                      *string      `json:"pageTitle"`
		MetaDescription                *string      `json:"metaDescription"`
		Asin                           *string      `json:"asin"`
		Purchasable                    bool         `json:"purchasable"`
		Sellable                       bool         `json:"sellable"`
		InventoryTracking              bool         `json:"inventoryTracking"`
		UnavailableUntil               *time.Time   `json:"unavailableUntil"`
		CatalogExternalID              *string      `json:"catalogExternalId"`
		MeasurementQuantity            string       `json:"measurementQuantity"`
		MeasurementUnit                string       `json:"measurementUnit"`
		PurchaseQuantity               string       `json:"purchaseQuantity"`
		PurchasePackageQuantity        string       `json:"purchasePackageQuantity"`
		PurchaseUnit                   string       `json:"purchaseUnit"`
		PurchasePrice                  *string      `json:"purchasePrice"`
		UnitPrice                      *string      `json:"unitPrice"`
		PurchasePriceChangeAlert       *string      `json:"purchasePriceChangeAlert"`
		InvalidFabricationUnitProducts []string     `json:"invalidFabricationUnitProducts"`
		FabricationUnits               []string     `json:"fabricationUnits"`
		Parents                        []Parent     `json:"parents"`
		ParentItems                    []ParentItem `json:"parentItems"`
		ParentVariants                 []string     `json:"parentVariants"`
		Items                          []Item       `json:"items"`
		Suppliers                      []string     `json:"suppliers"`
		ImageURL                       *string      `json:"imageUrl"`
		Type                           string       `json:"type"`
		Taxes                          []Tax        `json:"taxes"`
		PurchaseTaxes                  []string     `json:"purchaseTaxes"`
		Weight                         *string      `json:"weight"`
		RetailPriceExclTax             string       `json:"retailPriceExclTax"`
		Tax                            string       `json:"tax"`
		RetailPriceInclTax             string       `json:"retailPriceInclTax"`
		RegularRetailPrice             *string      `json:"regularRetailPrice"`
		WebshopForPickup               bool         `json:"webshopForPickup"`
		WebshopForShipping             bool         `json:"webshopForShipping"`
		WebshopPosition                *int         `json:"webshopPosition"`
		WebshopClickCollectPosition    *int         `json:"webshopClickCollectPosition"`
		WebshopOrderPayPosition        int          `json:"webshopOrderPayPosition"`
		WebshopForServing              bool         `json:"webshopForServing"`
		Inventory                      []Inventory  `json:"inventory"`
		Variants                       []string     `json:"variants"`
		VariantOptions                 []string     `json:"variantOptions"`
		ModifierSetIDs                 []string     `json:"modifierSetIds"`
		GroupProduct                   *string      `json:"groupProduct"`
		Family                         *string      `json:"family"`
		Subfamily                      *string      `json:"subfamily"`
		FamilyCategory                 *string      `json:"familyCategory"`
		PurchaseFamily                 *string      `json:"purchaseFamily"`
		PurchaseSubfamily              *string      `json:"purchaseSubfamily"`
		PurchaseFamilyCategory         *string      `json:"purchaseFamilyCategory"`
		UnitConversions                []string     `json:"unitConversions"`
		InventoryProductID             *string      `json:"inventoryProductId"`
		InventoryProduct               *string      `json:"inventoryProduct"`
		PurchasableProducts            []string     `json:"purchasableProducts"`
		UnitPriceWap                   *string      `json:"unitPriceWap"`
		PurchasePriceWap               *string      `json:"purchasePriceWap"`
		CatalogProductID               *string      `json:"catalogProductId"`
		CatalogProductSellable         *bool        `json:"catalogProductSellable"`
	}

	Item struct {
		ProductID                string      `json:"productId"`
		ProductVariantID         *string     `json:"productVariantId"`
		Quantity                 string      `json:"quantity"`
		MeasurementQuantity      string      `json:"measurementQuantity"`
		MeasurementUnit          string      `json:"measurementUnit"`
		UnitPrice                *string     `json:"unitPrice"`
		PurchaseQuantity         string      `json:"purchaseQuantity"`
		PurchasePackageQuantity  string      `json:"purchasePackageQuantity"`
		PurchaseUnit             string      `json:"purchaseUnit"`
		PurchasePrice            *string     `json:"purchasePrice"`
		PurchasePriceChangeAlert *string     `json:"purchasePriceChangeAlert"`
		Name                     string      `json:"name"`
		Deleted                  bool        `json:"deleted"`
		FabricationUnit          string      `json:"fabricationUnit"`
		FabricationUnits         []string    `json:"fabricationUnits"`
		UnitConversions          []string    `json:"unitConversions"`
		Inventory                []Inventory `json:"inventory"`
	}

	Inventory struct {
		WarehouseID  string  `json:"warehouseId"`
		OnHand       string  `json:"onHand"`
		Committed    string  `json:"committed"`
		Available    string  `json:"available"`
		Incoming     string  `json:"incoming"`
		UnitPrice    *string `json:"unitPrice"`
		CurrentValue string  `json:"currentValue"`
		SalesValue   string  `json:"salesValue"`
	}

	Tax struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Rate        string  `json:"rate"`
		Active      bool    `json:"active"`
		TaxType     string  `json:"taxType"`
		PmsCode     string  `json:"pmsCode"`
		RatePercent string  `json:"ratePercent"`
		Label       *string `json:"label"`
	}

	Parent struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Deleted bool   `json:"deleted"`
	}

	ParentItem struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		Quantity        string `json:"quantity"`
		FabricationUnit string `json:"fabricationUnit"`
	}
)

var _ client.Request = (*ProductGetRequest)(nil)

func (r *ProductGetRequest) Method() string {
	return http.MethodGet
}

func (r *ProductGetRequest) PathTemplate() string {
	return "/v1/products/{{.id}}"
}

func ProductsWithID(ID string) ProductGetRequestOption {
	return func(r *ProductGetRequest) {
		r.ID = strings.TrimSpace(ID)
	}
}

func (m *MewsPosClient) NewProductGetRequest(options ...ProductGetRequestOption) *ProductGetRequest {
	r := &ProductGetRequest{
		client: m,
	}

	for _, o := range options {
		o(r)
	}

	return r
}

func (r *ProductGetRequest) Do(ctx context.Context) (*Product, error) {
	if r.ID == "" {
		return nil, errors.New("product ID is required")
	}

	var responseBody ProductGetResponse
	err := r.client.Do(ctx, r, &responseBody)
	if err != nil {
		return nil, err
	}
	return &responseBody.Product, nil

}
