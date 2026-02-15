package security

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// GetNodeID generates a unique hash based on the hostname and hardware info
func GetNodeID() string {
	hostname, _ := os.Hostname()
	// We hash the hostname to create a unique but anonymous ID
	hash := sha256.Sum256([]byte(hostname))
	return fmt.Sprintf("PHANE-%x", hash[:4])
}
