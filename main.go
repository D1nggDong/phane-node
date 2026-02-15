package main

import (
    "embed"
    "html/template"
    "log"
    "net/http"
)

//go:embed templates/* static/*
var embeddedFiles embed.FS

func main() {
    mux := http.NewServeMux()
    mux.Handle("/static/", http.FileServer(http.FS(embeddedFiles)))

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" { http.NotFound(w, r); return }
        render(w, "templates/index.html")
    })
    mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { render(w, "templates/auth.html") })
    mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) { render(w, "templates/auth.html") })
    mux.HandleFunc("/dash", func(w http.ResponseWriter, r *http.Request) { render(w, "templates/dashboard.html") })
    mux.HandleFunc("/skills", func(w http.ResponseWriter, r *http.Request) { render(w, "templates/skills.html") })

    log.Println("🚀 PHANE FRONTEND COMPLETE: http://localhost:8080")
    http.ListenAndServe(":8080", mux)
}

func render(w http.ResponseWriter, path string) {
    tmpl, _ := template.ParseFS(embeddedFiles, path)
    tmpl.Execute(w, nil)
}
