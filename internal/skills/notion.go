package skills

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"phane-node/internal/security"
)

type NotionSkill struct{}

func (n NotionSkill) Name() string { return "Notion Sync" }
func (n NotionSkill) Description() string { return "Update Notion databases autonomously." }

func (n NotionSkill) Execute(encryptedKey string) error {
	// 1. Decrypt the key using the Master Key in RAM
	master := security.GetSessionKey()
	apiKey, err := security.DecryptKey(encryptedKey, master)
	if err != nil {
		return fmt.Errorf("vault decryption failed: %v", err)
	}

	// 2. Simple Notion API Handshake (Conceptual)
	// This keeps the 1GB RAM footprint tiny by using native http
	req, _ := http.NewRequest("GET", "https://api.notion.com/v1/users/me", nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Notion-Version", "2022-06-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("notion connection failed")
	}
	defer resp.Body.Close()

	return nil
}
