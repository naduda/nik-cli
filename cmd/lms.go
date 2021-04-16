package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nik-cli/lms"
	"nik-cli/lms/scheduler"
	"nik-cli/server"
)

var serverPort int
var schedulerStartAt string
var schedulerEvery string

var lmsCmd = &cobra.Command{
	Use:   "lms",
	Short: "short fot lms",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start fot test",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			scheduler.Run(schedulerEvery, schedulerStartAt)
		}()
		server := server.NewInstance(serverPort, true)
		lms.InitHandlers(&server)
		if err := server.Run(); err != nil {
			fmt.Println("Error: ", err.Error())
		}
	},
}

func init() {
	startCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 8485, "Http server's port")
	lmsCmd.PersistentFlags().StringVar(&schedulerStartAt, "at", "", "Start schedule at HH:mm")
	lmsCmd.PersistentFlags().StringVarP(&schedulerEvery, "every", "e", "", "Start schedule every 1h, 1m, 1s, 1ms")
	lmsCmd.AddCommand(startCmd)
	RootCmd.AddCommand(lmsCmd)
}
