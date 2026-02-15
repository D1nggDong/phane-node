package core

import (
	"fmt"
	"phane-node/internal/security"
)

// DeployOneClickSkill ensures the user has a Pro license before automation
func DeployOneClickSkill(skillID string, userKey string) error {
	isPro, err := security.VerifyProStatus(userKey)
	if err != nil || !isPro {
		return fmt.Errorf("this is a Pro feature. Please upgrade at phane.ai")
	}

	// Logic to automatically pull API keys from Phane Cloud
	fmt.Printf("🚀 Deploying %s via One-Click Automation\n", skillID)
	return nil
}
