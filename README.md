# Synaps Go SDK

> [!WARNING]
> This SDK is currently in Beta.

The Synaps Go SDK is divided in two packages [individual](#Individual) and [corporate](#Corporate).
Both provides a convenient way to interact with the Synaps API, individual for **Know your customer** sessions (KYC), and corporate for **Know your business** sessions (KYB).

# Requirements

Before you start using this SDK, ensure that you have the following:

- **Go Programming Language**: version 1.18 or higher.

- **Synaps API Key**: Your Synaps API key. You can find it on the [manager app](https://manager-kyc.synaps.io) within the developer section of your app.

# Individual

Individual SDK enables you to initiate sessions, retrieve session details, and obtain information about different steps within a session (Liveness, Identity, Proof of address, etc.).

> For more details, please refer to the Synaps API documentation at [https://docs.synaps.io](https://docs.synaps.io).

## Installation

To use the Synaps Individual Go SDK, you can add it as a dependency in your project using `go get`:

```bash
go get github.com/synaps-io/synaps-go/pkg/individual
```

## Usage

The SDK facilitates the initiation of sessions, tracking user KYC progress, retrieving verification results, and handling events through webhooks.
This section provides an overview of the fundamental steps to integrate the SDK into your project and start leveraging its functionalities.

> A complete example can be found in the [examples/individual/main.go](https://github.com/synaps-io/synaps-go/blob/main/examples/individual/main.go) file within the repository.

### Import

```go
import (
	"github.com/synaps-io/synaps-go/pkg/individual"
)
```

### Configuring client

Set the `SYNAPS_API_KEY` env variable to your api key value and create a new Synaps client from environment: 

```go
client := individual.NewClientFromEnv()
```
> This will also check for `.env` file

Or create it from variables:
```go
client := individual.NewClient("$YOUR_API_KEY")
```

### Initialize session

You can choose when to create a KYC session for your user (on register, on first withdrawal, on first deposit, ...), the most common way is on register.

Initialize a new session:

```go
// ... Create user

initSessionRes, err := client.InitSession(individual.InitSessionParams{})
if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID

// ... Store user with sessionID
```

Initialize a new session with an `alias` (most common use case is to set username as an alias):

```go
// ... Create user

initSessionRes, err := client.InitSession(individual.InitSessionParams{Alias: "john-doe"})
if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID

// ... Store user with sessionID
```

Initialize a new session with `metadata`:

```go
// ... Create user

initSessionRes, err := client.InitSession(individual.InitSessionParams{Metadata: map[string]string{
    "firstname": "John",
    "lastname": "Doe",
    "email": "john@doe.io",
}})

if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID

// ... Store user with sessionID
```

### Get session details

(Refer to the [documentation](https://docs.synaps.io/session#get-session-details) for details about the session details response)

While waiting for KYC session to be completed you can use our API to provide informations about their KYC to your users. Here is how to use SDK to get theses informations:

```go
// ... Get user and associated sessionID 

details, err := client.GetSessionDetails(sessionID)
if err != nil {
	log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
}

fmt.Printf("session status: %s\n", details.Session.Status)

// ... Send back to users 
```

### Get step details
(Refer to the [documentation](https://docs.synaps.io/steps#get-step-details) for details about the step details response)

Steps in your **KYC flow** (Liveness, ID document, Proof of residency, ...), are configurable, thus you can have multiple steps of the same type:

```go
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
```

An exemple for getting ID step details (with only one step ID Document):
```go
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

if IDStepDetails.Status == individual.StatusPending || IDStepDetails.Status == individual.StatusApproved {
	fmt.Printf("ID Document firstname: %s\n", IDStepDetails.Document.Fields["firstname"])
}
```

Iterating through steps:
```go
...
var response any
var err error

for _, step := range details.Session.Steps {
	switch step.Type {
	case individual.LivenessStep:
		response, err = client.GetLivenessStep(sessionID, step.ID)
        // Do your stuff...
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
    // Do your stuff...
}
```
## API Reference

For more details about the API, please refer to the [Synaps API Reference](https://docs.synaps.io/session).

# Corporate

Corporate SDK enables you to initiate sessions, retrieve session details, and obtain information about different steps within a session (Company details, Legal representative, Beneficial owners, etc.).

> For more details, please refer to the Synaps API documentation at [https://docs.synaps.io](https://docs.synaps.io).

## Installation

To use the Synaps corporate Go SDK, you can add it as a dependency in your project using `go get`:

```bash
go get github.com/synaps-io/synaps-go/pkg/corporate
```

## Usage

The SDK facilitates the initiation of sessions, tracking user KYC progress, retrieving verification results, and handling events through **webhooks**.
This section provides an overview of the fundamental steps to integrate the SDK into your project and start leveraging its functionalities.

> A complete example can be found in the [examples/corporate/main.go](https://github.com/synaps-io/synaps-go/blob/main/examples/corporate/main.go) file within the repository.

### Import

```go
import (
	"github.com/synaps-io/synaps-go/pkg/corporate"
)
```

### Configuring client

Set the `SYNAPS_API_KEY` env variable to your api key value and create a new Synaps client from environment: 

```go
client := corporate.NewClientFromEnv()
```
> This will also check for `.env` file

Or create it from variables:
```go
client := corporate.NewClient("$YOUR_API_KEY")
```

### Initialize session

You can choose when to create a KYB session for the user (on register, on first withdrawal, on first deposit, ...), the most common way is on register.

Initialize a new session:

```go
// ... Create user

initSessionRes, err := client.InitSession(corporate.InitSessionParams{})
if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID

// ... Store user with sessionID
```

Initialize a new session with an `alias` (most common use case is to set username as an alias):

```go
// ... Create user

initSessionRes, err := client.InitSession(corporate.InitSessionParams{Alias: "john-doe"})
if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID

// ... Store user with sessionID
```

Initialize a new session with `metadata`:

```go
// ... Create user

initSessionRes, err := client.InitSession(corporate.InitSessionParams{Metadata: map[string]string{
    "firstname": "John",
    "lastname": "Doe",
    "email": "john@doe.io",
}})

if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID

// ... Store user with sessionID
```

### Get session details

(Refer to the [documentation](https://docs.synaps.io/session#get-session-details) for details about the session details response)

While waiting for KYB session to be completed you can use our API to provide informations about their KYB to your users. Here is how to use SDK to get theses informations:

```go
// ... Get user and associated sessionID 

details, err := client.GetSessionDetails(sessionID)
if err != nil {
	log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
}

fmt.Printf("session status: %s\n", details.Session.Status)

```

### Get step details
(Refer to the [documentation](https://docs.synaps.io/steps#get-step-details) for details about the step details response)

Getting company details step:

```go
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
```

An exemple for getting documents step details:
```go
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
fmt.Printf("Shareholder register status: %s\n", DocumentsStepDetails.RegisterOfDirectorsShareholders.Status)

if DocumentsStepDetails.CertificateOfIncorporation.Status == corporate.StatusApproved {
	for _, file := range DocumentsStepDetails.CertificateOfIncorporation.Files {
		fmt.Printf("Certificate of incorporation document: %s\n", file.URL)
	}
}
```

Iterating through steps:
```go
...
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
```

## API Reference

For more details about the API, please refer to the [Synaps API Reference](https://docs.synaps.io/session).

# Webhooks

Whether you using corporate or individual SDK, we provide webhooks, which are triggered for any state change in a session, so you can update your database and notify your users as needed.

In order to receive webhooks, you'll need to create an endpoint that can receive and handle the webhook events. Below is an example, for individual SDK, of how to set up the necessary components.

> You can find the complete example in the [examples/individual/webhook/main.go](https://github.com/synaps-io/synaps-go/blob/main/examples/individual/webhook/main.go) or [examples/corporate/webhook/main.go](https://github.com/synaps-io/synaps-go/blob/main/examples/corporate/webhook/main.go) file within the repository.

## Import
```go
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/synaps-io/synaps-go/pkg/individual" // Or "github.com/synaps-io/synaps-go/pkg/corporate"
)
```

## Handle webhook
Create your handler function for processing incoming requests:
```go
func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

        // Unmarshaling body
	var payload individual.WebhookPayload // Or corporate.WebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Error unmarshaling request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

        // Checking for secret
	if r.URL.Query().Get("secret") != os.Getenv("SYNAPS_WEBHOOK_SECRET") {
		log.Printf("Error wrong webhook secret")
		http.Error(w, "Error invalid secret", http.StatusUnauthorized)
		return
	}

	handleEvent(payload)

	w.WriteHeader(http.StatusOK)
}
```

## Handle event

Create a function to handle the webhook event:

```go
func handleEvent(payload individual.WebhookPayload) {
	switch payload.Service {
	case individual.IDDocumentStep:
		log.Printf("Received ID document event: %s: %s", payload.Status, payload.Reason)

        
		// ... Find user by sessionID and send email to user

	case individual.LivenessStep:
		log.Printf("Received liveness event: %s: %s", payload.Status, payload.Reason)

		// ... Find user by sessionID and send email to user
	}
}
```

## Serve endpoint

```go
func main() {
	_, ok := os.LookupEnv("SYNAPS_WEBHOOK_SECRET")
	if !ok {
		log.Fatalf("Error missing webhook secret")
	}

	http.HandleFunc("/webhook", handleWebhook)
	fmt.Println("Webhook server listening on port 80...")
	http.ListenAndServe(":80", nil)
}
```

> Ensure that your endpoint is reachable from the internet so webhook server can reach it

Once done, add your endpoint URL to Synaps [manager](https://manager-kyc.synaps.io) (see [documentation](https://docs.synaps.io/quickstart#6-configure-webhooks) for guidance).

Congratulations, you're now all set!

Be sure not to overlook theses steps to ensure security:
- Verifying that the secret in the query parameters is matching the one given to you on the manager. This step ensures that you are exclusively receiving events from Synaps, as shown in the [individual example](https://github.com/synaps-io/synaps-go/blob/main/examples/individual/webhook/main.go#L37) or [corporate example](https://github.com/synaps-io/synaps-go/blob/main/examples/corporate/webhook/main.go#L37).
- Utilizing HTTPS to establish a secure communication channel. This practice ensures the confidentiality and integrity of the data being exchanged.


# License

This SDK is released under the [MIT License](LICENSE).
