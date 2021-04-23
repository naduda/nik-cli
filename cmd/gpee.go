package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nik-cli/gpee"
	"strconv"
)

var readDate string

var gpeeCmd = &cobra.Command{
	Use:   "gpee",
	Short: "https://www.gpee.com.ua/",
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read data by date",
	Run: func(cmd *cobra.Command, args []string) {
		inst, err := gpee.NewGpee(gpeeLogin, gpeePassword)
		if err != nil {
			panic(err.Error())
		}
		gpeeId := strconv.Itoa(gpeeStationId)
		data, err := inst.HistoryPerDate(gpeeId, readDate)
		if err == nil {
			for _, v := range data {
				fmt.Println(v.Date, v.Hour, v.E)
			}
		}
	},
}

func init() {
	readCmd.PersistentFlags().StringVarP(&readDate, "date", "d", "", "Date of gpee history")
	readCmd.PersistentFlags().StringVarP(&gpeeLogin, "login", "l", "", "Gpee login")
	readCmd.PersistentFlags().StringVarP(&gpeePassword, "password", "p", "", "Gpee password")
	readCmd.PersistentFlags().IntVar(&gpeeStationId, "id", 0, "Gpee station id")
	gpeeCmd.AddCommand(readCmd)
	RootCmd.AddCommand(gpeeCmd)
}
