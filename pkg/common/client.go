// Package common contains shared code between individual and corporate package
package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Base Client type
type Client struct {
	HttpClient *http.Client
	ApiKey     string
	BaseURL    string
}

// MakeRequest is a generic function that create http request and send it, it then serialize response to the provided type
func MakeRequest[T any](httpClient *http.Client, method string, path string, body io.Reader, headers map[string]string) (*T, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create init session request: %w", err)
	}

	for key, header := range headers {
		req.Header.Add(key, header)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make init session request: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var error Error
		if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error output: %w", err)
		}
		defer res.Body.Close()

		return nil, fmt.Errorf("request failed with status code %d: %s", res.StatusCode, error.Message)
	}

	var output T
	if err := json.NewDecoder(res.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("failed to unmarshal output: %w", err)
	}
	defer res.Body.Close()

	return &output, nil
}
