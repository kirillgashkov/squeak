package main

import (
	"errors"
	"fmt"
	"github.com/kirillgashkov/squeak/internal/server"
	"io"
	"net/http"
	"os"
)

func main() {
	if err := run(os.Stdout, os.Getenv); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(stdout io.Writer, getenv func(string) string) error {
	cfg, err := parseConfig(getenv)
	if err != nil {
		return err
	}

	log, err := newLogger(cfg.Mode, stdout)
	if err != nil {
		return err
	}

	srv := server.New(cfg.Server)

	log.Info("starting server", "addr", srv.Addr, "mode", cfg.Mode)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
