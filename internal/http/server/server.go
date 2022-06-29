package server

import (
	"context"
	"example-service/docs"
	"example-service/pkg/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"example-service/internal/http/handler"
	"example-service/internal/http/middleware"
	"example-service/pkg/logger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start() error
}

func New(handler handler.Handler) Server {
	server := &serverImpl{
		handler: handler,
	}

	server.withRouter()

	return server
}

type serverImpl struct {
	handler handler.Handler

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
	router.Use(gin.Recovery())

	router.Use(middleware.Logger())

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
