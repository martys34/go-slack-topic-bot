package boshio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/pkg/errors"
)

type Stemcell struct {
	Alias string
	Name  string
}

var _ message.Messager = &Stemcell{}

type boshioStemcellApiV1 []struct {
	Version string `json:"version"`
}

func (m Stemcell) Message() (string, error) {
	res, err := http.DefaultClient.Get(fmt.Sprintf("https://bosh.io/api/v1/stemcells/%s", m.Name))
	if err != nil {
		return "", err
	}

	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "reading response")
	}

	var data boshioStemcellApiV1

	err = json.Unmarshal(resBodyBytes, &data)
	if err != nil {
		return "", errors.Wrap(err, "unmarshalling")
	}

	if len(data) == 0 {
		return "", nil
	}

	return fmt.Sprintf("%s/%s", m.Alias, data[0].Version), nil
}
