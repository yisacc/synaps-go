package corporate

import "github.com/yisacc/synaps-go/pkg/common"

type Status string

const (
	StatusRejected             Status = Status(common.StatusRejected)             // Rejected KYB session
	StatusSubmissionRequired   Status = Status(common.StatusSubmissionRequired)   // Submission required to continue KYB
	StatusResubmissionRequired Status = Status(common.StatusResubmissionRequired) // Resubmission required to continue KYB
	StatusPending              Status = Status(common.StatusPending)              // Synaps is processing KYB session
	StatusApproved             Status = Status(common.StatusApproved)             // Approved KYB session
)
