package core

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/labstack/gommon/log"
)

// Client - basic http client
type (
	client struct {
		APIURL   string
		Headers  http.Header
		Executor Executor
	}
	Executor interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

func NewRestClient(url string, headers http.Header, exec Executor) client {
	return client{
		APIURL:   url,
		Headers:  headers,
		Executor: exec,
	}
}

func (c client) CallService(path, method string, reqBody []byte) ([]byte, error) {
	log.Print("calling url ", c.APIURL+path)
	req, _ := http.NewRequest(method, c.APIURL+path, bytes.NewBuffer(reqBody))
	req.Header = c.Headers

	resp, err := c.Executor.Do(req)
	if err != nil {
		log.Errorf("error calling url %s - error: %s", c.APIURL+path, err)
		return nil, err
	}
	defer resp.Body.Close()
	resBody, _ := ioutil.ReadAll(resp.Body)

	return resBody, nil
}
