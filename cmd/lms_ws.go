package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"nik-cli/lms"
	"nik-cli/lms/ws"
)

var lmsCmdLogin string
var lmsCmdPassword string

var lmsWebsocketCmd = &cobra.Command{
	Use:   "lms-ws",
	Short: "https://www.gpee.com.ua/",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to websocket",
	Run: func(cmd *cobra.Command, args []string) {
		lms, err := lms.NewLms(lmsCmdLogin, lmsCmdPassword)
		if err != nil {
			log.Fatal(err)
		}
		wsClient := ws.NewInstance(lms.Cookies)
		if err := wsClient.Connect(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	connectCmd.PersistentFlags().StringVarP(&lmsCmdLogin, "login", "l", "", "Lms login")
	connectCmd.PersistentFlags().StringVarP(&lmsCmdPassword, "password", "p", "", "Lms password")
	lmsWebsocketCmd.AddCommand(connectCmd)
}
