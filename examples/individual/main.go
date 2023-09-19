package main

import (
	"fmt"
	"log"

	"github.com/synaps-hub/synaps-sdk-go/pkg/individual"
)

func main() {
	client := individual.NewClientFromEnv()

	initSessionRes, err := client.InitSession(individual.InitSessionParams{Alias: "john-doe"})
	if err != nil {
		log.Fatalf("failed to init session: %s", err)
	}
	sessionID := initSessionRes.SessionID

	// Getting session details
	details, err := client.GetSessionDetails(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
	}

	fmt.Printf("session status: %s\n", details.Session.Status)

	processLiveness(client, details)

	processID(client, details)

	processSteps(client, details)
}

// Getting liveness step details (with multiple liveness step)
func processLiveness(client *individual.Client, details individual.SessionDetailsResponse) {
	sessionID := details.Session.ID

	for _, step := range details.Session.Steps {
		if step.Type == individual.LivenessStep {
			livenessStepDetails, err := client.GetLivenessStep(sessionID, step.ID)
			if err != nil {
				log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", step.Type, sessionID, err)
			}

			fmt.Printf("Liveness step status: %s\n", step.Status)

			switch step.Status {
			case individual.StatusApproved:
				fmt.Printf("Liveness file url: %s\n", livenessStepDetails.Verification.Liveness.File.URL)
			case individual.StatusRejected:
				fmt.Printf("Liveness reject reason: %s\n", livenessStepDetails.Reason.Message)
			default:
				fmt.Printf("Liveness step is not finished yet\n")
			}
		}
	}
}

// Getting id document step details (with only one ID step)
func processID(client *individual.Client, details individual.SessionDetailsResponse) {
	var IDStep *individual.Step
	for _, step := range details.Session.Steps {
		if step.Type == individual.IDDocumentStep {
			IDStep = &step
			break
		}
	}

	sessionID := details.Session.ID
	if IDStep == nil {
		log.Fatalf("failed to get step for session[%s]", sessionID)
	}

	IDStepDetails, err := client.GetIDDocumentStep(sessionID, IDStep.ID)
	if err != nil {
		log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", IDStep.Type, sessionID, err)
	}

	fmt.Printf("ID step status: %s\n", IDStepDetails.Status)

	if IDStepDetails.Status == individual.StatusApproved {
		fmt.Printf("ID Document firstname: %s\n", IDStepDetails.Document.Fields["firstname"])
	}

	if IDStepDetails.Status == individual.StatusRejected {
		fmt.Printf("ID Document rejected: %s\n", IDStepDetails.Reason.Message)
	}
}

// Iterating over steps
func processSteps(client *individual.Client, details individual.SessionDetailsResponse) {
	sessionID := details.Session.ID

	var response any
	var err error
	for _, step := range details.Session.Steps {
		switch step.Type {
		case individual.LivenessStep:
			response, err = client.GetLivenessStep(sessionID, step.ID)
		case individual.IDDocumentStep:
			response, err = client.GetIDDocumentStep(sessionID, step.ID)
		case individual.EmailStep:
			response, err = client.GetEmailStep(sessionID, step.ID)
		case individual.PhoneStep:
			response, err = client.GetPhoneStep(sessionID, step.ID)
		case individual.ProofOfAddressStep:
			response, err = client.GetProofOfAddressStep(sessionID, step.ID)
		case individual.AMLStep:
			response, err = client.GetAMLStep(sessionID, step.ID)
		}

		if err != nil {
			log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", step.Type, sessionID, err)
			continue
		}
		fmt.Printf("Response is:\n%+v\n", response)
	}
}
