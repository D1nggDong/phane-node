package skills

import "fmt"

type Skill interface {
	Name() string
	Connect() error
	Execute(cmd string) (string, error)
}

// Registry stores all 50+ integrations
var Registry = make(map[string]Skill)

func Register(s Skill) {
	Registry[s.Name()] = s
	fmt.Printf("✅ Skill Loaded: %s\n", s.Name())
}
