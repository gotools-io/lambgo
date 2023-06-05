package core

import (
	"errors"
	"log"

	echo "github.com/labstack/echo/v4"
)

type lambgoAPI struct {
	router *echo.Echo
}

// Router defines all the endpoints of the app
func NewAPI(h Handler, noun string) (*lambgoAPI, error) {
	if (Handler{}) == h {
		return nil, errors.New("you didn't implement any handler")
	}

	lambgoAPI := lambgoAPI{
		router: echo.New(),
	}

	switch {
	case h.Creator != nil:
		lambgoAPI.router.POST("/api/"+noun, h.Create)
	case h.Reader != nil:
		lambgoAPI.router.GET("/api/"+noun, h.ReadAll)
		lambgoAPI.router.GET("/api/"+noun+"/:id", h.Read)
	case h.Updater != nil:
		lambgoAPI.router.PUT("/api/"+noun, h.Update)
	case h.Deleter != nil:
		lambgoAPI.router.DELETE("/api/"+noun+"/:id", h.Delete)
	case h.Executor != nil:
		lambgoAPI.router.POST("/api/"+noun+"/"+h.Executor.Action(), h.Execute)
	}

	lambgoAPI.router.GET("/health", h.BasicCheck)

	return &lambgoAPI, nil
}

func (a *lambgoAPI) Check(host, port string) error {
	switch {
	case host == "":
		return errors.New("empty host configuration")
	case port == "":
		return errors.New("empty port configuration")
	}
	return nil
}

func (a *lambgoAPI) Start(host, port string) error {
	log.Print("starting lambgo service \n")
	a.Check(host, port)
	log.Print("sucesfully started lambgo service \n")
	return a.router.Start(host + ":" + port)
}
