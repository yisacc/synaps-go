package individual

import (
	"fmt"

	"github.com/synaps-hub/synaps-sdk-go/pkg/common"
)

// Get liveness step details from sessionID and stepID
func (c *Client) GetLivenessStep(sessionID string, stepID string) (LivenessStepResponse, error) {
	res, err := common.MakeRequest[LivenessStepResponse](c.HttpClient, "GET", c.BaseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return LivenessStepResponse{}, fmt.Errorf("get liveness step details request failed: %s", err)
	}

	return *res, nil
}

// Get phone step details from sessionID and stepID
func (c *Client) GetPhoneStep(sessionID string, stepID string) (PhoneStepResponse, error) {
	res, err := common.MakeRequest[PhoneStepResponse](c.HttpClient, "GET", c.BaseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return PhoneStepResponse{}, fmt.Errorf("get phone step details request failed: %s", err)
	}

	return *res, nil
}

// Get ID document step details from sessionID and stepID
func (c *Client) GetIDDocumentStep(sessionID string, stepID string) (IDStepResponse, error) {
	res, err := common.MakeRequest[IDStepResponse](c.HttpClient, "GET", c.BaseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return IDStepResponse{}, fmt.Errorf("get id document step details request failed: %s", err)
	}

	return *res, nil
}

// Get email step details from sessionID and stepID
func (c *Client) GetEmailStep(sessionID string, stepID string) (EmailStepResponse, error) {
	res, err := common.MakeRequest[EmailStepResponse](c.HttpClient, "GET", c.BaseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return EmailStepResponse{}, fmt.Errorf("get email step details request failed: %s", err)
	}

	return *res, nil
}

// Get proof of address step details from sessionID and stepID
func (c *Client) GetProofOfAddressStep(sessionID string, stepID string) (ProofOfAddressStepResponse, error) {
	res, err := common.MakeRequest[ProofOfAddressStepResponse](c.HttpClient, "GET", c.BaseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return ProofOfAddressStepResponse{}, fmt.Errorf("get proof of address step details request failed: %s", err)
	}

	return *res, nil
}

// Get AML step details from sessionID and stepID
func (c *Client) GetAMLStep(sessionID string, stepID string) (AMLStepResponse, error) {
	res, err := common.MakeRequest[AMLStepResponse](c.HttpClient, "GET", c.BaseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return AMLStepResponse{}, fmt.Errorf("get AML step details request failed: %s", err)
	}

	return *res, nil
}
