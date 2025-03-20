package corporate

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/yisacc/synaps-go/pkg/common"
)

type (
	// Init session function params
	//
	// see: [common.InitSessionParams]
	//
	// function: [corporate.InitSession]
	InitSessionParams common.InitSessionParams
	// Init session response
	//
	// see: [common.InitSessionResponse]
	//
	// function: [corporate.InitSession]
	InitSessionResponse common.InitSessionResponse
)

// Init new session with alias and metadata, if provided
func (c *Client) InitSession(params InitSessionParams) (sessionID InitSessionResponse, err error) {
	headers := map[string]string{"Api-Key": c.ApiKey, "Content-Type": "application/json"}
	body, err := json.Marshal(params)
	if err != nil {
		return InitSessionResponse{}, fmt.Errorf("failed to marshal input: %s", err)
	}

	res, err := common.MakeRequest[InitSessionResponse](c.HttpClient, "POST", c.BaseURL+"/session/init", bytes.NewReader(body), headers)
	if err != nil {
		return InitSessionResponse{}, fmt.Errorf("init session request failed: %s", err)
	}
	return *res, nil
}
