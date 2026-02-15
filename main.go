package main

import (
	"fmt"
	"phane-node/internal/server"
)

func main() {
	fmt.Println("🛰️ PHANE Sovereign OS - v1.0.0")
	fmt.Println("1GB RAM Optimization: ACTIVE")
	
	// Start the local dashboard and auth listener
	server.StartServer()
}
