package core_test

import (
	"net/http"
	"testing"

	lambgo "github.com/gotools-io/lambgo/core"
	"github.com/gotools-io/lambgo/mocks"
)

func Test_client_CallService(t *testing.T) {
	type fields struct {
		APIURL   string
		Headers  http.Header
		Executor lambgo.Executor
	}
	type args struct {
		path    string
		method  string
		reqBody []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Call service ok",
			fields: fields{
				APIURL:   "https://testurl.com/testapi",
				Headers:  http.Header{},
				Executor: mocks.NewMockHTTP(false),
			},
			args: args{
				path:    "/testpath",
				method:  http.MethodGet,
				reqBody: nil,
			},
			wantErr: false,
		},
		{
			name: "Call service failure",
			fields: fields{
				APIURL:   "https://testurl.com/testapi",
				Headers:  http.Header{},
				Executor: mocks.NewMockHTTP(true),
			},
			args: args{
				path:    "/testpath",
				method:  http.MethodGet,
				reqBody: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := lambgo.NewRestClient(
				tt.fields.APIURL,
				tt.fields.Headers,
				tt.fields.Executor,
			)
			_, err := c.CallService(tt.args.path, tt.args.method, tt.args.reqBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CallService() error = %v, wantErr %v", err, tt.wantErr)

			}
		})
	}
}
