package security

import "phane-node/internal/db"

// SaveLicense stores the Pro key in the local encrypted SQLite database
func SaveLicense(key string) error {
	_, err := db.Instance.Exec("INSERT OR REPLACE INTO settings (key, value) VALUES ('pro_license_key', ?)", key)
	return err
}

// GetSavedLicense retrieves the stored key for the heartbeat check
func GetSavedLicense() string {
	var key string
	db.Instance.QueryRow("SELECT value FROM settings WHERE key = 'pro_license_key'").Scan(&key)
	return key
}
