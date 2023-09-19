package corporate

type StepType string

// Every corporate step type
const (
	AMLStep                 StepType = "AML"
	DocumentsStep           StepType = "DOCUMENTS"
	CompanyDetailsStep      StepType = "COMPANY_DETAILS"
	BeneficialOwnersStep    StepType = "BENEFICIAL_OWNERS"
	LegalRepresentativeStep StepType = "LEGAL_REPRESENTATIVE"
)

type File struct {
	URL  string `json:"url"`
	Type string `json:"type"`
	Size int    `json:"size"`
}

type Document struct {
	Files  []File `json:"files"`
	Status Status `json:"status"`
}

// Company details step response
type CompanyDetailsStepResponse struct {
	Name               string `json:"name"`
	City               string `json:"city"`
	Country            string `json:"country"`
	Address            string `json:"address"`
	Zipcode            string `json:"zipcode"`
	RegistrationNumber string `json:"registration_number"`
	Structure          string `json:"structure"`
}

// Documents step response
type DocumentsStepResponse struct {
	CertificateOfIncorporation         Document `json:"CERTIFICATE_OF_INCORPORATION"`
	MemorandumOfAssociation            Document `json:"MEMORANDUM_OF_ASSOCIATION"`
	RegistryOfDirectorsAndShareholders Document `json:"REGISTRY_OF_DIRECTORS_AND_SHAREHOLDERS"`
}

// Legal representative step response
type LegalRepresentativeStepResponse struct {
	ID            string        `json:"id"`
	FirstName     string        `json:"firstname"`
	LastName      string        `json:"lastname"`
	BirthDate     string        `json:"birth_date"`
	Nationality   string        `json:"nationality"`
	Authorization Authorization `json:"authorization"`
	Steps         []Step        `json:"steps"`
	Status        Status        `json:"status"`
}

type Authorization struct {
	Required bool `json:"required"`
	File     File `json:"file"`
}

// Beneficial owner step response
type BeneficialOwnersStepResponse struct {
	Corporates  []Corporate  `json:"corporates"`
	Individuals []Individual `json:"individuals"`
}

type Edge struct {
	NodeID   int     `json:"node_id"`
	Shares   float64 `json:"shares"`
	Director bool    `json:"director"`
}

type Corporate struct {
	NodeID         int                        `json:"node_id"`
	CompanyDetails CompanyDetailsStepResponse `json:"company_details"`
	AML            AMLStepResponse            `json:"aml"`
	Documents      map[string]Document        `json:"documents"`
	Edges          []Edge                     `json:"edges"`
}

type Individual struct {
	NodeID      int    `json:"node_id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	BirthDate   string `json:"birth_date"`
	Nationality string `json:"nationality"`
	Ultimate    bool   `json:"ultimate"`
	Links       []Edge `json:"links"`
	ID          string `json:"id"`
	Steps       []Step `json:"steps"`
	Status      Status `json:"status"`
	Edges       []Edge `json:"edges"`
}

// AML step response
type AMLStepResponse struct {
	ActionRequired bool `json:"action_required"`
	Matches        int  `json:"matches"`
	Screening      []struct {
		ID   string `json:"id"`
		Info struct {
			FullName   string   `json:"full_name"`
			EntityType string   `json:"entity_type"`
			Aka        []string `json:"aka"`
			Countries  string   `json:"countries"`
			BirthDate  string   `json:"birth_date"`
			Associates []struct {
				Association string `json:"association"`
				Name        string `json:"name"`
			} `json:"associates"`
		} `json:"info"`
		Sanctions    []AmlScreeningListing `json:"sanctions"`
		PEP          []AmlScreeningListing `json:"pep"`
		AdverseMedia []struct {
			Date    string `json:"date"`
			Snippet string `json:"snippet"`
			Title   string `json:"title"`
			URL     string `json:"url"`
		} `json:"adverse_media"`
		Types []string `json:"types"`
	} `json:"screening"`
	Status   Status `json:"status"`
	Decision struct {
		ComplianceOfficer string `json:"compliance_officer"`
		UserID            string `json:"user_id"`
		Output            string `json:"output"`
		InternalNote      string `json:"internal_note"`
		Reasons           []struct {
			Reason string `json:"reason"`
			Note   string `json:"note"`
		} `json:"reasons"`
	} `json:"decision"`
	Monitoring bool `json:"monitoring"`
}

type AmlScreeningListing struct {
	Types     []string `json:"types"`
	Countries []struct {
		CountryCode string `json:"country_code"`
		Name        string `json:"name"`
	} `json:"countries"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	ListingEnded   string `json:"listing_ended"`
	ListingStarted string `json:"listing_started"`
	Fields         []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"fields"`
}
