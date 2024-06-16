package main

import (
	"time"

	"github.com/iamsoloma/butterfly"
)


type Server struct {
	listenAddr  string
	idleTimeout time.Duration
	bodyLimit   int
	Carbine        butterfly.Storer[string, string]
	Meta        butterfly.Storer[string, string]
}

func NewServer(listenAddr string, bodyLimit int, idleTimeout time.Duration) *Server {
	return &Server{
		listenAddr:  listenAddr,
		idleTimeout: idleTimeout,
		bodyLimit:   bodyLimit,
		Carbine:        butterfly.NewCarbineStore[string, string](),
		Meta:        butterfly.NewCarbineStore[string, string](),
	}
}
