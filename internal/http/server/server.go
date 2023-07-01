package server

import (
	"context"
	"example-service/docs"
	"example-service/pkg/config"
	"example-service/pkg/tracer"
	"fmt"
	"gitlab.marathon.edu.vn/pkg/go/xprom"
	"net/http"

	"example-service/internal/http/handler"
	"example-service/internal/http/middleware"
	"example-service/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.marathon.edu.vn/pkg/go/xerrors"
)

type Server interface {
	Start() error
}

func New(handler handler.Handler, measurer xprom.Measurer) Server {
	server := &serverImpl{
		handler:  handler,
		measurer: measurer,
	}

	server.withRouter()

	return server
}

type serverImpl struct {
	handler  handler.Handler
	measurer xprom.Measurer

	router *gin.Engine
}

func (i *serverImpl) Start() error {
	log := logger.CToL(context.Background(), "server.Start")
	log.Infof("Listening and serving HTTP on :%d", config.GetHTTPConfig().Port)

	return i.router.Run(fmt.Sprintf(":%d", config.GetHTTPConfig().Port))
}

func (i *serverImpl) withRouter() {
	if config.GetAppConfig().Env != "local" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(middleware.Logger())
	//Prometheus Include Recovery
	i.measurer.GinMiddleware(router)
	//Tracer middleware
	router.Use(tracer.Middleware(router))

	_ = xerrors.Initialize()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})
	docs.SwaggerInfo.BasePath = "/v1"
	v1 := router.Group("/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", i.handler.TaskHandler().Get)
			//tasks.GET("/tasks/:id", i.handler.TaskHandler().GetOne)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	i.router = router
}
