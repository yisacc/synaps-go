package corporate

// Webhook payload you receive in case of event
type WebhookPayload struct {
	Reason    string       `json:"reason"`
	Service   StepType     `json:"service"`
	SessionID string       `json:"session_id"`
	Status    WebhookEvent `json:"status"`
	StepID    string       `json:"step_id"`
}

type WebhookEvent string

// Every corporate webhook event type
// ref: [common.Status]
const (
	EventRejected             WebhookEvent = "REJECTED"              // KYB session state changed to rejected
	EventSubmissionRequired   WebhookEvent = "SUBMISSION_REQUIRED"   // KYB session state changed to submission required
	EventResubmissionRequired WebhookEvent = "RESUBMISSION_REQUIRED" // KYB session state changed to resubmission required
	EventPending              WebhookEvent = "PENDING_VERIFICATION"  // KYB session state changed to pending verification
	EventApproved             WebhookEvent = "APPROVED"              // KYB session state changed to approved
	EventReset                WebhookEvent = "RESET"                 // KYB session got reset
)
