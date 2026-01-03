# MGate v1.0 🚀

A high-performance, multi-protocol networking gateway written in Go. Designed for extreme speed, low latency, and protocol flexibility.

## Features
- **Layer 7 Flash Proxy**: High-speed HTTP/HTTPS reverse proxy with keep-alive optimization.
- **Layer 4 Zero-Copy Tunnel**: Raw TCP tunneling using kernel-level data transfer (Splice optimization).
- **Atomic Round-Robin**: Lock-free load balancing for massive concurrency handling.
- **The Magic Engine**: Chained optimization logic to tune Go runtime scheduler dynamically.
- **Zero-Config Logic**: Single binary deployment, no complex configuration files.

## Installation
```bash
go get github.com/meowtux/mgate
```
Quick Start
```go
package main

import "github.com/meowtux/mgate"

func main() {
    // Ignite engine
    gate := mgate.New()

    // Optimization magic (Chaining up to 10x)
    gate.Magic().Magic().Magic()

    // HTTP Load Balancer
    gate.AddHTTPGate(":80", "http://127.0.0.1:3000", "http://127.0.0.1:3001")

    // TCP Tunnel (MySQL/SSH)
    gate.AddTunnel(":3306", "127.0.0.1:3307")

    gate.Ignite()
}
```
Optimization Specs
 * Runtime: Pre-tuned GOMAXPROCS for I/O bound tasks.
 * Memory: Optimized GC (Garbage Collector) thresholds for networking throughput.
 * I/O: Bi-directional stream copying with non-blocking logic.

## MIT License. Built for performance.

- this was made by : MeowTux
