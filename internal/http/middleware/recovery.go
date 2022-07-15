package middleware

import (
	"example-service/pkg/errors"
	"example-service/pkg/logger"
	"example-service/pkg/utils/apiwrapper"
	"github.com/gin-gonic/gin"
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
			apiwrapper.Abort(c, &apiwrapper.Response{Error: errors.InternalServerError.New()})
		}
	}()
	c.Next()
}
