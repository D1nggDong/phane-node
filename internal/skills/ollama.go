package skills

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type OllamaSkill struct{}

func (o OllamaSkill) Name() string { return "Ollama Intelligence" }
func (o OllamaSkill) Description() string { return "Connect local LLMs (Llama 3, Mistral) for autonomous reasoning." }

func (o OllamaSkill) Execute(model string, prompt string) (string, error) {
	url := "http://localhost:11434/api/generate"
	payload, _ := json.Marshal(map[string]interface{}{
		"model":  model,
		"prompt": prompt,
		"stream": false,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["response"].(string), nil
}
