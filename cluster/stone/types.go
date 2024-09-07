package main

import (
	"time"

	btrflstore "github.com/iamsoloma/butterfly/store"
)

type Server struct {
	listenAddr  string
	idleTimeout time.Duration
	bodyLimit   int
	Memory      btrflstore.MemoryStore
	StoragePath     string
}

func NewServer(listenAddr string, bodyLimit int, idleTimeout time.Duration, StoragePath string) *Server {
	return &Server{
		listenAddr:  listenAddr,
		idleTimeout: idleTimeout,
		bodyLimit:   bodyLimit,
		Memory:      *btrflstore.NewMemoryStore(),
		StoragePath: StoragePath,
	}
}
