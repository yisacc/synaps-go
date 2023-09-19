package common

type Status string

const (
	StatusRejected             Status = "REJECTED"              // Rejected KYC session
	StatusSubmissionRequired   Status = "SUBMISSION_REQUIRED"   // Submission required to continue KYC
	StatusResubmissionRequired Status = "RESUBMISSION_REQUIRED" // Resubmission required to continue KYC
	StatusPending              Status = "PENDING_VERIFICATION"  // Synaps is processing KYC session
	StatusApproved             Status = "APPROVED"              // Approved KYC session
)
