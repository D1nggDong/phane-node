package skills

import (
	"bytes"
	"net/http"
)

type GenericLLMSkill struct {
	Endpoint string
}

func (g GenericLLMSkill) Name() string { return "External LLM Bridge" }
func (g GenericLLMSkill) Description() string { return "Connect to any OpenAI-compatible API (LM Studio, LocalAI, etc.)" }

func (g GenericLLMSkill) Execute(prompt string) error {
	// Logic to send prompt to g.Endpoint
	return nil
}
