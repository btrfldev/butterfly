package main

import "os"



func(s *Server) InitNode(DustAddress string) {
	if err := os.MkdirAll(s.StoragePath, 0777); err!=nil{
		panic(err)
	}
}