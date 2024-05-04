package client

import "net/http"

type Data struct {
	Lib   string `json:"lib"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Client struct {
	Address string
	Data
	client *http.Client
}