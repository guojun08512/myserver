package cmd

import (
	"myserver/pkg/server"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	RunE: func(cmd *cobra.Command, Args []string) error {
		server.Start()
		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
