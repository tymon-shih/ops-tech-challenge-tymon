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

func (h handler) health(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusNotImplemented)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET requests are allowed!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok!"))

	data := []byte{79, 107, 33}
	fmt.Fprintf(w, "\nAn alternate way to write \"Ok!\" is the following: %v", string(data))

}

func (h handler) token(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	w.Write([]byte("Only POST requests are allowed!"))
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading body"))
		return
	}

	// if len(body) == 0 {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Empty body"))
	// 	return
	// }

	w.WriteHeader(http.StatusAccepted)

	out := createMAC(body, h.key)
	// output := fmt.Sprintf("HMAC value is: %v", out)

	h.mu.Lock()
	h.stats["requests"] += 1
	h.mu.Unlock()

	// w.Write([]byte(output))
	// w.Write([]byte("\nHMAC in hexademical format: "))
	fmt.Fprintf(w, "%x", out)
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
