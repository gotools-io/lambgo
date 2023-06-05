package core

import (
	"errors"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	Creator interface {
		Create(c any) (any, error)
	}
	Reader interface {
		Read(id any) (any, error)
		ReadAll(limit, marker any) ([]any, error)
	}
	Updater interface {
		Update(r any) (any, error)
	}
	Deleter interface {
		Delete(id any) error
	}
	Executor interface {
		Execute(e any) (any, error)
		Action() string
	}
}

// BasicCheck - basic health check response
func (h Handler) BasicCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// Create - creater lambgo
func (h Handler) Create(c echo.Context) error {
	var lambgo any
	if err := c.Bind(&lambgo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	l, err := h.Creator.Create(lambgo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, l)
}

// Read - readr lambgo
func (h Handler) Read(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errors.New("id param not present"))
	}
	l, err := h.Reader.Read(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, l)
}

// ReadAll - read all lambgo
func (h Handler) ReadAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	marker := c.QueryParam("marker")

	l, err := h.Reader.ReadAll(limit, marker)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, l)
}

// Update - updater lambgo
func (h Handler) Update(c echo.Context) error {
	var lambgo any
	if err := c.Bind(&lambgo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	l, err := h.Updater.Update(lambgo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, l)
}

// Delete - deleter lambgo
func (h Handler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errors.New("id param not present"))
	}

	err := h.Deleter.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

// Executor - executor lambgo
func (h Handler) Execute(c echo.Context) error {
	var lambgo any
	if err := c.Bind(&lambgo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	l, err := h.Executor.Execute(lambgo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, l)
}
