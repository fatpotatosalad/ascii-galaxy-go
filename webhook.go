package main

/*import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type GiteaPushPayload struct {
	Pusher     GiteaUser    `json: "pusher"`
	After      string       `json: "after"`
	Ref        string       `json: "ref"`
	Repository GiteaRepo    `json: "repository"`
	HeadCommit *GiteaCommit `json: "head_commit"`
}

type GiteaUser struct {
	Username string `json:"username"`
	Login    string `json:"login"`
}

type GiteaRepo struct {
	Name string `json:"name"`
}

type GiteaCommit struct {
	Timestamp string `json:"timestamp"`
}

type GitlabPushPayload struct {
	Pusher     GitlabUser   `json: "pusher"`
	After      string       `json: "after"`
	Ref        string       `json: "ref"`
	Repository GiteaRepo    `json: "repository"`
	HeadCommit *GiteaCommit `json: "head_commit"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}

	var payload WebhookPayload

	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("Invalid JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// handle webhook event

	log.Printf("webhook received:")
	log.Printf("Event: %s", payload.Event)
	log.Printf("ID: %s", payload.ID)
	log.Printf("Data: %+v", payload.Data)

	response := map[string]interface{}{
		"success": true,
		"message": "webhook received successfully!",
		"id":      payload.ID,
	}

	w.Header().Set(
		"Content-Type", "application/json",
	)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to write response: %v", err)
	}

}
*/
