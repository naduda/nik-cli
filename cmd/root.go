package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "nik-cli",
	Short: "http://nik.net.ua/ua/",
}

func init() {
	RootCmd.AddCommand(gpeeCmd)
	RootCmd.AddCommand(lmsCmd)
	RootCmd.AddCommand(lmsWebsocketCmd)
}
