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


	query, err := json.Marshal(butterfly.Query{Objects: objects})
	if lerr := logger.CheckErr(err, errLocation, "Can`t marshal json"); lerr != nil {
		panic(lerr)
	}


	req, err := http.NewRequest(http.MethodGet, a.DustAddress+"/put", bytes.NewBuffer(query))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if lerr := logger.CheckErr(err, errLocation, "Can`t do request to dust"); lerr != nil {
		panic(lerr)
	}

	
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&status)
	if lerr := logger.CheckErr(err, errLocation, "Can`t unmarshal dust`s responce"); lerr != nil {
		panic(lerr)
	}

	if status.Status == http.StatusOK {
		return nil
	} else {
		return logger.NewErr(errLocation, strconv.Itoa(status.Status)+":"+status.Message)
	}
}
