package gin_engine

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
)

// Start this method should be used to start service
func (engine *GinEngine) Start() {
	gin.SetMode((func(env string) string {
		switch env {
		case "production", "prod":
			return gin.ReleaseMode
		case "testing", "test":
			return gin.TestMode
		default:
			return gin.DebugMode
		}
	})(engine.environment))
	engine.engine = gin.New()
	engine.logger.Info("setting up routes ...", zap.Any("routeCount", len(engine.routes)))

	if !engine.disableDefaultHealthCheck {
		engine.logger.Info("initializing default health check system")
		engine.engine.GET("/__health", engine.HealthCheck)
	}

	defer engine.cancel()
	server := &http.Server{Addr: fmt.Sprintf("%s:%s", engine.host, engine.port), Handler: engine.engine}
	engine.SetupRoutes(engine.routes, engine.engine.Group("/"))
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			engine.logger.Error("error while starting service", zap.Error(err))
			return
		}
	}()

	engine.logger.Info(fmt.Sprintf("service listening on => %s:%s", engine.host, engine.port))
	engine.quitSignal = make(chan os.Signal, 1)
	signal.Notify(engine.quitSignal, os.Interrupt)

	<-engine.quitSignal
	engine.logger.Warn("attempting to gracefully shutdown service")

	engine.cancel()
	if err := server.Shutdown(engine.ctx); err != nil {
		engine.logger.Info("error while attempting graceful shutdown", zap.Error(err))
		os.Exit(0)
	}

	select {
	case <-engine.ctx.Done():
		engine.logger.Info("timeout of 5 seconds.")
	}

	engine.logger.Info("Server exiting")
}
