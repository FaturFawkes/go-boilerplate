package gin_engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck will be used as default routing on each service it supposed to be
// handle docker to check service is up or not
func (engine *GinEngine) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "")
	return
}
