package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nik-cli/crypt"
	"nik-cli/lms"
	"nik-cli/lms/scheduler"
	"nik-cli/lms/scheduler/model"
	"strconv"
)

var serverPort int
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
	Short: "read data from gpee and load to lms",
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
	Short: "configuration file for lms sync",
}

var makeCfgCmd = &cobra.Command{
	Use:   "new",
	Short: "create new configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		s := crypt.NewStorage(cfgPassword, "1.cfg")
		s.Data = model.LmsConfig{
			StartAt: schedulerStartAt,
			Every:   schedulerEvery,
			Lms: model.ConfigLms{
				Login: lmsLogin,
				Psw:   lmsPassword,
				Ids:   []model.ConfigId{},
			},
		}
		if err := s.Save(); err != nil {
			fmt.Printf("cfg: %s\n", err.Error())
		}
	},
}

var addToCfgCmd = &cobra.Command{
	Use:   "add",
	Short: "add station to configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		s := crypt.NewStorage(cfgPassword, "1.cfg")
		if err := s.Load(); err != nil {
			fmt.Printf("cfg: %s\n", err.Error())
		}
		id := model.ConfigId{
			Lms:   lmsStationId,
			Gpee:  gpeeStationId,
			Login: lmsLogin,
			Psw:   lmsPassword,
		}
		s.Data.Lms.Ids = append(s.Data.Lms.Ids, id)
		if err := s.Save(); err != nil {
			fmt.Printf("cfg: %s\n", err.Error())
		}
	},
}

func init() {
	startCmd.PersistentFlags().IntVar(&serverPort, "port", 8485, "Http server's port")
	startCmd.PersistentFlags().StringVar(&schedulerStartAt, "at", "", "Start schedule at HH:mm")
	startCmd.PersistentFlags().StringVarP(&schedulerEvery, "every", "e", "", "Start schedule every 1h, 1m, 1s, 1ms")
	lmsCmd.AddCommand(startCmd)

	syncCmd.PersistentFlags().StringVarP(&lmsDate, "date", "d", "", "Date: format = dd.MM.yyyy")
	syncCmd.PersistentFlags().StringVarP(&lmsLogin, "login", "l", "", "Lms login")
	syncCmd.PersistentFlags().StringVarP(&lmsPassword, "password", "p", "", "Lms password")
	syncCmd.PersistentFlags().IntVar(&lmsStationId, "lmsId", 0, "Lms station id")
	syncCmd.PersistentFlags().StringVar(&gpeeLogin, "gl", "", "Gpee login")
	syncCmd.PersistentFlags().StringVar(&gpeePassword, "gp", "", "Gpee password")
	syncCmd.PersistentFlags().IntVar(&gpeeStationId, "gpeeId", 0, "Gpee station id")
	lmsCmd.AddCommand(syncCmd)

	makeCfgCmd.PersistentFlags().StringVarP(&lmsLogin, "login", "l", "", "Lms login")
	makeCfgCmd.PersistentFlags().StringVarP(&lmsPassword, "password", "p", "", "Lms password")
	makeCfgCmd.PersistentFlags().StringVar(&schedulerStartAt, "at", "", "Start schedule at HH:mm")
	makeCfgCmd.PersistentFlags().StringVarP(&schedulerEvery, "every", "e", "", "Start schedule every 1h, 1m, 1s, 1ms")
	makeCfgCmd.PersistentFlags().StringVar(&cfgPassword, "cp", "12345678", "Password for encrypt/decrypt configuration file")
	cfgCmd.AddCommand(makeCfgCmd)

	addToCfgCmd.PersistentFlags().StringVarP(&lmsLogin, "login", "l", "", "Lms login")
	addToCfgCmd.PersistentFlags().StringVarP(&lmsPassword, "password", "p", "", "Lms password")
	addToCfgCmd.PersistentFlags().IntVar(&lmsStationId, "lms", 0, "Lms station id")
	addToCfgCmd.PersistentFlags().IntVar(&gpeeStationId, "gpee", 0, "Gpee station id")
	addToCfgCmd.PersistentFlags().StringVar(&cfgPassword, "cp", "12345678", "Password for encrypt/decrypt configuration file")
	cfgCmd.AddCommand(addToCfgCmd)

	lmsCmd.AddCommand(cfgCmd)
	RootCmd.AddCommand(lmsCmd)
}
