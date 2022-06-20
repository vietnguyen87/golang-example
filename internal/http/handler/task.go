package handler

import (
	"github.com/gin-gonic/gin"
	"mapi-service/dto"
	"mapi-service/internal/model/converter"
	"mapi-service/internal/repository"
	"mapi-service/pkg/logger"
	"net/http"
)

type TaskHandler interface {
	GetTasks(c *gin.Context)
}

type taskHandlerImpl struct {
	repository repository.Repository
}

func NewTaskHandler(
	repository repository.Repository,
) TaskHandler {
	return &taskHandlerImpl{
		repository: repository,
	}
}

func (i *taskHandlerImpl) GetTasks(c *gin.Context) {
	log := logger.CToL(c.Request.Context(), "GetTasks")

	tasks, err := i.repository.TaskRepository().GetTasks(
		logger.LToC(c.Request.Context(), log),
	)
	if err != nil {
		log.WithField("err", err).Errorf("GetTasks returns error: %s", err.Error())
		abortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &dto.GetTasksResponse{
		Data: converter.TasksToDTO(tasks),
	})
}
