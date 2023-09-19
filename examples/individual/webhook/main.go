package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/synaps-hub/synaps-sdk-go/pkg/individual"
)

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

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Unmarshaling body
	var payload individual.WebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Error unmarshaling request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if r.URL.Query().Get("secret") != os.Getenv("SYNAPS_WEBHOOK_SECRET") {
		log.Printf("Error wrong webhook secret")
		http.Error(w, "Error invalid secret", http.StatusUnauthorized)
		return
	}

	// Handling event
	handleEvent(payload)

	w.WriteHeader(http.StatusOK)
}

func main() {
	_, ok := os.LookupEnv("SYNAPS_WEBHOOK_SECRET")
	if !ok {
		log.Fatalf("Error missing webhook secret")
	}

	http.HandleFunc("/webhook", handleWebhook)
	fmt.Println("Webhook server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
