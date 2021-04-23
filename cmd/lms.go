package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"nik-cli/crypt"
	"nik-cli/lms"
	"nik-cli/lms/scheduler"
	"strconv"
)

var schedulerStartAt string
var schedulerEvery string
var lmsLogin string
var lmsPassword string
var lmsStationId int
var gpeeLogin string
var gpeePassword string
var gpeeStationId int
var lmsDate string
var cfgPassword string
var cfgName string

var StopCh = make(chan bool)

var lmsCmd = &cobra.Command{
	Use:   "lms",
	Short: "https://lms.ua.energy/",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start periodically sync",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			scheduler.Run(schedulerEvery, schedulerStartAt, cfgPassword)
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
	Short: "Read data from gpee and load to lms",
	Run: func(cmd *cobra.Command, args []string) {
		gpeeId := strconv.Itoa(gpeeStationId)
		err := lms.Sync(lmsDate, lmsLogin, lmsPassword, gpeeLogin, gpeePassword, gpeeId, lmsStationId)
		if err != nil {
			panic(err.Error())
		}
	},
}

var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Configuration file for lms sync",
}

var encryptFileCfgCmd = &cobra.Command{
	Use:   "crypt",
	Short: "Encrypt configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		s := crypt.NewStorage(cfgPassword, "config.cfg")
		if err := s.ReadDataFromFile(cfgName); err != nil {
			fmt.Printf("cfg: %s\n", err.Error())
			return
		}
		if err := s.Save(); err != nil {
			fmt.Printf("cfg: %s\n", err.Error())
		}
	},
}

var printCfgCmd = &cobra.Command{
	Use:   "print",
	Short: "Print configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		s := crypt.NewStorage(cfgPassword, "config.cfg")
		if err := s.Load(); err != nil {
			fmt.Printf("cfg: %s\n", err.Error())
			return
		}
		jsn, err := json.MarshalIndent(s.Data, "", "  ")
		if err != nil {
			fmt.Printf("cfg: %s\n", err.Error())
		} else {
			fmt.Println(string(jsn))
		}
	},
}

func init() {
	//startCmd.PersistentFlags().IntVar(&serverPort, "port", 8485, "Http server's port")
	startCmd.PersistentFlags().StringVar(&schedulerStartAt, "at", "", "(Optional) Start schedule at HH:mm")
	startCmd.PersistentFlags().StringVarP(&schedulerEvery, "every", "e", "", "Start schedule every 1h, 1m, 1s, 1ms")
	startCmd.PersistentFlags().StringVar(&cfgPassword, "cp", "12345678", "Password for encrypt/decrypt configuration file")
	lmsCmd.AddCommand(startCmd)

	syncCmd.PersistentFlags().StringVarP(&lmsDate, "date", "d", "", "Date: format = dd.MM.yyyy")
	syncCmd.PersistentFlags().StringVarP(&lmsLogin, "login", "l", "", "Lms login")
	syncCmd.PersistentFlags().StringVarP(&lmsPassword, "password", "p", "", "Lms password")
	syncCmd.PersistentFlags().IntVar(&lmsStationId, "lmsId", 0, "Lms station id")
	syncCmd.PersistentFlags().StringVar(&gpeeLogin, "gl", "", "Gpee login")
	syncCmd.PersistentFlags().StringVar(&gpeePassword, "gp", "", "Gpee password")
	syncCmd.PersistentFlags().IntVar(&gpeeStationId, "gpeeId", 0, "Gpee station id")
	lmsCmd.AddCommand(syncCmd)

	printCfgCmd.PersistentFlags().StringVar(&cfgPassword, "cp", "12345678", "Password for encrypt/decrypt configuration file")
	cfgCmd.AddCommand(printCfgCmd)

	encryptFileCfgCmd.PersistentFlags().StringVar(&cfgPassword, "cp", "12345678", "Password for encrypt/decrypt configuration file")
	encryptFileCfgCmd.PersistentFlags().StringVarP(&cfgName, "name", "n", "", "Destination of the source file")
	cfgCmd.AddCommand(encryptFileCfgCmd)

	lmsCmd.AddCommand(cfgCmd)
	RootCmd.AddCommand(lmsCmd)
}
