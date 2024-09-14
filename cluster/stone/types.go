package main

import (
	"time"

	btrflstore "github.com/iamsoloma/butterfly/store"
)

type Server struct {
	listenPort  string
	idleTimeout time.Duration
	//bodyLimit   int
	Memory      btrflstore.MemoryStore
	StoragePath string
	DustAddress string
	NodeInfo    NodeInfo
}

type NodeInfo struct {
	ID            int
	Region        string
	Role          string
	PublicAddress string
}

func NewServer(listenPort string /*bodyLimit int,*/, idleTimeout time.Duration, StoragePath, PublicAddress, DustAddress string) *Server {

	return &Server{
		listenPort:  listenPort,
		idleTimeout: idleTimeout,
		//bodyLimit:   bodyLimit,
		Memory:      *btrflstore.NewMemoryStore(),
		StoragePath: StoragePath,
		DustAddress: DustAddress,
		NodeInfo: NodeInfo{
			PublicAddress: PublicAddress,
		},
	}
}
