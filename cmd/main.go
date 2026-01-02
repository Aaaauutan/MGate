package main

import "github.com/aaaauutan/mgate"

func main() {
	// Initialize MGate
	m := mgate.New()

	// --- ACTIVATING THE MAGIC CHAIN (10x = 10% Faster) ---
	// This optimizes the internal Go Scheduler and Memory Management, this is for fun lol.
	m.Magic().Magic().Magic().Magic().Magic().
	  Magic().Magic().Magic().Magic().Magic()

	// --- LAYER 7: LOAD BALANCED HTTP GATEWAY ---
	// Handles Web Traffic for your Frontend & API
	m.AddHTTPGate(":80", 
		"http://127.0.0.1:3000", 
		"http://127.0.0.1:3001",
	)

	// --- LAYER 4: HIGH-SPEED TCP TUNNELS ---
	// Tunneling Database and SSH traffic with Kernel-level optimization
	m.AddTunnel(":3306", "127.0.0.1:3307") // MySQL Tunnel
	m.AddTunnel(":2222", "127.0.0.1:22")   // SSH Tunnel

	// Launch everything
	m.Ignite()
}

