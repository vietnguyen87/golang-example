package handler

import (
	"example-service/dto"
	"example-service/internal/model/converter"
	"example-service/internal/repository"
	"example-service/pkg/logger"
	"github.com/gin-gonic/gin"
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

// @BasePath /v1

// GetTasks godoc
// @Summary GetTasks example
// @Schemes
// @Description Do GetTasks
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /tasks [get]
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
