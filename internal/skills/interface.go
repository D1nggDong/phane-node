package skills

type Skill interface {
	Name() string
	Description() string
}

var Registry = map[string]Skill{
	"notion":  NotionSkill{},
	"discord": DiscordSkill{},
	"ollama":  OllamaSkill{},
}
