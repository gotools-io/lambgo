package mocks

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type (
	creatorMock struct {
		fail         bool
		errorMessage string
	}
	readerMock struct {
		fail         bool
		errorMessage string
	}
	updaterMock struct {
		fail         bool
		errorMessage string
	}
	deleterMock struct {
		fail         bool
		errorMessage string
	}
	executorMock struct {
		fail         bool
		errorMessage string
		action       string
	}
)

func NewCreatorMock(fail bool, errorMessage string) creatorMock {
	return creatorMock{
		fail:         fail,
		errorMessage: errorMessage,
	}
}

func NewReaderMock(fail bool, errorMessage string) readerMock {
	return readerMock{
		fail:         fail,
		errorMessage: errorMessage,
	}
}

func NewUpdaterMock(fail bool, errorMessage string) updaterMock {
	return updaterMock{
		fail:         fail,
		errorMessage: errorMessage,
	}
}

func NewDeleterMock(fail bool, errorMessage string) deleterMock {
	return deleterMock{
		fail:         fail,
		errorMessage: errorMessage,
	}
}

func NewExecutorMock(fail bool, errorMessage string) executorMock {
	return executorMock{
		fail:         fail,
		errorMessage: errorMessage,
	}
}

func (c creatorMock) Create(r any) (any, error) {
	if c.fail {
		return "", errors.New(c.errorMessage)
	}
	return "it works", nil
}

func (c readerMock) Read(r any) (any, error) {
	if c.fail {
		return "", errors.New(c.errorMessage)
	}
	return "it works", nil
}

func (c readerMock) ReadAll(limit, marker any) ([]any, error) {
	if c.fail {
		return []any{"it works", "it works"}, errors.New(c.errorMessage)
	}
	return []any{"it works", "it works"}, nil
}

func (c updaterMock) Update(r any) (any, error) {
	if c.fail {
		return "", errors.New(c.errorMessage)
	}
	return "it works", nil
}
func (c deleterMock) Delete(id any) error {
	if c.fail {
		return errors.New(c.errorMessage)
	}
	return nil
}

func (c executorMock) Execute(e any) (any, error) {
	if c.fail {
		return "", errors.New(c.errorMessage)
	}
	return "it works", nil
}
func (c executorMock) Action() string {
	return c.action
}

func NewContextMock(method, endpoint, body, path string, values ...string) echo.Context {
	r := httptest.NewRequest(method, endpoint, bytes.NewReader([]byte(body)))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(r, w)
	c.SetPath(path)
	c.SetParamNames("id")
	c.SetParamValues(values...)
	c.Request().Body = r.Body
	return c
}

func NewContextMockResponse(statusCode int, errorMessage string) error {
	r := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(r, w)
	return c.JSON(statusCode, errors.New(errorMessage))
}

type mockHTTP struct {
	Fail bool
}

func NewMockHTTP(fail bool) mockHTTP {
	return mockHTTP{
		Fail: fail,
	}
}

func NewMockHTTPResponse() *http.Response {
	json := `{"result":"ok'"}`
	r := io.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		Body: r,
	}
}

func (m mockHTTP) Do(req *http.Request) (*http.Response, error) {
	if m.Fail {
		return nil, errors.New("some error happened while calling the REST API")
	}

	return NewMockHTTPResponse(), nil
}
