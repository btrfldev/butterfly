package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type dustPutRequest struct {
	Lib string `json:"lib"`
	Key string `json:"key"`
	Value string `json:"value"`
}


func (c *Client) Put() (err error) {
	reqData := dustPutRequest{
		Lib:   c.Data.Lib,
		Key:   c.Data.Key,
		Value: c.Data.Value,
	}

	b, err := json.Marshal(&reqData)
	if err != nil {
		return  err
	}
	req, err := http.NewRequest("GET", c.Address+"/put", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusInternalServerError {
		return errors.New("empty data")
	} else if res.StatusCode == http.StatusBadRequest {
		return  errors.New("bad request")
	}else if res.StatusCode!=http.StatusOK {
		return errors.New("unexpected error")
	} else if res.StatusCode == http.StatusOK {
		return nil
	}

	return nil
}