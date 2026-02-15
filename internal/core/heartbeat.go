package core

import (
	"time"
	"runtime"
)

func StartHeartbeat() {
	ticker := time.NewTicker(10 * time.Second) // Fast pulse for demo
	go func() {
		for range ticker.C {
			// 1. Perform the real work
			checkAndExecuteTasks()

			// 2. Broadcast the "Pulse" to the UI
			GlobalHub.Broadcast(map[string]string{
				"type": "pulse",
				"timestamp": time.Now().Format("15:04:05"),
			})

			// 3. Keep Pi 4B 1GB lean
			runtime.GC()
		}
	}()
}

func checkAndExecuteTasks() {
	// (Existing SQLite logic here)
}
