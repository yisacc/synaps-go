package common

// Response in case of error from the Synaps API
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
