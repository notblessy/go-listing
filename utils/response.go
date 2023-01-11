package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseSuccess struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ResponseInternalServerError(c echo.Context, response *ResponseError) error {
	return c.JSON(http.StatusInternalServerError, response)
}

func ResponseBadRequest(c echo.Context, response *ResponseError) error {
	defaultValueError(response)
	return c.JSON(http.StatusBadRequest, response)
}

func ResponseCreated(c echo.Context, response *ResponseSuccess) error {
	defaultValueSuccess(response)
	return c.JSON(http.StatusCreated, response)
}

func ResponseOK(c echo.Context, response *ResponseSuccess) error {
	defaultValueSuccess(response)
	return c.JSON(http.StatusOK, response)
}

func defaultValueSuccess(response *ResponseSuccess) {
	response.Success = true
}

func defaultValueError(response *ResponseError) {
	response.Success = false
	if response.Message == "" {
		response.Message = "ERROR"
	}
}
