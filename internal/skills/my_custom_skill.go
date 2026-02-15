package skills

type MyCustomSkill struct{}

func (s MyCustomSkill) Name() string {
    return "My Custom Tool"
}

func (s MyCustomSkill) Description() string {
    return "A manually configured automation tool."
}

// Execute handles the logic when the agent triggers this skill
func (s MyCustomSkill) Execute(payload string) error {
    // Your custom logic here (e.g., API calls, local script execution)
    return nil
}
