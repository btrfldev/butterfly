package main

import (
	"time"

	"github.com/iamsoloma/butterfly"
)


type Server struct {
	listenAddr  string
	idleTimeout time.Duration
	bodyLimit   int
	Dust        butterfly.Storer[string, string]
	Info        butterfly.Storer[string, string]
}

func NewServer(listenAddr string, bodyLimit int, idleTimeout time.Duration) *Server {
	return &Server{
		listenAddr:  listenAddr,
		idleTimeout: idleTimeout,
		bodyLimit:   bodyLimit,
		Dust:        butterfly.NewDustStore[string, string](),
		Info:        butterfly.NewDustStore[string, string](),
	}
}
