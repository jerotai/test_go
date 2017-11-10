package cmd

import (
        "test_go/module/game"
        "test_go/core/model/database/system"
        "test_go/module/mq"

	"github.com/spf13/cobra"
        "fmt"
        "sync"
)

func init() {
        RootCmd.AddCommand(devCmd)
}

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start Dev",
	Long:  `Start Dev`,
	Run: func(cmd *cobra.Command, args []string) {
                switch args[0] {
                        case "dbconnect":
                                system.AdminList()
                        case "GameList":
                                response := game.NewGameList().List()
                                for _, element := range response {
                                        fmt.Printf("%+v", element)
                                }
                        case "WaitGroup":
                                wg := &sync.WaitGroup{}

                                for i := 0; i < 5; i++ {
                                        wg.Add(1)
                                        go func(wg *sync.WaitGroup, i int) {
                                                fmt.Printf("i:%d \n", i)
                                                wg.Done()
                                        }(wg, i)
                                }
                                wg.Wait()
                                fmt.Println("exit")
                        case "QSend":
                                for i := 0; i <10000; i++ {
                                        mq.NewSend().Send()
                                }

                                fmt.Println("QSend")
                        case "QReceive":
                                mq.NewReceive().Receive()
                                fmt.Println("QReceive")
                        default:
                                fmt.Print("default")
                }
	},
}
