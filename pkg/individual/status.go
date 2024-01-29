package individual

import "github.com/synaps-io/synaps-go/pkg/common"

type Status string

const (
	StatusRejected             Status = Status(common.StatusRejected)             // Rejected KYC session
	StatusSubmissionRequired   Status = Status(common.StatusSubmissionRequired)   // Submission required to continue KYC
	StatusResubmissionRequired Status = Status(common.StatusResubmissionRequired) // Resubmission required to continue KYC
	StatusPending              Status = Status(common.StatusPending)              // Synaps is processing KYC session
	StatusApproved             Status = Status(common.StatusApproved)             // Approved KYC session
)
