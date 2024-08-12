package delivery

import "github.com/gin-gonic/gin"

type (
	IDelivery interface {
		HandleRequest(c *gin.Context)
	}

	IMiddleware interface {
		ValidateJwtToken() gin.HandlerFunc
	}
)
