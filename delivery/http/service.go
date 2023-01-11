package http

import (
	"github.com/labstack/echo/v4"
	"github.com/notblessy/go-listing/model"
)

// HTTPService :nodoc:
type HTTPService struct {
	productUsecase model.ProductUsecase
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// RegisterProductUsecase :nodoc:
func (h *HTTPService) RegisterProductUsecase(p model.ProductUsecase) {
	h.productUsecase = p
}

// Routes :nodoc:
func (h *HTTPService) Routes(route *echo.Echo) {
	route.POST("/products", h.createProductHandler)
	route.GET("/products", h.findAllProductHandler)
}
