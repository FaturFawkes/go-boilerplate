package gin_engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// setupRoutes internal method to set up routes if any
func (engine *GinEngine) SetupRoutes(routes []Route, group *gin.RouterGroup) {
	if routes == nil || len(routes) == 0 {
		engine.logger.Info("nothing to setup")
		return
	}

	for _, pRoute := range routes {
		if pRoute.Middleware != nil {
			group.Use(pRoute.Middleware...)
		}

		if pRoute.Action != nil {
			if pRoute.Method == http.MethodGet {
				group.GET(pRoute.Path, pRoute.Action)
			}

			if pRoute.Method == http.MethodPost {
				group.POST(pRoute.Path, pRoute.Action)
			}

			if pRoute.Method == http.MethodPut {
				group.PUT(pRoute.Path, pRoute.Action)
			}

			if pRoute.Method == http.MethodDelete {
				group.DELETE(pRoute.Path, pRoute.Action)
			}

			if pRoute.Method == http.MethodPatch {
				group.PATCH(pRoute.Path, pRoute.Action)
			}

			if pRoute.Method == http.MethodHead {
				group.HEAD(pRoute.Path, pRoute.Action)
			}

			if pRoute.Method == http.MethodOptions {
				group.OPTIONS(pRoute.Path, pRoute.Action)
			}
		}

		if pRoute.Routes != nil {
			engine.SetupRoutes(pRoute.Routes, group.Group(pRoute.Path))
		}
	}
}
