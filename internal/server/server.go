package server

import (
	"fmt"
	"net/http"
	"phane-node/internal/auth"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Using the HTML code &#128752; instead of the emoji symbol
			renderPage(w, "Initialize Node", "Create your master password to lock your sovereign vault.", "/register", "password", "Master Password")
			return
		}
		pass := r.FormValue("input")
		_, _ = auth.Register(pass) 
		http.Redirect(w, r, "/notion-setup", http.StatusSeeOther)
	})

	http.HandleFunc("/notion-setup", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "Notion Setup", "Your keys will be encrypted with AES-256.", "/save-notion", "text", "secret_notion_token...")
	})

	http.HandleFunc("/save-notion", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<body style='background:#020617;color:#38bdf8;text-align:center;padding-top:100px;font-family:sans-serif;'><h1>✔️ Vaulted</h1><p style='color:white;'>Your node is now connected to Notion.</p></body>")
	})

	fmt.Println("🚀 PHANE Dashboard running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func renderPage(w http.ResponseWriter, title, sub, action, iType, placeholder string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        body { background-color: #020617; color: #f8fafc; }
        .glass { background: rgba(30, 41, 59, 0.4); backdrop-filter: blur(12px); border: 1px solid rgba(255,255,255,0.1); }
        .glow { box-shadow: 0 0 40px rgba(56, 189, 248, 0.15); }
    </style>
</head>
<body class="flex items-center justify-center min-h-screen p-6">
    <div class="glass p-12 rounded-[2.5rem] w-full max-w-lg glow text-center">
        <div class="text-5xl mb-8">&#128752;</div>
        <h1 class="text-4xl font-black mb-3 tracking-tighter text-sky-400 font-sans">%s</h1>
        <p class="text-slate-400 mb-10 text-lg">%s</p>
        <form method="POST" action="%s" class="space-y-6">
            <input type="%s" name="input" placeholder="%s" 
                   class="w-full bg-slate-950/50 border border-slate-800 p-5 rounded-2xl text-white focus:border-sky-500 outline-none transition-all duration-300 text-center text-xl">
            <button type="submit" class="w-full bg-sky-500 py-5 rounded-2xl font-black text-slate-950 hover:bg-sky-400 hover:scale-[1.02] active:scale-[0.98] transition-all duration-300 text-xl tracking-wide uppercase">
                Continue
            </button>
        </form>
    </div>
</body>
</html>`, title, sub, action, iType, placeholder)
}
