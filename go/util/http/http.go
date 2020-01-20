package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Get(url string, data []byte, headers map[string]string) (*http.Response, []byte, error) {
	body := bytes.NewReader(data)
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	return resp, ret, err

}

func PostJson(url string, data []byte, headers map[string]string) (*http.Response, []byte, error) {
	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	return resp, ret, err
}