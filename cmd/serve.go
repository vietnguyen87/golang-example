package cmd

import (
	"context"
	"example-service/internal/http/handler"
	"example-service/internal/http/server"
	"example-service/internal/repository"
	"example-service/pkg/config"
	"example-service/pkg/gormclient"
	"example-service/pkg/logger"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.marathon.edu.vn/pkg/go/xprom"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server on the predefined port",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CToL(context.Background(), "serveCmd")

		s := setupServer()
		if err := s.Start(); err != nil {
			log.WithError(err).Fatalf("Failed to start HTTP server with error: %s", err.Error())
		}
	},
}

func setupServer() server.Server {
	log := logger.CToL(context.Background(), "setupServer")

	db, err := gormclient.New()
	if err != nil {
		log.WithError(err).Fatalf("gormclient.New returns error: %s", err.Error())
	}
	measurer := xprom.New(
		xprom.Namespace("mrt"),
		xprom.ListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort)),
		xprom.Ignore("/metrics", "/ping", "/swagger/*any"),
		xprom.Subsystem("example"),
	)
	repositoryService := repository.New(db)
	return server.New(
		handler.New(repositoryService),
		measurer,
	)
}
