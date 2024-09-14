package main

import (
	"time"

	btrflstore "github.com/iamsoloma/butterfly/store"
)

type Server struct {
	listenAddr  string
	idleTimeout time.Duration
	//bodyLimit   int
	Memory      btrflstore.MemoryStore
	StoragePath string
	DustAddress string
	NodeInfo    NodeInfo
}

type NodeInfo struct {
	ID     string
	Region string
	Role   string
}

func NewServer(listenAddr string /*bodyLimit int,*/, idleTimeout time.Duration, StoragePath, DustAddress string) *Server {
	return &Server{
		listenAddr:  listenAddr,
		idleTimeout: idleTimeout,
		//bodyLimit:   bodyLimit,
		Memory:      *btrflstore.NewMemoryStore(),
		StoragePath: StoragePath,
		DustAddress: DustAddress,
		NodeInfo: NodeInfo{

		},
	}
}
