package apiwrapper

import (
	"net/http"

	"gitlab.marathon.edu.vn/pkg/go/xcontext"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"gitlab.marathon.edu.vn/pkg/go/xerrors"
)

const (
	statusSuccess = 1
	statusFail    = 0
)

/**
Example success response:
{
    "code": 1,
    "message": "Thành công!",
    "data": {

    },
    "trace_id": ""
}

Example failure response:
{
    "code": -1,
    "message": "Yêu cầu không hợp lệ!",
	"result":{
		"data":{}
	}
    "trace_id": ""
}
*/
type Response struct {
	Error      error       `json:"error"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}
type APIResponse struct {
	//Status  int64            `json:"status"`
	Code    xerrors.ErrorType `json:"code"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data,omitempty"`
	//Error   interface{}      `json:"error,omitempty"`
	TraceID string `json:"trace_id"`
}
type apiHandlerFn func(c *gin.Context) *Response

func Wrap(fn apiHandlerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		rsp := fn(c)
		apiRsp := transformAPIResponse(c, rsp)
		if rsp.StatusCode == 0 {
			c.JSON(http.StatusOK, apiRsp)
			return
		}
		c.JSON(rsp.StatusCode, apiRsp)
	}
}

func Abort(c *gin.Context, rsp *Response) {
	apiRsp := transformAPIResponse(c, rsp)
	if rsp.StatusCode == 0 {
		c.AbortWithStatusJSON(http.StatusOK, apiRsp)
		return
	}
	c.AbortWithStatusJSON(rsp.StatusCode, apiRsp)
}

func AbortWithStatus(c *gin.Context, statusCode int, rsp *Response) {
	apiRsp := transformAPIResponse(c, rsp)
	c.AbortWithStatusJSON(statusCode, apiRsp)
}

func transformAPIResponse(c *gin.Context, rsp *Response) *APIResponse {
	traceID := cast.ToString(c.Request.Context().Value(xcontext.KeyContextID.String()))
	if xerrors.Is(rsp.Error, xerrors.Success) {
		apiRsp := &APIResponse{
			//Status:  statusSuccess,
			Code:    xerrors.GetErrorType(rsp.Error),
			Message: xerrors.GetMessage(rsp.Error),
			Data:    rsp.Data,
			TraceID: traceID,
		}
		return apiRsp
	}
	apiRsp := &APIResponse{
		//Status:  statusFail,
		Code:    xerrors.GetErrorType(rsp.Error),
		Message: xerrors.GetMessage(rsp.Error),
		//Error:   rsp.Error,
		Data:    rsp.Data,
		TraceID: traceID,
	}
	_ = c.Error(rsp.Error)
	return apiRsp
}
