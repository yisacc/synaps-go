package corporate

type Step struct {
	Status Status   `json:"status"`
	Type   StepType `json:"type"`
}

// Session details response
type SessionDetailsResponse struct {
	App struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"app"`
	Session struct {
		ID      string `json:"id"`
		Alias   string `json:"alias"`
		Status  Status `json:"status"`
		Sandbox bool   `json:"sandbox"`
		Steps   []Step `json:"steps"`
	} `json:"session"`
}
