package cmd

import (
	"github.com/spf13/cobra"
	"nik-cli/lms"
	"nik-cli/lms/scheduler"
)

var serverPort int
var schedulerStartAt string
var schedulerEvery string
var lmsLogin string
var lmsPassword string
var lmsStationId int
var gpeeLogin string
var gpeePassword string
var gpeeStationId string
var lmsDate string

var StopCh = make(chan bool)

var lmsCmd = &cobra.Command{
	Use:   "lms",
	Short: "short fot lms",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "short fot test",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			scheduler.Run(schedulerEvery, schedulerStartAt)
		}()
		<-StopCh
		//server := server.NewInstance(serverPort, true)
		//lms.InitHandlers(&server)
		//if err := server.Run(); err != nil {
		//	fmt.Println("Error: ", err.Error())
		//}
	},
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "short fot sync",
	Run: func(cmd *cobra.Command, args []string) {
		err := lms.Sync(lmsDate, lmsLogin, lmsPassword, gpeeLogin, gpeePassword, gpeeStationId, lmsStationId)
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	startCmd.PersistentFlags().IntVar(&serverPort, "port", 8485, "Http server's port")
	startCmd.PersistentFlags().StringVar(&schedulerStartAt, "at", "", "Start schedule at HH:mm")
	startCmd.PersistentFlags().StringVarP(&schedulerEvery, "every", "e", "", "Start schedule every 1h, 1m, 1s, 1ms")

	lmsCmd.PersistentFlags().StringVarP(&lmsDate, "date", "d", "", "Date: format = dd.MM.yyyy")
	lmsCmd.PersistentFlags().StringVarP(&lmsLogin, "login", "l", "", "Lms login")
	lmsCmd.PersistentFlags().StringVarP(&lmsPassword, "password", "p", "", "Lms password")
	lmsCmd.PersistentFlags().IntVar(&lmsStationId, "lmsId", 0, "Lms station id")

	lmsCmd.PersistentFlags().StringVar(&gpeeLogin, "gl", "", "Gpee login")
	lmsCmd.PersistentFlags().StringVar(&gpeePassword, "gp", "", "Gpee password")
	lmsCmd.PersistentFlags().StringVar(&gpeeStationId, "gpeeId", "", "Gpee station id")

	lmsCmd.AddCommand(startCmd)
	lmsCmd.AddCommand(syncCmd)
	RootCmd.AddCommand(lmsCmd)
}
