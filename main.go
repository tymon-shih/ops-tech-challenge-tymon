package main

import (
	"net/http"
	"os"
)

// var requestIPs []string

func main() {
	h := &handler{
		key:   []byte(os.Getenv("SECRET")),
		stats: make(map[string]uint64),
	}
	http.HandleFunc("/token", h.token)
	http.HandleFunc("/metrics", h.metrics)
	http.HandleFunc("/health", h.health)
	http.ListenAndServe(":8080", nil)
}
