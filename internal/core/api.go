package core

import (
	"encoding/json"
	"net/http"
)

type TaskRequest struct {
	Prompt string `json:"prompt"`
}

func HandleTaskDeployment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", 405)
		return
	}

	var req TaskRequest
	json.NewDecoder(r.Body).Decode(&req)

	// 1. Send to Orchestrator to break down the task
	go RunAutonomousTask(req.Prompt)

	// 2. Respond to UI immediately
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "monitoring started"})
}
