package go_mews_pos

import (
	"context"
	"net/http"
	"time"
)

type (
	RegisterGetRequest struct {
		client     *MewsPosClient
		RegisterID string `path:"register_id"`
		Include    string `query:"include,omitempty"`
	}

	RegisterGetResponse struct {
		Data     Register   `json:"data"`
		Included []Included `json:"included"`
	}

	RegisterGetOption func(*RegisterGetRequest)

	Register struct {
		ID            string                 `json:"id"`
		Type          string                 `json:"type"`
		Attributes    *RegisterAttributes    `json:"attributes,omitempty"`
		SelfLinks     *SelfLinks             `json:"links,omitempty"`
		Relationships *RegisterRelationships `json:"relationships,omitempty"`
	}

	RegisterRelationships struct {
		RegisterRelationshipOutlet RegisterRelationOutlet `json:"outlet"`
	}

	RegisterRelationOutlet struct {
		Data Outlet `json:"data"`
	}

	Outlet struct {
		ID         string            `json:"id"`
		Type       string            `json:"type"`
		Attributes *OutletAttributes `json:"attributes,omitempty"`
	}

	RegisterAttributes struct {
		Name          string    `json:"name"`
		InvoicesCount int       `json:"invoicesCount"`
		Index         int       `json:"index"`
		Virtual       bool      `json:"virtual"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
	}

	OutletAttributes struct {
		Name       string    `json:"name"`
		Address1   *string   `json:"address1"`
		Address2   *string   `json:"address2"`
		City       *string   `json:"city"`
		State      *string   `json:"state"`
		PostalCode *string   `json:"postalCode"`
		Index      int       `json:"index"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	}
)

func RegisterWithIncludeOutlet() RegisterGetOption {
	return func(r *RegisterGetRequest) {
		r.Include = "outlet"
	}
}

func RegisterWithID(registerID string) RegisterGetOption {
	return func(r *RegisterGetRequest) {
		r.RegisterID = registerID
	}
}

func (m *MewsPosClient) NewRegisterGetRequest(opts ...RegisterGetOption) *RegisterGetRequest {
	r := &RegisterGetRequest{
		client: m,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *RegisterGetRequest) Method() string {
	return http.MethodGet
}

func (r *RegisterGetRequest) PathTemplate() string {
	return "/api/v2/registers/{{.register_id}}"
}

func (r *RegisterGetRequest) Do(ctx context.Context) (*Register, error) {
	var resp RegisterGetResponse
	err := r.client.Do(ctx, r, &resp)
	if err != nil {
		return nil, err
	}

	outlets, err := getIncludedByType[Outlet](r.client, resp.Included, "outlets")
	if err != nil {
		return nil, err
	}

	if resp.Data.Relationships != nil {
		if outlet, ok := outlets[resp.Data.Relationships.RegisterRelationshipOutlet.Data.ID]; ok {
			resp.Data.Relationships.RegisterRelationshipOutlet.Data = outlet
		}
	}

	return &resp.Data, nil
}
