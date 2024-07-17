package main

import (
	"errors"
	"github.com/kirillgashkov/squeak/internal/server"
	"strconv"
	"time"
)

const (
	modeKey                    = "SQUEAK_MODE"
	serverHostKey              = "SQUEAK_SERVER_HOST"
	serverPortKey              = "SQUEAK_SERVER_PORT"
	serverReadHeaderTimeoutKey = "SQUEAK_SERVER_READ_HEADER_TIMEOUT"
)

const (
	defaultMode                    = modeDevelopment
	defaultServerHost              = "127.0.0.1"
	defaultServerPort              = 8000
	defaultServerReadHeaderTimeout = 1 * time.Second
)

const (
	modeDevelopment = "development"
	modeProduction  = "production"
)

type config struct {
	Mode   string
	Server *server.Config
}

func parseConfig(getenv func(string) string) (*config, error) {
	var err error

	mode := defaultMode
	if getenv(modeKey) != "" {
		mode = getenv(modeKey)
		switch mode {
		case modeDevelopment, modeProduction:
		default:
			return nil, errors.New("invalid mode")
		}
	}

	serverCfg, err := parseServerConfig(getenv)
	if err != nil {
		return nil, err
	}

	return &config{Mode: mode, Server: serverCfg}, nil
}

func parseServerConfig(getenv func(string) string) (*server.Config, error) {
	var err error

	host := defaultServerHost
	if getenv(serverHostKey) != "" {
		host = getenv(serverHostKey)
	}

	port := defaultServerPort
	if getenv(serverPortKey) != "" {
		port, err = strconv.Atoi(getenv(serverPortKey))
		if err != nil {
			return nil, errors.Join(errors.New("invalid server port"), err)
		}
	}

	readHeaderTimeout := defaultServerReadHeaderTimeout
	if getenv(serverReadHeaderTimeoutKey) != "" {
		readHeaderTimeout, err = time.ParseDuration(getenv(serverReadHeaderTimeoutKey))
		if err != nil {
			return nil, errors.Join(errors.New("invalid server read header timeout"), err)
		}
	}

	return &server.Config{
		Host:              host,
		Port:              port,
		ReadHeaderTimeout: readHeaderTimeout,
	}, nil
}
