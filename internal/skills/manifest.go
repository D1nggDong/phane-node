package skills

import (
	"fmt"
	"phane-node/internal/security"
)

type Integration struct {
	ID       string
	Name     string
	Category string
	Icon     string
}

// The Master List of 50+ Integrations for PHANE OS
var MasterManifest = []Integration{
	// Communication
	{"discord", "Discord", "Social", "💬"},
	{"telegram", "Telegram", "Social", "✈️"},
	{"whatsapp", "WhatsApp", "Social", "🟢"},
	{"slack", "Slack", "Work", "🌈"},
	{"messenger", "Messenger", "Social", "🔵"},
	
	// Productivity & Code
	{"notion", "Notion", "Productivity", "📓"},
	{"google_cal", "Google Calendar", "Productivity", "📅"},
	{"gmail", "Gmail", "Email", "✉️"},
	{"outlook", "Outlook", "Email", "📧"},
	{"github", "GitHub", "Dev", "💻"},
	{"gitlab", "GitLab", "Dev", "🦊"},
	{"trello", "Trello", "Productivity", "📋"},
	{"linear", "Linear", "Productivity", "📐"},
	{"clickup", "ClickUp", "Productivity", "🆙"},
	{"asana", "Asana", "Productivity", "❇️"},

	// Storage & Cloud
	{"drive", "Google Drive", "Cloud", "📂"},
	{"dropbox", "Dropbox", "Cloud", "📦"},
	{"icloud", "iCloud", "Cloud", "☁️"},
	{"aws", "AWS", "Infrastructure", "☁️"},
	{"digitalocean", "DigitalOcean", "Infrastructure", "🌊"},

	// AI & Search
	{"openai", "OpenAI", "AI", "🤖"},
	{"anthropic", "Anthropic", "AI", "🧠"},
	{"groq", "Groq", "AI", "⚡"},
	{"perplexity", "Perplexity", "Search", "🔍"},
	{"brave", "Brave Search", "Search", "🦁"},

	// Finance & Crypto
	{"stripe", "Stripe", "Finance", "💳"},
	{"paypal", "PayPal", "Finance", "🅿️"},
	{"coinbase", "Coinbase", "Crypto", "🪙"},
	{"metamask", "MetaMask", "Crypto", "🦊"},
	{"binance", "Binance", "Crypto", "🔶"},

	// Smart Home & IoT
	{"homekit", "Apple HomeKit", "IoT", "🏠"},
	{"google_home", "Google Home", "IoT", "🏠"},
	{"alexa", "Amazon Alexa", "IoT", "🗣️"},
	{"hue", "Philips Hue", "IoT", "💡"},
	{"sonos", "Sonos", "Audio", "🔊"},

	// Social Media & Content
	{"twitter", "X / Twitter", "Social", "🐦"},
	{"instagram", "Instagram", "Social", "📸"},
	{"linkedin", "LinkedIn", "Professional", "💼"},
	{"youtube", "YouTube", "Media", "📺"},
	{"spotify", "Spotify", "Media", "🎧"},
	{"reddit", "Reddit", "Social", "🤖"},
	{"medium", "Medium", "Content", "📝"},

	// Automation & Misc
	{"zapier", "Zapier", "Automation", "🧡"},
	{"ifttt", "IFTTT", "Automation", "🔵"},
	{"make", "Make.com", "Automation", "🟣"},
	{"shopify", "Shopify", "E-commerce", "🛍️"},
	{"ghost", "Ghost", "CMS", "👻"},
	{"wordpress", "WordPress", "CMS", "🌐"},
	{"zoom", "Zoom", "Video", "📹"},
	{"calendly", "Calendly", "Productivity", "📅"},
}

func GetSkill(id string) (Integration, error) {
	for _, s := range MasterManifest {
		if s.ID == id { return s, nil }
	}
	return Integration{}, fmt.Errorf("skill not found")
}

func EncryptAndStore(id string, rawToken string, masterKey string) (string, error) {
	// Strictly follow security requirement [cite: 2026-02-12]
	encrypted, err := security.EncryptKey(rawToken, masterKey)
	if err != nil { return "", err }
	return encrypted, nil
}
