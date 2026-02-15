package skills

import (
	"sync"
	"phane-node/internal/security"
)

type SkillMetadata struct {
	ID          string
	Name        string
	Icon        string
	Description string
	Connected   bool
}

var (
	// In a production build, this list would contain all 50+ integrations
	AvailableSkills = []SkillMetadata{
		{"discord", "Discord", "💬", "Bridge your server to your local AI.", false},
		{"notion", "Notion", "📓", "Manage calendars and databases.", false},
		{"gmail", "Gmail", "✉️", "Automate email responses and filtering.", false},
		{"github", "GitHub", "💻", "Monitor repositories and automate commits.", false},
	}
	mu sync.Mutex
)

func ConnectSkill(id string, token string, masterKey string) error {
	// Encrypt the sensitive token before saving to the local database
	encrypted, err := security.EncryptKey(token, masterKey)
	if err != nil { return err }
	
	// Logic to save 'encrypted' to sqlite (phane_users.db) goes here
	return nil
}
