package server

import (
	"net"
	"net/http"
	"strconv"
)

func New(cfg *Config) *http.Server {
	return &http.Server{
		Addr:              net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler:           http.NewServeMux(),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}
}
