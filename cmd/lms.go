package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"nik-cli/gpee"
	"nik-cli/lms"
	"nik-cli/lms/scheduler"
	"nik-cli/server"
	"strings"
	"time"
)

var serverPort int
var schedulerStartAt string
var schedulerEvery string
var lmsLogin string
var lmsPassword string

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
		inst := lms.NewLms()
		cookies, err := inst.Login(lmsLogin, lmsPassword)
		if err != nil {
			fmt.Println(err.Error())
		}
		req, err := http.NewRequest("GET", lms.Url+"/prs", nil)
		if err != nil {
			panic(err.Error())
		}
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}

		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)
		if err != nil {
			panic(err.Error())
		}
		//goland:noinspection GoUnhandledErrorResult
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		body := string(bodyBytes)
		body = body[strings.Index(body, "<script id=\"availableGenerations\""):]
		body = body[strings.Index(body, "[") : strings.Index(body, "]")+1]
		fmt.Println(body)

		var jsonStr = []byte(`{"mode": "getPRS", "data": {"generation_id": 412, "date": "2021-04-16"}}`)
		req, err = http.NewRequest("POST", lms.Url+"/api/generation/prs/get/", bytes.NewBuffer(jsonStr))
		if err != nil {
			panic(err.Error())
		}
		req.Header.Set("content-type", "application/json;charset=UTF-8")
		req.Header.Set("x-requested-with", "XMLHttpRequest")
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}

		httpClient = &http.Client{}
		resp, err = httpClient.Do(req)
		if err != nil {
			panic(err.Error())
		}
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		body = string(bodyBytes)
		fmt.Println(body)
	},
}

var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "short fot excel",
	Run: func(cmd *cobra.Command, args []string) {
		d := time.Now().Add(-24 * time.Hour)
		data, err := gpee.HistoryPerDate("login", "password", "stationId", d.Format("02.01.2006"))
		if err != nil {
			panic(err.Error())
		}
		lms.MakeExcelFile(d.Format("2006-01-02"), data)
	},
}

func init() {
	startCmd.PersistentFlags().IntVar(&serverPort, "port", 8485, "Http server's port")
	lmsCmd.PersistentFlags().StringVar(&schedulerStartAt, "at", "", "Start schedule at HH:mm")
	lmsCmd.PersistentFlags().StringVarP(&schedulerEvery, "every", "e", "", "Start schedule every 1h, 1m, 1s, 1ms")
	lmsCmd.PersistentFlags().StringVarP(&lmsLogin, "login", "l", "", "Lms login")
	lmsCmd.PersistentFlags().StringVarP(&lmsPassword, "password", "p", "", "Lms login")
	lmsCmd.AddCommand(excelCmd)
	lmsCmd.AddCommand(startCmd)
	lmsCmd.AddCommand(testCmd)
	RootCmd.AddCommand(lmsCmd)
}
