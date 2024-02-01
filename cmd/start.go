package cmd

import (
	"github.com/Mynzhigit/snippetbox/conf"
	"github.com/Mynzhigit/snippetbox/server"

	"github.com/spf13/cobra"
)

// Version command
func init() { // nolint: gochecknoinits
	rootCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start " + conf.Executable,
		Long:  "Start " + conf.Executable,
		Run: func(cmd *cobra.Command, args []string) {
			server.ListenAndServe(&logger)
		},
	})
}
