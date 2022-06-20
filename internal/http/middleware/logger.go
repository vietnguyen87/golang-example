package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"example-service/pkg/logger"
)

func Logger() gin.HandlerFunc {
	return LoggerMiddleware{}.HandlerFunc
}

type LoggerMiddleware struct{}

func (i LoggerMiddleware) HandlerFunc(c *gin.Context) {
	log := logger.CToL(c.Request.Context(), "LoggerMiddleware")

	start := time.Now()

	c.Next()

	param := i.prepareParam(c, start)
	msg := i.prepareMessage(param)

	if param.StatusCode < 400 {
		i.injectFields(log, param).Info(msg)
	} else if param.StatusCode < 500 {
		i.injectFields(log, param).Warn(msg)
	} else {
		i.injectFields(log, param).Error(msg)
	}
}

func (i LoggerMiddleware) prepareParam(c *gin.Context, start time.Time) gin.LogFormatterParams {
	param := gin.LogFormatterParams{
		Request: c.Request,
		Keys:    c.Keys,
	}

	param.TimeStamp = time.Now()
	param.Latency = param.TimeStamp.Sub(start)

	param.ClientIP = c.ClientIP()
	param.Method = c.Request.Method
	param.StatusCode = c.Writer.Status()
	param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

	param.BodySize = c.Writer.Size()

	path := c.Request.URL.Path
	if c.Request.URL.RawQuery != "" {
		path += "?" + c.Request.URL.RawQuery
	}

	param.Path = path

	return param
}

func (i LoggerMiddleware) prepareMessage(param gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"[GIN] %d | %dms | %s | %s %v",
		param.StatusCode,
		param.Latency.Milliseconds(),
		param.ClientIP,
		param.Method,
		param.Path,
	)
}

func (i LoggerMiddleware) injectFields(log *logrus.Entry, param gin.LogFormatterParams) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"code":          param.StatusCode,
		"latency":       param.Latency.Milliseconds(),
		"client_ip":     param.ClientIP,
		"method":        param.Method,
		"path":          param.Path,
		"error_message": param.ErrorMessage,
	})
}
