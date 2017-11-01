package cmd

import (
        "test_go/api/controller/game"

	"github.com/spf13/cobra"
        "fmt"
)

var devName string
//var devArgu string

func init() {
        devCmd.Flags().StringVarP(&devName, "name", "n", "", "dev's name")
        RootCmd.AddCommand(devCmd)
}

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start Dev",
	Long:  `Start Dev`,
	Run: func(cmd *cobra.Command, args []string) {
                switch args[0] {
                        case "dbconnect":
                                game.GameList()
                        default:
                                fmt.Print("default")
                }
	},
}
