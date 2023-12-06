package utils

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

func HttpRequest(url, method string, body interface{}, resp interface{}) error {
	var err error
	var bodyBytes []byte
	var reader *bytes.Reader
	if body != nil {
		bodyBytes, err = jsoniter.Marshal(body)
		if err != nil {
			return err
		}
	}

	if bodyBytes != nil {
		reader = bytes.NewReader(bodyBytes)
	}

	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = jsoniter.Unmarshal(respBody, &resp)
	if err != nil {
		return err
	}
	return nil
}
