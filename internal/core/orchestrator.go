package core

import (
	"fmt"
	"log"
	"phane-node/internal/skills"
)

// RunAutonomousTask takes a natural language prompt and decides which skill to use
func RunAutonomousTask(userPrompt string) error {
	// 1. Ask Ollama to interpret the intent
	// In a real scenario, we'd format a system prompt to return JSON
	ollama := skills.OllamaSkill{}
	decision, err := ollama.Execute("llama3", fmt.Sprintf("Analyze this task: '%s'. Which tool should I use: notion or discord? Respond with only the tool name.", userPrompt))
	
	if err != nil {
		return err
	}

	log.Printf("🤖 Orchestrator Decision: %s", decision)

	// 2. Trigger the specific skill from the Registry
	if skill, exists := skills.Registry[decision]; exists {
		log.Printf("⚡ Executing Skill: %s", skill.Name())
		// Here we would pass the decrypted keys and data
		return nil 
	}

	return fmt.Errorf("no suitable skill found for decision: %s", decision)
}
