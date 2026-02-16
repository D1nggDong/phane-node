package main

import (
	"fmt"
	"phane-node/internal/server"
)

func main() {
	// Clean ASCII for the terminal to stop the weird symbols
	fmt.Println("---------------------------------")
	fmt.Println("| PHANE Sovereign OS - v1.0.0   |")
	fmt.Println("| Status: 1GB RAM OPTIMIZED     |")
	fmt.Println("---------------------------------")
	fmt.Println("🚀 Node is listening via Cloudflare Tunnel...")
	
	server.StartServer()
}
