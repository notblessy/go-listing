package http

import (
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
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: err.Error(),
		})
	}

	// if err := c.Validate(&data); err != nil {
	// 	logger.Error(ErrBadRequest)
	// 	return utils.ResponseBadRequest(c, &utils.Response{
	// 		Message: fmt.Sprintf("error validate: %s", ErrBadRequest),
	// 		Data:    ErrBadRequest,
	// 	})
	// }

	id, err := h.productUsecase.Create(&data)
	if err != nil {
		logger.Error(err)
		return utils.ResponseError(c, &utils.Response{
			Message: err.Error(),
		})
	}

	return utils.ResponseOK(c, &utils.Response{
		Data: id,
	})
}
