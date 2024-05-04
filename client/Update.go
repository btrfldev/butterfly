package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type dustUpdateRequest struct {
	Lib string `json:"lib"`
	Key string `json:"key"`
}


func (c *Client )Update() (has bool, err error) {
	reqData := dustUpdateRequest{
		Lib: c.Data.Lib,
		Key: c.Data.Key,
	}

	b, err := json.Marshal(&reqData)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("GET", c.Address+"/update", bytes.NewBuffer(b))
	if err != nil {
		return false, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusInternalServerError {
		return false, nil
	} else if res.StatusCode == http.StatusBadRequest {
		return false, errors.New("bad request")
	}else if res.StatusCode!=http.StatusOK {
		return false, errors.New("unexpected error")
	} else if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return true, nil
}