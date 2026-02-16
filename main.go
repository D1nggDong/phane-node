package main

import (
	"fmt"
	"phane-node/internal/auth"
	"phane-node/internal/server"
	"log"
)

func main() {
	fmt.Println("---------------------------------")
	fmt.Println("| PHANE Sovereign OS - v1.0.0   |")
	fmt.Println("---------------------------------")
	
	err := auth.InitDB()
	if err != nil {
		log.Fatal("Failed to open vault:", err)
	}
	fmt.Println("🗄️ Vault Database: READY")
	fmt.Println("🚀 Node is listening via Cloudflare Tunnel...")
	
	server.StartServer()
}
