package httputils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Get(url string) (r *http.Request, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func Post(url string, body *interface{}) (r *http.Request, err error) {
	responseBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(responseBody)
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return nil, err
	}

	return req, nil
}
