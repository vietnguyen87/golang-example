package server

import (
	"context"
	"fmt"
	"mapi-service/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"

	"mapi-service/internal/http/handler"
	"mapi-service/internal/http/middleware"
	"mapi-service/pkg/logger"
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

	router.GET("/api/v1/tasks", i.handler.TaskHandler().GetTasks)
	//router.POST("/api/v1/tasks", i.handler.TaskHandler().CreateTask)
	//router.PATCH("/api/v1/tasks/:id/completed", i.handler.TaskHandler().MarkTaskAsCompleted)
	//router.DELETE("/api/v1/tasks/:id", i.handler.TaskHandler().DeleteTask)

	i.router = router
}
