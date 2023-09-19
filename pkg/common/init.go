package common

// Base init session response
type InitSessionResponse struct {
	SessionID string `json:"session_id"`
	Sandbox   bool   `json:"sandbox"`
}

// Base init session params
type InitSessionParams struct {
	Metadata map[string]string `json:"metadata,omitempty"`
	Alias    string            `json:"alias,omitempty"`
}
