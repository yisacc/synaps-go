package corporate

import (
	"fmt"

	"github.com/synaps-hub/synaps-sdk-go/pkg/common"
)

// Get documents step from sessionID
func (c *Client) GetDocumentsStep(sessionID string) (DocumentsStepResponse, error) {
	res, err := common.MakeRequest[DocumentsStepResponse](c.HttpClient, "GET", c.BaseURL+"/corporate/session/"+sessionID+"/step/"+string(DocumentsStep), nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return DocumentsStepResponse{}, fmt.Errorf("get documents step details request failed: %s", err)
	}

	return *res, nil
}

// Get company details step from sessionID
func (c *Client) GetCompanyDetailsStep(sessionID string) (CompanyDetailsStepResponse, error) {
	res, err := common.MakeRequest[CompanyDetailsStepResponse](c.HttpClient, "GET", c.BaseURL+"/corporate/session/"+sessionID+"/step/"+string(CompanyDetailsStep), nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return CompanyDetailsStepResponse{}, fmt.Errorf("get company details step details request failed: %s", err)
	}

	return *res, nil
}

// Get beneficial owner step from sessionID
func (c *Client) GetBeneficialOwnersStep(sessionID string) (BeneficialOwnersStepResponse, error) {
	res, err := common.MakeRequest[BeneficialOwnersStepResponse](c.HttpClient, "GET", c.BaseURL+"/corporate/session/"+sessionID+"/step/"+string(BeneficialOwnersStep), nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return BeneficialOwnersStepResponse{}, fmt.Errorf("get beneficial owners step details request failed: %s", err)
	}

	return *res, nil
}

// Get legal representative step from sessionID
func (c *Client) GetLegalRepresentativeStep(sessionID string) (LegalRepresentativeStepResponse, error) {
	res, err := common.MakeRequest[LegalRepresentativeStepResponse](c.HttpClient, "GET", c.BaseURL+"/corporate/session/"+sessionID+"/step/"+string(LegalRepresentativeStep), nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return LegalRepresentativeStepResponse{}, fmt.Errorf("get legal representative step details request failed: %s", err)
	}

	return *res, nil
}

// Get AML step from sessionID
func (c *Client) GetAMLStep(sessionID string) (AMLStepResponse, error) {
	res, err := common.MakeRequest[AMLStepResponse](c.HttpClient, "GET", c.BaseURL+"/corporate/session/"+sessionID+"/step/"+string(AMLStep), nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return AMLStepResponse{}, fmt.Errorf("get AML step details request failed: %s", err)
	}

	return *res, nil
}
