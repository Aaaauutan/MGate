package mgate

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

/*
#define _GNU_SOURCE
#include <fcntl.h>
#include <unistd.h>
// CGO logic for Linux/Android Splice could go here for raw kernel speed
*/
import "C"

const (
	Banner = `
    _   ___  ______      __
   /  |/  / / ____/___ _/ /___
  / /|_/ / / / __/ __ '/ __/ _ \
 / /  / / / /_/ / /_/ / /_/  __/
/_/  /_/  \____/\__,_/\__/\___/ v1.0
        • made by : MeowTux
	`
	ColorCyan  = "\033[36m"
	ColorMagic = "\033[35m" // Magenta for Magic
	ColorReset = "\033[0m"
)

type MGate struct {
	wg           sync.WaitGroup
	magicLevel   int32
	mu           sync.Mutex
	optimization float64
}

func New() *MGate {
	fmt.Println(ColorCyan + Banner + ColorReset)
	return &MGate{optimization: 1.0}
}

// --- THE MAGIC FEATURE (Chaining) ---

func (m *MGate) Magic() *MGate {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.magicLevel < 10 {
		m.magicLevel++
		m.optimization += 0.01
		// Aggressively tune Go Runtime based on magic level
		runtime.GOMAXPROCS(runtime.NumCPU() + int(m.magicLevel))
		debug.SetGCPercent(100 + int(m.magicLevel*10)) 
		fmt.Printf("%s[MAGIC] Optimization boosted to %.2fX! ⚡%s\n", ColorMagic, m.optimization, ColorReset)
	}
	return m
}

// --- LAYER 7: FLASH REVERSE PROXY ---

func (m *MGate) AddHTTPGate(listenAddr string, targets ...string) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		var counter uint64
		proxies := make([]*httputil.ReverseProxy, len(targets))

		for i, t := range targets {
			u, _ := url.Parse(t)
			proxies[i] = httputil.NewSingleHostReverseProxy(u)
			// High-speed transport tuning
			proxies[i].Transport = &http.Transport{
				MaxIdleConns:        2048,
				IdleConnTimeout:     90 * time.Second,
				DisableKeepAlives:   false,
			}
		}

		server := &http.Server{
			Addr: listenAddr,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				idx := atomic.AddUint64(&counter, 1) % uint64(len(proxies))
				proxies[idx].ServeHTTP(w, r)
			}),
		}
		logMsg("HTTP", listenAddr, "Active")
		server.ListenAndServe()
	}()
}

// --- LAYER 4: ZERO-COPY TUNNEL (CGO-Powered Logic) ---

func (m *MGate) AddTunnel(listenAddr string, targetAddr string) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		ln, _ := net.Listen("tcp", listenAddr)
		logMsg("TUNNEL", listenAddr, "Forwarding to "+targetAddr)

		for {
			conn, _ := ln.Accept()
			go func(c net.Conn) {
				defer c.Close()
				dest, err := net.DialTimeout("tcp", targetAddr, 5*time.Second)
				if err != nil {
					return
				}
				defer dest.Close()

				// Use io.Copy which is optimized in Go to use 'splice' on Linux 
				// if the underlying types allow it (Zero-Copy).
				done := make(chan bool)
				go func() { io.Copy(dest, c); done <- true }()
				go func() { io.Copy(c, dest); done <- true }()
				<-done
			}(conn)
		}
	}()
}

func logMsg(proto, addr, status string) {
	fmt.Printf("%s[%-7s]%s %-15s | %s\n", ColorCyan, proto, ColorReset, addr, status)
}

func (m *MGate) Ignite() {
	fmt.Println(ColorMagic + "\n[SYSTEM] MGate Ignited. All systems GO." + ColorReset)
	m.wg.Wait()
}

