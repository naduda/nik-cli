package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nik-cli/gpee"
)

var readDate string

var gpeeCmd = &cobra.Command{
	Use:   "gpee",
	Short: "short fot gpee",
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "short fot read",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := gpee.HistoryPerDate(gpeeLogin, gpeePassword, gpeeStationId, readDate)
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
	readCmd.PersistentFlags().StringVar(&gpeeStationId, "id", "", "Gpee station id")
	gpeeCmd.AddCommand(readCmd)
	RootCmd.AddCommand(gpeeCmd)
}
