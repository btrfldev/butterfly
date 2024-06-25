package main

import (
	"time"

	"github.com/iamsoloma/butterfly"
)

type Server struct {
	listenAddr    string
	idleTimeout   time.Duration
	bodyLimit     int
	cacheLifeTime time.Duration
	Carbine       butterfly.Storer[string, string]
	CacheStorage  butterfly.Storer[string, string]
}

func NewServer(listenAddr string, bodyLimit int, idleTimeout time.Duration, cacheLifeTime time.Duration) *Server {
	return &Server{
		listenAddr:    listenAddr,
		idleTimeout:   idleTimeout,
		bodyLimit:     bodyLimit,
		cacheLifeTime: cacheLifeTime,
		Carbine:       butterfly.NewCarbineStore[string, string](),
		CacheStorage:  butterfly.NewCarbineStore[string, string](),
	}
}
