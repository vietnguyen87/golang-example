package handler

import (
	"example-service/dto"
	"example-service/internal/helper"
	"example-service/internal/model/converter"
	"example-service/internal/repository"
	"example-service/pkg/logger"
	"example-service/pkg/utils/apiwrapper"
	"github.com/gin-gonic/gin"
	"gitlab.marathon.edu.vn/pkg/go/xerrors"
	"go.elastic.co/apm/v2"
)

type TaskHandler interface {
	Get(c *gin.Context) *apiwrapper.Response
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

// Get godoc
// @Summary GetTasks example
// @Schemes
// @Description Do GetTasks
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /tasks [get]
func (i *taskHandlerImpl) Get(c *gin.Context) *apiwrapper.Response {
	ctx := c.Request.Context()
	log := logger.CToL(ctx, "GetTasks")
	span, ctx := apm.StartSpan(c.Request.Context(), "GetTasks", "request")
	defer span.End()

	// create new ao.Span and context.Context for this part of the request
	//return &apiwrapper.Response{Error: errors.BadRequestErr.Report(errors.New("test lỗi"))}
	body := dto.GetReq{}
	if err := c.ShouldBind(&body); err != nil {
		log.WithField("err", err).Errorf("Get returns error when ShouldBindJSON: %s", err.Error())
		return &apiwrapper.Response{Error: xerrors.BadRequestErr.Report(err)}
	}
	query := helper.BuildQuery("", nil, nil, helper.BuildPagination(body.Page, body.Limit))
	tasks, total, err := i.repository.TaskRepository().Find(
		logger.LToC(ctx, log), query,
	)
	if err != nil {
		log.WithField("err", err).Errorf("GetTasks returns error: %s", err.Error())
		return &apiwrapper.Response{Error: xerrors.BadRequestErr.Report(err)}
	}
	return apiwrapper.SuccessWithDataResponse(&dto.GetResponse{
		Tasks:      converter.TasksToDTO(tasks),
		Total:      total,
		Pagination: dto.SetPagination(body.Page, body.Limit),
	})
}
