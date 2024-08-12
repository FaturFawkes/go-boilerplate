package gin_engine

import "github.com/gin-gonic/gin"

// Response send json response with desired format
func (engine *GinEngine) Response(c *gin.Context, httpCode int, success bool, message string, data interface{},
	meta interface{}) {
	c.JSON(httpCode, struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta,omitempty"`
	}{
		Status:  success,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
	return
}
