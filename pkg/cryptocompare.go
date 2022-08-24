package pkg

import (
	"bytes"
	"net/http"
)

type CryptoCompare interface {
	create(url string, protocol string, method string, token string)
}

type Requestor struct {
	option_fsyms string
	option_fsymf string
	option_app   string
	request      *http.Request
}

func (api *Requestor) create(url string, protocol string, method string, token string, body byte) error {
	jsonBody := []byte{body}
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		api.request = req
		api.request.Header.Add("Content-Type", "application/json")
		api.request.Header.Add("Authorization", "ApiKey "+token)
	}
	return nil
}
