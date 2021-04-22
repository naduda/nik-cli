package main

import (
	"github.com/kardianos/service"
	"nik-cli/cmd"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	_ = cmd.RootCmd.Execute()
}

func (p *program) Stop(s service.Service) error {
	cmd.StopCh <- true
	return nil
}

// GOOS=windows GOARCH=386 go build -o nik-cli.exe .
func main() {
	//svcConfig := &service.Config{
	//	Name:        "nik",
	//	DisplayName: "Nik cli as a service",
	//	Description: "-",
	//}
	//
	//prg := &program{}
	//s, err := service.New(prg, svcConfig)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = s.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	_ = cmd.RootCmd.Execute()
}
