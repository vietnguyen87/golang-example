package handler

import (
	"example-service/dto"
	"example-service/internal/constants"
	"example-service/internal/http/handler/builder"
	"example-service/internal/repository"
	"example-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type TaskHandler interface {
	Get(c *gin.Context)
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

// @BasePath  /v1

// Get godoc
// @Summary  GetTasks example
// @Schemes
// @Description  Do GetTasks
// @Tags         example
// @Accept       json
// @Produce      json
// @Param request query dto.ListReq true "body"
// @Failure      400  {object}  dto.ErrorResp
// @Failure      500  {object}  dto.ErrorResp
// @Success      200  {object}  dto.Results
// @Router       /tasks [get]
func (i *taskHandlerImpl) Get(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.CToL(ctx, "GetTasks")
	span, ctx := apm.StartSpan(c.Request.Context(), "GetTasks", "request")
	defer span.End()

	req := dto.ListReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.WithField("err", err).Errorf("Get returns error when ShouldBindJSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, builder.BuildErrorResponse(err, err.Error()))
		return
	}

	query := builder.BuildQuery(req.Q, req.SearchFields, builder.BuildFilters(req.Filters), builder.BuildSort(req.Sort), builder.BuildPagination(req.Pagination))
	query.SetSelectFields(req.SelectFields)
	query.SetHaveCount(req.HaveCount)

	tasks, total, err := i.repository.TaskRepository().Find(ctx, query)
	if err != nil {
		log.WithField("err", err).Errorf("TaskRepository().Find err: %s", err.Error())
		c.JSON(http.StatusInternalServerError, builder.BuildErrorResponse(err, constants.InternalError))
		return
	}
	c.JSON(http.StatusOK, builder.BuildTasksResponse(tasks, total, req.Pagination))
}
