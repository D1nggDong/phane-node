package security

import (
	"log"
	"os"
)

// EmergencyLock wipes the Master Key from RAM and shuts down the process
func EmergencyLock() {
	lock.Lock()
	sessionKey = "" // Wipe the key
	lock.Unlock()

	log.Println("🚨 EMERGENCY LOCK ACTIVATED: Vault keys purged from RAM.")
	
	// Optional: Exit the program to ensure total isolation
	os.Exit(0)
}
