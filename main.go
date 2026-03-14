package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	requestCount int64
	currentRPS   int64
)

func main() {
	// Start a ticker to calculate RPS every second
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			// Swap the counter with 0 and get the previous value
			rps := atomic.SwapInt64(&requestCount, 0)
			atomic.StoreInt64(&currentRPS, rps)
		}
	}()

	// Target endpoint to flood
	http.HandleFunc("/target", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requestCount, 1)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API endpoint for frontend to get the current RPS
	http.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		rps := atomic.LoadInt64(&currentRPS)

		json.NewEncoder(w).Encode(map[string]int64{
			"requests_per_second": rps,
		})
	})

	// Serve the frontend
	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Println("Layer 7 Dstat running on port :8080")
	log.Println("Send GET/POST to http://<server-ip>:8080/target to test load")
	log.Println("View dashboard at http://server-ip>:8080/")

	server := &http.Server{
		Addr:           ":8080",
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
