package individual

import (
	"fmt"

	"github.com/synaps-hub/synaps-sdk-go/pkg/common"
)

// Get session details from sessionID
func (c *Client) GetSessionDetails(sessionID string) (SessionDetailsResponse, error) {
	res, err := common.MakeRequest[SessionDetailsResponse](c.HttpClient, "GET", c.BaseURL+"/individual/session/"+sessionID, nil, map[string]string{"Api-Key": c.ApiKey})
	if err != nil {
		return SessionDetailsResponse{}, fmt.Errorf("session details request failed: %s", err)
	}
	return *res, nil
}
