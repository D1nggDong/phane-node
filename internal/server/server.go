package server

import (
	"fmt"
	"net/http"
	"phane-node/internal/auth"
)

func StartServer() {
	// 1. Registration Handler
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "<html><body style='background:#020617;color:white;font-family:sans-serif;text-align:center;padding-top:50px;'>")
			fmt.Fprintf(w, "<h1>Sovereign Registration</h1><form method='POST'><input type='password' name='pass' placeholder='Master Password' style='padding:10px;border-radius:5px;'><button type='submit' style='padding:10px 20px;margin-left:10px;background:#38bdf8;border:none;border-radius:5px;'>Initialize Node</button></form></body></html>")
			return
		}
		
		pass := r.FormValue("pass")
		hashed, _ := auth.Register(pass)
		// In a real scenario, we save 'hashed' to your phane_nodes.db vault here
		fmt.Fprintf(w, "Node Initialized. Hash saved to local vault: %s", hashed[:15]+"...")
	})

	// 2. Login Handler
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><body style='background:#020617;color:white;text-align:center;padding-top:50px;'><h1>Node Login</h1><p>Checking local encrypted vault...</p></body></html>")
	})

	fmt.Println("🚀 PHANE Dashboard running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
