package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type handler struct {
	key   []byte
	stats map[string]uint64
	mu    sync.Mutex
}

func (h handler) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	h.mu.Lock()
	h.stats["requests"] += 1
	h.mu.Unlock()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	out := createMAC(body, h.key)
	fmt.Fprintf(w, "%x", out)

	w.WriteHeader(http.StatusAccepted)
}

func (h handler) metrics(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(h.stats)
	w.WriteHeader(http.StatusOK)
}

func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
