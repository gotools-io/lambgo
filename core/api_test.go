package core_test

import (
	"testing"

	lambgo "github.com/gotools-io/lambgo/core"
	"github.com/gotools-io/lambgo/mocks"
)

func TestNewAPI(t *testing.T) {
	type args struct {
		h    lambgo.Handler
		noun string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Create new POST - OK",
			args: args{
				h: lambgo.Handler{
					Creator: mocks.NewCreatorMock(false, ""),
				},
				noun: "test",
			},
			wantErr: false,
		},
		{
			name: "Create new POST - Empty handler should return error",
			args: args{
				h:    lambgo.Handler{},
				noun: "test",
			},
			wantErr: true,
		},
		{
			name: "Create new GET - OK",
			args: args{
				h: lambgo.Handler{
					Reader: mocks.NewReaderMock(false, ""),
				},
				noun: "test",
			},
			wantErr: false,
		},
		{
			name: "Create new PUT - OK",
			args: args{
				h: lambgo.Handler{
					Updater: mocks.NewUpdaterMock(false, ""),
				},
				noun: "test",
			},
			wantErr: false,
		},
		{
			name: "Create new DELETE - OK",
			args: args{
				h: lambgo.Handler{
					Deleter: mocks.NewDeleterMock(false, ""),
				},
				noun: "test",
			},
			wantErr: false,
		},
		{
			name: "Create new Execute - OK",
			args: args{
				h: lambgo.Handler{
					Executor: mocks.NewExecutorMock(false, ""),
				},
				noun: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lambgo.NewAPI(tt.args.h, tt.args.noun)
			if (err != nil) != tt.wantErr && got == nil {
				t.Fail()
			}
		})
	}
}

func Test_lambgoAPI_Check(t *testing.T) {
	type args struct {
		h    lambgo.Handler
		noun string
		host string
		port string
	}
	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{
		{
			name: "Launch an API - OK",
			args: args{
				h: lambgo.Handler{
					Creator: mocks.NewCreatorMock(false, ""),
				},
				noun: "test",
				host: "localhost",
				port: "8080",
			},
			wantErr: false,
		},
		{
			name: "Launch an API - Host not configured - should fail",
			args: args{
				h: lambgo.Handler{
					Creator: mocks.NewCreatorMock(false, ""),
				},
				noun: "test",
				host: "",
				port: "8080",
			},
			wantErr: true,
		},
		{
			name: "Launch an API - Port not configured - should fail",
			args: args{
				h: lambgo.Handler{
					Creator: mocks.NewCreatorMock(false, ""),
				},
				noun: "test",
				host: "localhost",
				port: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := lambgo.NewAPI(tt.args.h, tt.args.noun)
			if err != nil {
				t.Fail()
			}

			err = a.Check(tt.args.host, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Fail()
			}
		})
	}
}
