package individual

type StepType string

// Every individual step type
const (
	LivenessStep       StepType = "LIVENESS"
	IDDocumentStep     StepType = "ID_DOCUMENT"
	ProofOfAddressStep StepType = "PROOF_OF_ADDRESS"
	EmailStep          StepType = "EMAIL"
	PhoneStep          StepType = "PHONE"
	AMLStep            StepType = "AML"
)

// Rejection code associated with reason
type ReasonCode string

// Every individual rejection reason
const (
	ForgedRejectionReason              ReasonCode = "FORGED_REJECTION"
	DocumentHiddenReason               ReasonCode = "DOCUMENT_HIDDEN"
	BadEnvironmentReason               ReasonCode = "BAD_ENVIRONMENT"
	BlackWhitePictureReason            ReasonCode = "BLACK_WHITE_PICTURE"
	BadQualityReason                   ReasonCode = "BAD_QUALITY"
	DocumentComplianceReason           ReasonCode = "DOCUMENT_COMPLIANCE"
	IdentityDocumentExpiredReason      ReasonCode = "IDENTITY_DOCUMENT_EXPIRED"
	DocumentInvalidFrontSideReason     ReasonCode = "DOCUMENT_INVALID_FRONT_SIDE"
	DocumentInvalidBackSideReason      ReasonCode = "DOCUMENT_INVALID_BACK_SIDE"
	IdentityDocumentDobDateMinorReason ReasonCode = "IDENTITY_DOCUMENT_DOB_DATE_MINOR"
	RestrictedNationalityTypeReason    ReasonCode = "RESTRICTED_NATIONALITY_TYPE"
)

// Step rejection reason
type StepReason struct {
	Code    ReasonCode
	Message string
}

type IDDocumentType string

// Every document type for ID Document step
const (
	Passport       IDDocumentType = "PASSPORT"
	NationalID     IDDocumentType = "NATIONAL_ID"
	DriverLicense  IDDocumentType = "DRIVER_LICENSE"
	ResidentPermit IDDocumentType = "RESIDENT_PERMIT"
)

type ProofOfAddressDocumentType string

// Every document type for proof of address step
const (
	ElectricityBill ProofOfAddressDocumentType = "ELECTRICITY_BILL"
	InternetBill    ProofOfAddressDocumentType = "INTERNET_BILL"
	LandlineBill    ProofOfAddressDocumentType = "LANDLINE_BILL"
	WaterBill       ProofOfAddressDocumentType = "WATER_BILL"
	GasBill         ProofOfAddressDocumentType = "GAS_BILL"
	BankStatement   ProofOfAddressDocumentType = "BANK_STATEMENT"
)

type PhoneMethod string

// Every phone verification method
const (
	Sms  PhoneMethod = "sms"
	Call PhoneMethod = "call"
)

type File struct {
	URL  string `json:"url"`
	Type string `json:"type"`
	Size int    `json:"size"`
}

type (
	// ID Document step relative data
	idDocumentData struct {
		Country string            `json:"country"`
		Type    IDDocumentType    `json:"type"`
		Fields  map[string]string `json:"fields"`
		Files   struct {
			Front File `json:"front"`
			Back  File `json:"back"`
			Face  File `json:"face"`
		} `json:"files"`
		OriginalFiles struct {
			Front File `json:"front"`
			Back  File `json:"back"`
		} `json:"original_files"`
	}
	// Proof of address step relative data
	proofOfAddressData struct {
		Country string                     `json:"country"`
		Type    ProofOfAddressDocumentType `json:"type"`
		Fields  map[string]string          `json:"fields"`
		Files   struct {
			Accomodation File `json:"accomodation"`
			Proof        File `json:"proof"`
		} `json:"files"`
		HostIDDocument idDocumentData `json:"host_id_document"`
		OriginalFiles  struct {
			Accomodation File `json:"accomodation"`
			Proof        File `json:"proof"`
		} `json:"original_files"`
	}
	// Phone step relative data
	phoneData struct {
		Phone struct {
			CallingCode string      `json:"calling_code"`
			Country     string      `json:"country"`
			Method      PhoneMethod `json:"method"`
			Number      string      `json:"number"`
		} `json:"phone"`
	}
	// Email step relative data
	emailData struct {
		Email struct {
			Value string `json:"value"`
		} `json:"email"`
	}
	// Liveness step relative data
	livenessData struct {
		Liveness struct {
			File File `json:"file"`
		} `json:"liveness"`
	}
	// AML step relative data
	amlData struct {
		AML struct {
			CustomerInfo struct {
				Firstname   string `json:"firstname"`
				Lastname    string `json:"lastname"`
				BirthDate   string `json:"birth_date"`
				Nationality string `json:"nationality"`
				Country     string `json:"country"`
			} `json:"customer_info"`
			Hits []struct {
				Fields []struct {
					Countries []struct {
						CountryCode string `json:"country_code"`
						Name        string `json:"name"`
					} `json:"countries"`
					Name   string `json:"name"`
					Source string `json:"source"`
					Value  []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"value"`
					Tag []string `json:"tag"`
				} `json:"fields"`
				Aliases []string `json:"aliases"`
				Media   []struct {
					Date    string `json:"date"`
					Snippet string `json:"snippet"`
					Title   string `json:"title"`
					URL     string `json:"url"`
				} `json:"media"`
				Info struct {
					FullName    string `json:"full_name"`
					BirthDate   string `json:"birth_date"`
					Nationality string `json:"nationality"`
				} `json:"info"`
				Type []string `json:"type"`
			} `json:"hits"`
			TotalHits      int  `json:"total_hits"`
			ActionRequired bool `json:"action_required"`
			Monitored      bool `json:"monitored"`
		} `json:"aml"`
	}
)

type StepMetadata struct {
	IP               string `json:"ip"`
	UserAgent        string `json:"user_agent"`
	Platform         string `json:"platform"`
	BrowserName      string `json:"browser_name"`
	BrowserVersion   string `json:"browser_version"`
	Device           string `json:"device"`
	CustomerLanguage string `json:"customer_language"`
}

// Phone step response
type PhoneStepResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification phoneData `json:"verification"`
}

// Email step response
type EmailStepResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification emailData `json:"verification"`
}

// Proof of address step response
type ProofOfAddressStepResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Document proofOfAddressData `json:"document"`
}

// ID document step response
type IDStepResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Document idDocumentData `json:"document"`
}

// Liveness step response
type LivenessStepResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification livenessData `json:"verification"`
}

// AML step response
type AMLStepResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification amlData `json:"verification"`
}
