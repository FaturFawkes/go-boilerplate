package delivery

import (
	"golang_boilerplate/domain/delivery"
	"golang_boilerplate/pkg/common/gin_engine"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouting(engine gin_engine.IGinEngine, ctrl delivery.IDelivery, middleware delivery.IMiddleware) {
	engine.AddRoute(gin_engine.Route{
		Path: "/orchestrator",
		Routes: []gin_engine.Route{
			{
				Path:   "/*any",
				Method: http.MethodGet,
				Action: ctrl.HandleRequest,
				Middleware: []gin.HandlerFunc{
					middleware.ValidateJwtToken(),
				},
			},
		},
	})
}
