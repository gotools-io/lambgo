package main

import (
	"log"
	"os"

	lambgo "github.com/gotools-io/lambgo/core"
)

type creator struct{}

func (c creator) Create(r any) (any, error) {
	//do something with your interface
	//maybe store in the Database
	return "Something is created", nil
}

// example run
func main() {
	handler := lambgo.Handler{
		Creator: creator{},
	}
	api, err := lambgo.NewAPI(handler, "test")
	if err != nil {
		log.Fatalf("unable to configure lambgo - err: %v", err.Error())
	}
	err = api.Start(os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("unable to start lambgo - err: %v", err.Error())
	}
}
