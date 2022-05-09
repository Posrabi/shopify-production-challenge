package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func New() *cobra.Command {
	app := &cobra.Command{
		Use:   "inv",
		Short: "inventory tracking service",
	}

	app.AddCommand(newServerCmd())

	return app
}
