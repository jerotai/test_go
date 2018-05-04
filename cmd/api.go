package cmd

import (
	"routes/api"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "ApiService",
	Short: "Start api server",
	Long:  `Start api server`,
	Run: func(cmd *cobra.Command, args []string) {
		api.NewApiService().Start()
	},
}
