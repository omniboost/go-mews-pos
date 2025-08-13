package go_mews_pos

import (
	"net/url"

	"github.com/omniboost/go-omniboost-http-client/client"
)

type (
	MewsPosClient struct {
		client.Client
	}

	Links struct {
		Prev *string `json:"prev"`
		Next *string `json:"next"`
	}

	SelfLinks struct {
		Self string `json:"self"`
	}

	Included struct {
		ID           string         `json:"id"`
		Type         string         `json:"type"`
		Attributes   map[string]any `json:"attributes"`
		Relationship map[string]any `json:"relationships,omitempty"`
		Links        *SelfLinks     `json:"links,omitempty"`
	}
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-mews-pos/" + libraryVersion
)

var DefaultBaseURL = url.URL{
	Scheme: "https",
	Host:   "api.mews-demo.com",
	Path:   "/pos",
}

func NewMewsPosClient(opts ...client.Option) *MewsPosClient {
	opts = append([]client.Option{
		client.WithBaseURL(DefaultBaseURL),
		client.WithUserAgent(userAgent),
	}, opts...)

	return &MewsPosClient{
		Client: client.NewClient(opts...),
	}
}

func WithApiKey(apiKey string) client.Option {
	return client.WithApiKeyAuth(
		"Authorization",
		"Bearer "+apiKey,
	)
}

func WithLegacyApiKey(apiKey string) client.Option {
	return client.WithApiKeyAuth(
		"X-AccessToken",
		apiKey,
	)
}

// getIncludedByType returns the included data by type, indexed by ID
func getIncludedByType[T any](m *MewsPosClient, data []Included, typeName string) (map[string]T, error) {
	jsoniterInstance := m.GetJsoniter()
	result := make(map[string]T)
	for _, included := range data {
		if included.Type != typeName {
			continue
		}
		raw, _ := jsoniterInstance.Marshal(included)
		var t T
		if err := jsoniterInstance.Unmarshal(raw, &t); err != nil {
			return nil, err
		}
		result[included.ID] = t
	}

	return result, nil
}

func getPageAfter(links Links) (string, error) {
	if links.Next == nil {
		return "", nil
	}

	u, err := url.Parse(*links.Next)
	return u.Query().Get("page[after]"), err
}
