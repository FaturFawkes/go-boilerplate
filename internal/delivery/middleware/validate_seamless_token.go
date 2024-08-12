package middleware

import (
	"golang_boilerplate/internal/delivery/response"
	"golang_boilerplate/pkg/common"
	"golang_boilerplate/pkg/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (mw *Middleware) ValidateJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get(config.EnvHeaderAuthorization)
		if auth == "" {
			c.JSON(http.StatusUnauthorized, response.DefaultResponse(false, "please provide authorization token", nil))
			c.Abort()
			return
		}

		if value := strings.Split(auth, " "); len(value) != 2 {
			c.JSON(http.StatusUnauthorized, response.DefaultResponse(false, "invalid token provided", nil))
			c.Abort()
			return
		} else {
			auth = value[1]
		}

		token, err := common.DecodeTokenWithoutValidate(auth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.DefaultResponse(false, "cannot validate your account", nil))
			c.Abort()
			return
		}

		c.Set("userId", token.UserId)
		c.Set("roleId", token.RoleId)
		c.Set("email", token.Email)

		c.Next()
	}
}
