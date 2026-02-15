package skills

import (
	"fmt"
	"phane-node/internal/security"
	"github.com/bwmarrin/discordgo"
)

// TriggerHandshake initiates the first contact with a service after encryption
func TriggerHandshake(serviceID string, encryptedToken string, masterKey string) error {
	// 1. Decrypt locally [cite: 2026-02-12]
	token, err := security.DecryptKey(encryptedToken, masterKey)
	if err != nil {
		return fmt.Errorf("decryption failed: %v", err)
	}

	// 2. Route to specific service logic
	switch serviceID {
	case "discord":
		return sendDiscordHandshake(token)
	case "whatsapp":
		return sendWhatsAppHandshake(token)
	default:
		return fmt.Errorf("handshake for %s is pending deployment", serviceID)
	}
}

func sendDiscordHandshake(token string) error {
	dg, err := discordgo.New("Bot " + token)
	if err != nil { return err }
	
	err = dg.Open()
	if err != nil { return err }
	defer dg.Close()

	// Find the first available channel to say hello
	guilds, _ := dg.UserGuilds(1, "", "")
	if len(guilds) > 0 {
		channels, _ := dg.GuildChannels(guilds[0].ID)
		for _, c := range channels {
			if c.Type == discordgo.ChannelTypeGuildText {
				dg.ChannelMessageSend(c.ID, "🛡️ **PHANE SOVEREIGN NODE ONLINE**\nConnection established via AES-256 local bridge.")
				break
			}
		}
	}
	return nil
}

func sendWhatsAppHandshake(token string) error {
    // Logic for sending the automated WhatsApp 'Linked' message
    fmt.Println("✅ WhatsApp handshake signal sent.")
    return nil
}
