package discord

import (
	"fmt"
	"phane-node/internal/security"
	"github.com/bwmarrin/discordgo"
)

type DiscordSkill struct {
	Session *discordgo.Session
	Token   string // Encrypted
}

func (d *DiscordSkill) Name() string { return "Discord Bridge" }

func (d *DiscordSkill) Connect(masterKey string) error {
	// Decrypt the token locally before connecting
	decryptedToken, err := security.DecryptKey(d.Token, masterKey)
	if err != nil { return err }

	dg, err := discordgo.New("Bot " + decryptedToken)
	if err != nil { return err }
	
	d.Session = dg
	return d.Session.Open()
}

func (d *DiscordSkill) Execute(channelID string, msg string) error {
	_, err := d.Session.ChannelMessageSend(channelID, msg)
	return err
}
