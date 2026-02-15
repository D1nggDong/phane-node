package network

import (
	"encoding/json"
	"net/http"
	"time"
)

type LicenseStatus struct {
	IsPro     bool      `json:"is_pro"`
	ExpiresAt time.Time `json:"expires_at"`
}

// CheckSubscription pings your central server to see if this Node ID has paid
func CheckSubscription(nodeID string) (bool, error) {
	// In production, this pings https://api.phane.ai/verify
	resp, err := http.Get("https://api.phane.ai/verify?node_id=" + nodeID)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var status LicenseStatus
	json.NewDecoder(resp.Body).Decode(&status)

	return status.IsPro, nil
}
