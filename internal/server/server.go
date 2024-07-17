package server

import (
	"net"
	"net/http"
	"strconv"
)

func New(cfg *Config) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handleGetHealth)

	return &http.Server{
		Addr:              net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler:           mux,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}
}

func handleGetHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		panic(err)
	}
}
