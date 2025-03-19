package go_mews_pos

import (
	"github.com/omniboost/go-omniboost-http-client/client"
	"log"
	"net/url"
	"os"
	"testing"
)

var testClient *MewsPosClient

func TestMain(m *testing.M) {
	opts := []client.Option{
		client.WithDisallowUnknownFields(true),
	}
	if os.Getenv("API_KEY") != "" {
		opts = append(opts, WithApiKey(
			os.Getenv("API_KEY"),
		))
	}
	if os.Getenv("DEBUG") != "" {
		opts = append(opts, client.WithDebug(true))
	}
	if baseUrlString := os.Getenv("BASE_URL"); baseUrlString != "" {
		baseUrl, err := url.Parse(baseUrlString)
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, client.WithBaseURL(*baseUrl))
	}
	testClient = NewMewsPosClient(opts...)

	os.Exit(m.Run())
}
