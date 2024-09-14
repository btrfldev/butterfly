package main

import (
	"time"

	btrflstore "github.com/iamsoloma/butterfly/store"
)

type Server struct {
	listenPort  string
	idleTimeout time.Duration
	bodyLimit   int
	Memory      btrflstore.MemoryStore
}

func NewServer(listenPort string, bodyLimit int, idleTimeout time.Duration) *Server {
	return &Server{
		listenPort:  listenPort,
		idleTimeout: idleTimeout,
		bodyLimit:   bodyLimit,
		Memory:      *btrflstore.NewMemoryStore(),
	}
}
