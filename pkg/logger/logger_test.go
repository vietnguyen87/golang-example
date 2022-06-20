package logger

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var _ = Describe("Logger", func() {
	It("should returns", func() {
		viper.Set("app.env", "local")

		ctxWithoutLog := context.TODO()

		log := CToL(ctxWithoutLog, "test")
		Expect(log.Logger.Formatter).To(BeEquivalentTo(&logrus.TextFormatter{FieldMap: fieldMap}))

		log.Logger.SetFormatter(&logrus.JSONFormatter{})
		ctxWithLog := LToC(ctxWithoutLog, log)

		logFromContext := CToL(ctxWithLog, "test")
		Expect(logFromContext.Logger.Formatter).To(BeEquivalentTo(&logrus.JSONFormatter{}))
	})
})
