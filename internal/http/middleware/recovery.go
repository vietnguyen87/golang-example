package middleware

import (
	"example-service/pkg/logger"
	"example-service/pkg/utils/apiwrapper"
	"github.com/gin-gonic/gin"
	"gitlab.marathon.edu.vn/pkg/go/xerrors"
	"runtime/debug"
)

func RecoverPanic(c *gin.Context) {
	log := logger.CToL(c.Request.Context(), "gin-recover")
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("method %v, path %v, err %v, trace %v",
				c.Request.Method,
				c.Request.URL.EscapedPath(),
				err,
				string(debug.Stack()))
			apiwrapper.Abort(c, &apiwrapper.Response{Error: xerrors.InternalServerError.New()})
		}
	}()
	c.Next()
}
