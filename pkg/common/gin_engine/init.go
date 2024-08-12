package gin_engine

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"time"
)

type (
	IGinEngine interface {
		// Start this method should be used to start service
		Start()

		// AddRoute is used for adding new route with existing method
		AddRoute(route Route)

		// HealthCheck will be used as default routing on each service it supposed to be
		// handle docker to check service is up or not
		HealthCheck(c *gin.Context)

		// setupRoutes internal method to set up routes if any
		SetupRoutes(routes []Route, group *gin.RouterGroup)

		// Response send json response with desired format
		Response(c *gin.Context, httpCode int, success bool, message string, data interface{}, meta interface{})
	}

	Route struct {
		Path   string
		Method string

		Action     func(c *gin.Context)
		Routes     []Route
		Middleware []gin.HandlerFunc
	}

	GinEngine struct {
		ctx    context.Context
		cancel context.CancelFunc

		logger     *zap.Logger
		engine     *gin.Engine
		routes     []Route
		quitSignal chan os.Signal

		host        string
		port        string
		environment string
		timeout     time.Duration

		disableDefaultHealthCheck bool
	}

	Constructor struct {
		Ctx    context.Context
		Logger *zap.Logger

		Host        string
		Port        string
		Environment string
		Timeout     time.Duration

		DisableDefaultHealthCheck bool
	}
)

// New will init gin class with default configuration and it usable
// with another service
func New(cs Constructor) (IGinEngine, error) {
	engine := new(GinEngine)
	ctx, cancel := context.WithCancel(cs.Ctx)

	engine.logger = cs.Logger
	engine.ctx = ctx
	engine.cancel = cancel
	engine.disableDefaultHealthCheck = cs.DisableDefaultHealthCheck
	engine.environment = cs.Environment

	engine.timeout = 10 * time.Second
	if cs.Timeout > 0*time.Second {
		engine.timeout = cs.Timeout
	}

	// set application host, in this case host will be optional
	engine.host = cs.Host

	// adding validation currently we just need to verify port only
	engine.port = cs.Port
	if engine.port == "" {
		return nil, errors.New("service port must be set")
	}

	return engine, nil
}
