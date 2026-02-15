package skills

import (
	"bytes"
	"encoding/json"
	"net/http"
	"phane-node/internal/security"
)

type DiscordSkill struct{}

func (d DiscordSkill) Name() string { return "Discord Bridge" }
func (d DiscordSkill) Description() string { return "Post alerts to Discord channels." }

func (d DiscordSkill) Execute(encryptedKey string, channelID string, message string) error {
	master := security.GetSessionKey()
	apiKey, _ := security.DecryptKey(encryptedKey, master)

	payload, _ := json.Marshal(map[string]string{"content": message})
	url := "https://discord.com/api/v9/channels/" + channelID + "/messages"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Authorization", "Bot "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	return nil
}
