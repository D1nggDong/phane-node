package security

import (
	"encoding/json"
	"net/http"
	"time"
)

type LicenseCheckResponse struct {
	Valid     bool      `json:"valid"`
	Features  []string  `json:"features"`
	ExpiresAt time.Time `json:"expires_at"`
}

// VerifyProStatus checks with phane.ai if the license is active
func VerifyProStatus(licenseKey string) (bool, error) {
	nodeID := GetNodeID()
	url := "https://api.phane.ai/v1/verify?node_id=" + nodeID + "&key=" + licenseKey

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result LicenseCheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Valid, nil
}
