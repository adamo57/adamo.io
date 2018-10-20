package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// New generates a new http server.
func New(r *mux.Router, address string) *http.Server {
	// create our own server to avoid long timeouts.
	return &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
