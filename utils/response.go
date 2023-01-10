package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(c echo.Context, response *Response) error {
	return c.JSON(http.StatusInternalServerError, response)
}

func ResponseBadRequest(c echo.Context, response *Response) error {
	defaultValueError(response)
	return c.JSON(http.StatusBadRequest, response)
}

func ResponseCreated(c echo.Context, response *Response) error {
	defaultValueSuccess(response)
	return c.JSON(http.StatusCreated, response)
}

func ResponseOK(c echo.Context, response *Response) error {
	defaultValueSuccess(response)
	return c.JSON(http.StatusOK, response)
}

func defaultValueSuccess(response *Response) {
	response.Success = true
	if response.Message == "" {
		response.Message = "SUCCESS"
	}
}

func defaultValueError(response *Response) {
	response.Success = false
	response.Data = nil
	if response.Message == "" {
		response.Message = "ERROR"
	}
}
