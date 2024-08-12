package mapper

import (
	"golang_boilerplate/internal/delivery/response"
	"golang_boilerplate/pkg/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Exception struct{}
)

func NewException() Exception {
	return Exception{}
}

func (Exception) ExceptionToJsonResponse(ctx *gin.Context, err error) {
	switch err.(type) {
	case exception.BadRequest, *exception.BadRequest:
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "INVALID_REQUEST",
		})
		return

	case exception.NotFound, *exception.NotFound:
		ctx.JSON(http.StatusNotFound, response.Response{
			Success: false,
			Message: "RESOURCE_NOT_FOUND",
		})
		return

	case exception.Unauthorized, *exception.Unauthorized:
		ctx.JSON(http.StatusUnauthorized, response.Response{
			Success: false,
			Message: "NOT_AUTHORIZED",
		})
		return

	case exception.ServiceUnavailable, *exception.ServiceUnavailable:
		ctx.JSON(http.StatusServiceUnavailable, response.Response{
			Success: false,
			Message: "SERVICE_UNAVAILABLE",
		})
		return

	default:
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "INTERNAL_SERVER_ERROR",
		})
		return
	}
}
