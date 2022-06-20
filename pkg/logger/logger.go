package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"mapi-service/pkg/config"
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
	if v == nil {
		return initLog(label)
	}

	if log, ok := v.(*logrus.Entry); ok {
		log = log.WithField("label", label)
		return log
	}

	return initLog(label)
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
