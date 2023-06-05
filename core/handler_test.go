package core_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	lambgo "github.com/gotools-io/lambgo/core"
	"github.com/gotools-io/lambgo/mocks"

	"github.com/labstack/echo/v4"
)

const request = "http://www.test.com/api/test"

func TestBasicCheck(t *testing.T) {
	w := httptest.NewRecorder()
	c := mocks.NewContextMock(http.MethodGet, "/health", "", "")
	h := lambgo.Handler{}
	h.BasicCheck(c)
	expected := httptest.NewRecorder()
	expected.WriteHeader(http.StatusOK)

	if !reflect.DeepEqual(w.Header(), expected.Header()) {
		t.Fail()
	}
}
func TestHandler_Create(t *testing.T) {
	type fields struct {
		Creator interface{ Create(c any) (any, error) }
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "Create handler OK",
			fields: fields{
				Creator: mocks.NewCreatorMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, `{"name":"Jhon Doe","address":"High St."}`, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusOK, ""),
		},
		{
			name: "Create handler - empty body request ",
			fields: fields{
				Creator: mocks.NewCreatorMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, ``, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusBadRequest, "Request body can't be empty"),
		},
		{
			name: "Create handler - creator function error ",
			fields: fields{
				Creator: mocks.NewCreatorMock(true, "creation error"),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, `{"name":"Jhon Doe","address":"High St."}`, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusInternalServerError, "creation error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lambgo.Handler{
				Creator: tt.fields.Creator,
			}
			result := h.Create(tt.args.c)
			if result != tt.want {
				t.Fail()
			}

		})
	}
}

func TestHandler_Read(t *testing.T) {
	type fields struct {
		Reader interface {
			Read(id any) (any, error)
			ReadAll(limit, marker any) ([]any, error)
		}
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "Reader - single read handler OK",
			fields: fields{
				Reader: mocks.NewReaderMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodGet, request, "", "/:id", "1234"),
			},
			want: mocks.NewContextMockResponse(http.StatusOK, ""),
		},
		{
			name: "Reader - single read handler - empty id param request ",
			fields: fields{
				Reader: mocks.NewReaderMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodGet, request, "", "/:id"),
			},
			want: mocks.NewContextMockResponse(http.StatusBadRequest, "id param not present"),
		},
		{
			name: "Reader - single read handler - reader function error ",
			fields: fields{
				Reader: mocks.NewReaderMock(false, "reading"),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodGet, request, "", "/:id", "1234"),
			},
			want: mocks.NewContextMockResponse(http.StatusInternalServerError, "reading error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lambgo.Handler{
				Reader: tt.fields.Reader,
			}
			result := h.Read(tt.args.c)
			if result != tt.want {
				t.Fail()
			}

		})
	}
}

func TestHandler_ReadAll(t *testing.T) {
	type fields struct {
		Reader interface {
			Read(id any) (any, error)
			ReadAll(limit, marker any) ([]any, error)
		}
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "Reader - single read handler OK",
			fields: fields{
				Reader: mocks.NewReaderMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodGet, request, "", ""),
			},
			want: mocks.NewContextMockResponse(http.StatusOK, ""),
		},
		{
			name: "Reader - single read handler - reader function error ",
			fields: fields{
				Reader: mocks.NewReaderMock(false, "reading"),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodGet, request, "", ""),
			},
			want: mocks.NewContextMockResponse(http.StatusInternalServerError, "reading error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lambgo.Handler{
				Reader: tt.fields.Reader,
			}
			result := h.ReadAll(tt.args.c)
			if result != tt.want {
				t.Fail()
			}

		})
	}
}

func TestHandler_Update(t *testing.T) {
	type fields struct {
		Updater interface{ Update(c any) (any, error) }
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "Update handler OK",
			fields: fields{
				Updater: mocks.NewUpdaterMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, `{"name":"Jhon Doe","address":"High St."}`, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusOK, ""),
		},
		{
			name: "Update handler - empty body request ",
			fields: fields{
				Updater: mocks.NewUpdaterMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, ``, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusBadRequest, "Request body can't be empty"),
		},
		{
			name: "Update handler - updater function error ",
			fields: fields{
				Updater: mocks.NewUpdaterMock(true, "update error"),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, `{"name":"Jhon Doe","address":"High St."}`, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusInternalServerError, "update error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lambgo.Handler{
				Updater: tt.fields.Updater,
			}
			result := h.Update(tt.args.c)
			if result != tt.want {
				t.Fail()
			}

		})
	}
}

func TestHandler_Delete(t *testing.T) {
	type fields struct {
		Deleter interface{ Delete(c any) error }
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "Delete handler OK",
			fields: fields{
				Deleter: mocks.NewDeleterMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodDelete, request, "", "/:id", "1234"),
			},
			want: mocks.NewContextMockResponse(http.StatusOK, ""),
		},
		{
			name: "Delete handler - no id param present on request ",
			fields: fields{
				Deleter: mocks.NewDeleterMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodDelete, request, "", ""),
			},
			want: mocks.NewContextMockResponse(http.StatusBadRequest, "Request body can't be empty"),
		},
		{
			name: "Delete handler - delete function error ",
			fields: fields{
				Deleter: mocks.NewDeleterMock(true, "delete error"),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodDelete, request, "", "/:id", "1234"),
			},
			want: mocks.NewContextMockResponse(http.StatusInternalServerError, "delete error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lambgo.Handler{
				Deleter: tt.fields.Deleter,
			}
			result := h.Delete(tt.args.c)
			if result != tt.want {
				t.Fail()
			}

		})
	}
}

func TestHandler_Execute(t *testing.T) {
	type fields struct {
		Executor interface {
			Execute(c any) (any, error)
			Action() string
		}
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "Execute handler OK",
			fields: fields{
				Executor: mocks.NewExecutorMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, `{"name":"Jhon Doe","address":"High St."}`, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusOK, ""),
		},
		{
			name: "Execute handler - empty body request ",
			fields: fields{
				Executor: mocks.NewExecutorMock(false, ""),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, ``, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusBadRequest, "Request body can't be empty"),
		},
		{
			name: "Executor handler - executor function error ",
			fields: fields{
				Executor: mocks.NewExecutorMock(true, "execution error"),
			},
			args: args{
				c: mocks.NewContextMock(http.MethodPost, request, `{"name":"Jhon Doe","address":"High St."}`, ""),
			},
			want: mocks.NewContextMockResponse(http.StatusInternalServerError, "execution error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lambgo.Handler{
				Executor: tt.fields.Executor,
			}
			result := h.Execute(tt.args.c)
			if result != tt.want {
				t.Fail()
			}

		})
	}
}
