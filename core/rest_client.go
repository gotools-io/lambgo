package core

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/gommon/log"
)

// Client - basic http client
type (
	client struct {
		apiURL   string
		headers  http.Header
		executor Executor
	}
	Executor interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

func NewRestClient(url string, headers http.Header, exec Executor) client {
	return client{
		apiURL:   url,
		headers:  headers,
		executor: exec,
	}
}

func (c client) CallService(path, method string, reqBody []byte) ([]byte, error) {
	log.Print("calling url ", c.apiURL+path)
	req, _ := http.NewRequest(method, c.apiURL+path, bytes.NewBuffer(reqBody))
	req.Header = c.headers

	resp, err := c.executor.Do(req)
	if err != nil {
		log.Errorf("error calling url %s - error: %s", c.apiURL+path, err)
		return nil, err
	}
	defer resp.Body.Close()
	resBody, _ := io.ReadAll(resp.Body)

	return resBody, nil
}
