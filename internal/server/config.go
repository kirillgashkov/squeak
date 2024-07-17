package server

import "time"

type Config struct {
	Host              string
	Port              int
	ReadHeaderTimeout time.Duration
}
