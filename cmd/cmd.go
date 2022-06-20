package cmd

import (
	"log"
	"mapi-service/pkg/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "academic core services",
	Long:  `A service for handle Teacher, Deal, Student...`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	config.Load()

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(configCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
