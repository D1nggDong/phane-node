package network

import (
	"log"
	"phane-node/internal/security"
)

// ActivateTunnel starts the secure remote bridge for Pro users
func ActivateTunnel() {
	// 1. Verify Subscription
	nodeID := security.GetNodeID()
	isPro, _ := CheckSubscription(nodeID)

	if !isPro {
		log.Println("❌ Tunneling is a Pro feature. Upgrade at /upgrade")
		return
	}

	// 2. Logic to start the outbound tunnel (Conceptual)
	log.Println("🚀 Pro Tunnel Active: https://" + nodeID + ".phane.ai")
}
