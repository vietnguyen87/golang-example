package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"mapi-service/internal/http/handler"
	"mapi-service/internal/http/server"
	"mapi-service/internal/repository"
	"mapi-service/pkg/gormclient"
	"mapi-service/pkg/logger"
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

	repositoryService := repository.New(db)
	return server.New(
		handler.New(repositoryService),
	)
}
