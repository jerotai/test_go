package cmd

import (
	"Stingray/api"

	"github.com/spf13/cobra"
	"os"
)

func init() {
	RootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "ApiService",
	Short: "Start api server",
	Long:  `Start api server`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Setenv("ConfPath", args[0])
		api.NewApiService().Start()
	},
}
