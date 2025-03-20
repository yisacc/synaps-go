// Package corporate is a package used to interact with corporate sessions in the Synaps API.
package corporate

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yisacc/synaps-go/pkg/common"
)

const DefaultBaseURL = "https://api.synaps.io/v4"

// Corporate Client type
type Client common.Client

// Create new client with custom API URL
func NewCustomClient(baseURL string, apiKey string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
		BaseURL:    baseURL,
	}
}

// Create new client with API key
func NewClient(apiKey string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
		BaseURL:    DefaultBaseURL,
	}
}

// Create new client from env
func NewClientFromEnv() *Client {
	godotenv.Load()

	apiKey, ok := os.LookupEnv("SYNAPS_API_KEY")
	if !ok {
		log.Fatalf("Missing required SYNAPS_API_KEY env variable")
	}

	baseURL, ok := os.LookupEnv("SYNAPS_BASE_URL")
	if !ok {
		baseURL = DefaultBaseURL
	}

	return NewCustomClient(baseURL, apiKey)
}
