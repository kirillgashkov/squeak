package main

import (
	"errors"
	"fmt"
	"github.com/kirillgashkov/squeak/internal/server"
	"net/http"
	"os"
)

func main() {
	if err := run(os.Getenv); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(getenv func(string) string) error {
	cfg, err := parseConfig(getenv)
	if err != nil {
		return err
	}

	srv := server.New(cfg.server)

	_, _ = fmt.Printf("starting server at %s\n", srv.Addr)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
