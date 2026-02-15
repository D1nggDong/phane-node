package core

import (
	"database/sql"
	"net/http"
	"phane-node/internal/security"
	_ "github.com/mattn/go-sqlite3"
)

// SaveSkillKey handles the form submission for new Marketplace skills
func SaveSkillKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	skillName := r.FormValue("skill_name")
	plainAPIKey := r.FormValue("api_key")

	// 1. Get the Master Key from volatile memory
	master := security.GetSessionKey()
	if master == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// 2. Encrypt the key using AES-256
	encrypted, err := security.EncryptKey(plainAPIKey, master)
	if err != nil {
		http.Error(w, "Encryption failed", 500)
		return
	}

	// 3. Save to SQLite
	db, _ := sql.Open("sqlite3", "./phane_nodes.db")
	defer db.Close()

	_, err = db.Exec("INSERT INTO tasks (name, action_skill, status) VALUES (?, ?, ?)", 
		skillName, encrypted, "active")

	if err != nil {
		http.Error(w, "Database error", 500)
		return
	}

	// 4. Send back to dashboard
	http.Redirect(w, r, "/dash", http.StatusSeeOther)
}
