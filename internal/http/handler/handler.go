package handler

import (
	"errors"
	"mapi-service/dto"
	"mapi-service/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	//"gitlab.marathon.edu.vn/marathon/services/mapiinternal/repository"
)

//go:generate mockery --name=Handler --case=snake
type Handler interface {
	TaskHandler() TaskHandler
}

func New(
	repository repository.Repository,
) Handler {
	return &handlerImpl{
		taskHandler: NewTaskHandler(repository),
	}
}

type handlerImpl struct {
	taskHandler TaskHandler
}

func (i *handlerImpl) TaskHandler() TaskHandler {
	return i.taskHandler
}

func abortWithError(c *gin.Context, code int, args ...error) {
	err := errors.New(strconv.Itoa(code))
	if len(args) == 1 {
		err = args[0]
	}

	c.JSON(code, &dto.ErrorResponse{Error: err.Error()})
}
