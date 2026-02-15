package security

import "sync"

var (
	sessionKey string
	lock       sync.RWMutex
)

// SetSessionKey stores the key in RAM only
func SetSessionKey(key string) {
	lock.Lock()
	defer lock.Unlock()
	sessionKey = key
}

// GetSessionKey retrieves the key for decryption tasks
func GetSessionKey() string {
	lock.RLock()
	defer lock.RUnlock()
	return sessionKey
}

// IsVaultLocked checks if the key is missing from RAM
func IsVaultLocked() bool {
	lock.RLock()
	defer lock.RUnlock()
	return sessionKey == ""
}
