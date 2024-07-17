package main

import (
	"fmt"
	"io"
	"log/slog"
)

func newLogger(mode string, stdout io.Writer) (*slog.Logger, error) {
	var handler slog.Handler
	switch mode {
	case modeDevelopment:
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		handler = slog.NewTextHandler(stdout, opts)
	case modeProduction:
		opts := &slog.HandlerOptions{Level: slog.LevelInfo}
		handler = slog.NewJSONHandler(stdout, opts)
	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}

	return slog.New(handler), nil
}
