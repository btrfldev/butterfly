package main

import (
	"time"

	btrflstore "github.com/iamsoloma/butterfly/store"
)

type Server struct {
	listenAddr    string
	idleTimeout   time.Duration
	bodyLimit     int
	Carbine       btrflstore.Storer[string, string]
}

func NewServer(listenAddr string, bodyLimit int, idleTimeout time.Duration) *Server {
	return &Server{
		listenAddr:    listenAddr,
		idleTimeout:   idleTimeout,
		bodyLimit:     bodyLimit,
		Carbine:       btrflstore.NewCarbineStore[string, string](),
	}
}
