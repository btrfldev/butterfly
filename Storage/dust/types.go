package main

import (
	"time"

	//"github.com/iamsoloma/butterfly"
	btrflstore "github.com/iamsoloma/butterfly/store"
)

type Server struct {
	listenAddr    string
	idleTimeout   time.Duration
	bodyLimit     int
	cacheLifeTime time.Duration
	Carbine       btrflstore.Storer[string, string] /*butterfly.Storer[string, string]*/
	CacheStorage  btrflstore.Storer[string, string]
}

func NewServer(listenAddr string, bodyLimit int, idleTimeout time.Duration, cacheLifeTime time.Duration) *Server {
	return &Server{
		listenAddr:    listenAddr,
		idleTimeout:   idleTimeout,
		bodyLimit:     bodyLimit,
		cacheLifeTime: cacheLifeTime,
		Carbine:       btrflstore.NewCarbineStore[string, string](),
		CacheStorage:  btrflstore.NewCarbineStore[string, string](),
	}
}
