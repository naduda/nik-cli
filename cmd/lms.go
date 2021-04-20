package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nik-cli/gpee"
	"nik-cli/lms"
	"nik-cli/lms/scheduler"
	"nik-cli/server"
	"time"
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
		server := server.NewInstance(serverPort, true)
		lms.InitHandlers(&server)
		if err := server.Run(); err != nil {
			fmt.Println("Error: ", err.Error())
		}
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "start fot test",
	Run: func(cmd *cobra.Command, args []string) {
		inst, err := lms.NewLms(lmsLogin, lmsPassword)
		if err != nil {
			panic(err.Error())
		}
		d, err := time.Parse("02.01.2006", lmsDate)
		if err != nil {
			panic(err.Error())
		}
		date := d.Format("2006-01-02")
		r, err := inst.Get(date, lmsStationId)

		data, err := gpee.HistoryPerDate(gpeeLogin, gpeePassword, gpeeStationId, d.Format("02.01.2006"))
		if err != nil {
			panic(err.Error())
		}
		err = inst.Put(date, lmsStationId, r.Prs.Version, data)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	startCmd.PersistentFlags().IntVar(&serverPort, "port", 8485, "Http server's port")
	lmsCmd.PersistentFlags().StringVar(&schedulerStartAt, "at", "", "Start schedule at HH:mm")
	lmsCmd.PersistentFlags().StringVarP(&schedulerEvery, "every", "e", "", "Start schedule every 1h, 1m, 1s, 1ms")

	lmsCmd.PersistentFlags().StringVarP(&lmsDate, "date", "d", "", "Date: format = dd.MM.yyyy")
	lmsCmd.PersistentFlags().StringVarP(&lmsLogin, "login", "l", "", "Lms login")
	lmsCmd.PersistentFlags().StringVarP(&lmsPassword, "password", "p", "", "Lms password")
	lmsCmd.PersistentFlags().IntVar(&lmsStationId, "lmsId", 0, "Lms station id")

	lmsCmd.PersistentFlags().StringVar(&gpeeLogin, "gl", "", "Gpee login")
	lmsCmd.PersistentFlags().StringVar(&gpeePassword, "gp", "", "Gpee password")
	lmsCmd.PersistentFlags().StringVar(&gpeeStationId, "gpeeId", "", "Gpee station id")

	lmsCmd.AddCommand(startCmd)
	lmsCmd.AddCommand(testCmd)
	RootCmd.AddCommand(lmsCmd)
}
