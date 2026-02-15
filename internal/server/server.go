package server

import (
	"fmt"
	"net/http"
	"phane-node/internal/auth"
)

func StartServer() {
	// Styled Registration Page
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        body { background-color: #020617; color: #f8fafc; }
        .glass { background: rgba(30, 41, 59, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255,255,255,0.1); }
        .glow { box-shadow: 0 0 20px rgba(56, 189, 248, 0.3); }
    </style>
</head>
<body class="flex items-center justify-center min-h-screen">
    <div class="glass p-10 rounded-3xl w-full max-w-md glow text-center">
        <div class="text-4xl mb-6">🛰️</div>
        <h1 class="text-3xl font-extrabold mb-2 text-sky-400">Initialize Node</h1>
        <p class="text-slate-400 mb-8">Set your master password to lock your sovereign vault.</p>
        <form method="POST" class="space-y-4">
            <input type="password" name="pass" placeholder="Master Password" 
                   class="w-full bg-slate-900 border border-slate-700 p-4 rounded-xl focus:outline-none focus:border-sky-500 transition">
            <button type="submit" class="w-full bg-sky-500 py-4 rounded-xl font-bold text-slate-900 hover:bg-sky-400 transition">
                Register Local Node
            </button>
        </form>
    </div>
</body>
</html>`)
			return
		}
		
		pass := r.FormValue("pass")
		hashed, _ := auth.Register(pass)
		fmt.Fprintf(w, "<html><body style='background:#020617;color:#38bdf8;padding:50px;text-align:center;'><h1>Success!</h1><p>Node initialized securely.</p></body></html>")
	})

	fmt.Println("🚀 PHANE Dashboard alive at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
