package main

import (
	"fmt"
	"log"

	"github.com/synaps-io/synaps-go/pkg/corporate"
)

func main() {
	client := corporate.NewClientFromEnv()

	initSessionRes, err := client.InitSession(corporate.InitSessionParams{Alias: "john-doe"})
	if err != nil {
		log.Fatalf("failed to init session: %s", err)
	}
	sessionID := initSessionRes.SessionID

	fmt.Printf("session id: %s\n", sessionID)

	// Getting session details
	details, err := client.GetSessionDetails(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
	}

	fmt.Printf("session status: %s\n", details.Session.Status)

	processCompanyDetails(client, details)

	processDocuments(client, details)

	processSteps(client, details)
}

// Getting liveness step details
func processCompanyDetails(client *corporate.Client, details corporate.SessionDetailsResponse) {
	sessionID := details.Session.ID

	for _, step := range details.Session.Steps {
		if step.Type == corporate.CompanyDetailsStep {
			companyStepDetails, err := client.GetCompanyDetailsStep(sessionID)
			if err != nil {
				log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", step.Type, sessionID, err)
			}

			fmt.Printf("Company details step status: %s\n", step.Status)

			switch step.Status {
			case corporate.StatusApproved:
				fmt.Printf("Company details address: %s\n", companyStepDetails.Address)
				fmt.Printf("Company details city: %s\n", companyStepDetails.City)
			case corporate.StatusRejected:
				fmt.Printf("Company details rejected\n")
			default:
				fmt.Printf("Company details step is not finished yet\n")
			}
		}
	}
}

// Getting documents step details
func processDocuments(client *corporate.Client, details corporate.SessionDetailsResponse) {
	var DocumentStep *corporate.Step
	for _, step := range details.Session.Steps {
		if step.Type == corporate.DocumentsStep {
			DocumentStep = &step
			break
		}
	}

	sessionID := details.Session.ID
	if DocumentStep == nil {
		log.Fatalf("failed to get step for session[%s]", sessionID)
	}

	DocumentsStepDetails, err := client.GetDocumentsStep(sessionID)
	if err != nil {
		log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", DocumentStep.Type, sessionID, err)
	}

	fmt.Printf("Certificate of incorporation status: %s\n", DocumentsStepDetails.CertificateOfIncorporation.Status)
	fmt.Printf("Memorandum of association status: %s\n", DocumentsStepDetails.MemorandumOfAssociation.Status)
	fmt.Printf("Shareholder register status: %s\n", DocumentsStepDetails.RegistryOfDirectorsAndShareholders.Status)

	if DocumentsStepDetails.CertificateOfIncorporation.Status == corporate.StatusApproved {
		for _, file := range DocumentsStepDetails.CertificateOfIncorporation.Files {
			fmt.Printf("Certificate of incorporation document: %s\n", file.URL)
		}
	}
}

// Iterating over steps
func processSteps(client *corporate.Client, details corporate.SessionDetailsResponse) {
	sessionID := details.Session.ID

	var response any
	var err error
	for _, step := range details.Session.Steps {
		switch step.Type {
		case corporate.CompanyDetailsStep:
			response, err = client.GetCompanyDetailsStep(sessionID)

		case corporate.DocumentsStep:
			response, err = client.GetDocumentsStep(sessionID)

		case corporate.BeneficialOwnersStep:
			response, err = client.GetBeneficialOwnersStep(sessionID)

		case corporate.LegalRepresentativeStep:
			response, err = client.GetLegalRepresentativeStep(sessionID)
		case corporate.AMLStep:
			response, err = client.GetAMLStep(sessionID)
		}

		if err != nil {
			log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", step.Type, sessionID, err)
			continue
		}
		fmt.Printf("Response is:\n%+v\n", response)
	}
}
