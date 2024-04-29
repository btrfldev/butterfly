package main

import (
	"time"

	"github.com/TinajXD/butterfly"
)

type Query struct {
	Lib   string `json:"lib"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Server struct {
	listenAddr  string
	idleTimeout time.Duration
	bodyLimit   int
	Dust        butterfly.Storer[string, string]
}

func NewServer(listenAddr string, bodyLimit int, idleTimeout time.Duration) *Server {
	return &Server{
		listenAddr: listenAddr,
		idleTimeout: idleTimeout,
		bodyLimit: bodyLimit,
		Dust:       butterfly.NewDustStore[string, string](),
	}
}
