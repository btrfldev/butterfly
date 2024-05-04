package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type dustGetRequest struct {
	Lib string `json:"lib"`
	Key string `json:"key"`
}
type dustGetResponce struct {
	Value string `json:"value"`
}

func (c *Client) Get() (has bool, err error) {
	reqData := dustGetRequest{Lib: c.Data.Lib, Key: c.Data.Key}
	respData := dustGetResponce{}
	b, err := json.Marshal(&reqData)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("GET", c.Address, bytes.NewBuffer(b))
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
	} else if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return true, err
		}
		err = json.Unmarshal(body, &respData)
		if err != nil {
			return true, err
		}

		c.Data.Value = respData.Value
	}

	return true, nil
}
