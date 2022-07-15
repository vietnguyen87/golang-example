package logger

import (
	"context"
	"example-service/pkg/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gitlab.marathon.edu.vn/pkg/go/xcontext"
)

const (
	contextKey = "logger"
)

var (
	fieldMap = logrus.FieldMap{
		logrus.FieldKeyMsg: "message",
	}
)

func CToL(ctx context.Context, label string) *logrus.Entry {
	// CToL stands for Context-To-Log
	v := ctx.Value(contextKey)
	traceID := cast.ToString(ctx.Value(xcontext.KeyContextID.String()))
	if log, ok := v.(*logrus.Entry); ok {
		return log.WithField("label", label).WithField("request-id", fmt.Sprint(traceID))
	}

	return initLog(label).WithField("request-id", fmt.Sprint(traceID))
}

func LToC(parent context.Context, logger *logrus.Entry) context.Context {
	// LToC stands for Log-To-Context
	return context.WithValue(parent, contextKey, logger)
}

func initLog(label string) *logrus.Entry {
	cfg := config.GetLoggingConfig()

	logger := logrus.New()

	switch cfg.Formatter {
	case config.LoggingJSONFormatter:
		logger.SetFormatter(&logrus.JSONFormatter{FieldMap: fieldMap})
	case config.LoggingTextFormatter:
		logger.SetFormatter(&logrus.TextFormatter{FieldMap: fieldMap})
	}

	switch cfg.Level {
	case config.LoggingInfoLevel:
		logger.SetLevel(logrus.InfoLevel)
	case config.LoggingDebugLevel:
		logger.SetLevel(logrus.DebugLevel)
	}

	return logger.WithField("label", label)
}
