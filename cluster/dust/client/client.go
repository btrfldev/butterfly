package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/iamsoloma/butterfly"
	"github.com/iamsoloma/butterfly/logger"
)

type Agent struct {
	DustAddress string
}

func (a *Agent) Put(objects []butterfly.Object) (err error) {
	errLocation := "btrfl.cluster.dust.client.Put"
	client := &http.Client{}
	status := butterfly.Status{}

	body, err := json.Marshal(objects)
	logger.CheckErr(err, errLocation, "Can`t marshal json")
	req, err := http.NewRequest(http.MethodGet, a.DustAddress, bytes.NewBuffer(body))
	resp, err := client.Do(req)
	logger.CheckErr(err, errLocation, "Can`t do request to dust")

	defer resp.Body.Close()
	respBody := []byte{}
	_, err = resp.Body.Read(respBody)
	logger.CheckErr(err, errLocation, "Can`t read dust`s responce")
	err = json.Unmarshal(respBody, &status)
	logger.CheckErr(err, errLocation, "Can`t unmarshal dust`s responce")

	if status.Status == http.StatusOK {
		return nil
	} else {
		return logger.NewErr(errLocation, strconv.Itoa(status.Status)+":"+status.Message)
	}
}
