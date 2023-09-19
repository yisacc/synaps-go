package individual

// Webhook payload you receive in case of event
type WebhookPayload struct {
	Reason    string       `json:"reason"`
	Service   StepType     `json:"service"`
	SessionID string       `json:"session_id"`
	Status    WebhookEvent `json:"status"`
	StepID    string       `json:"step_id"`
}

type WebhookEvent string

// Every individual webhook event type
//
// see Status: [common.Status]
const (
	EventRejected 		  WebhookEvent = "REJECTED"              // KYC session state changed to rejected
	EventSubmissionRequired   WebhookEvent = "SUBMISSION_REQUIRED"   // KYC session state changed to submission required
	EventResubmissionRequired WebhookEvent = "RESUBMISSION_REQUIRED" // KYC session state changed to resubmission required
	EventPending              WebhookEvent = "PENDING_VERIFICATION"  // KYC session state changed to pending verification
	EventApproved             WebhookEvent = "APPROVED"              // KYC session state changed to approved
	EventReset                WebhookEvent = "RESET"                 // KYC session got reset
)
