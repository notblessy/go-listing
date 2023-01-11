package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/notblessy/go-listing/model"
	"github.com/notblessy/go-listing/utils"

	"github.com/sirupsen/logrus"
)

// createProductHandler :nodoc:
func (h *HTTPService) createProductHandler(c echo.Context) error {
	logger := logrus.WithField("context", utils.Dump(c))
	var data model.Product

	if err := c.Bind(&data); err != nil {
		logger.Error(err)
		return utils.ResponseBadRequest(c, &utils.ResponseError{
			Message: err.Error(),
		})
	}

	if err := c.Validate(&data); err != nil {
		logger.Error(err)
		return utils.ResponseBadRequest(c, &utils.ResponseError{
			Message: fmt.Sprintf("error validate: %s", err),
		})
	}

	id, err := h.productUsecase.Create(&data)
	if err != nil {
		logger.Error(err)
		return utils.ResponseInternalServerError(c, &utils.ResponseError{
			Message: err.Error(),
		})
	}

	return utils.ResponseOK(c, &utils.ResponseSuccess{
		Data: id,
	})
}

func (h *HTTPService) findAllProductHandler(c echo.Context) error {
	logger := logrus.WithField("context", utils.Dump(c))

	req := model.ProductQuery{
		Sort: c.QueryParam("sort"),
	}

	products, err := h.productUsecase.FindAll(&req)
	if err != nil {
		logger.Error(err)
		return utils.ResponseInternalServerError(c, &utils.ResponseError{
			Message: fmt.Sprintf("%s", err),
		})
	}

	return utils.ResponseOK(c, &utils.ResponseSuccess{
		Data: products,
	})
}
