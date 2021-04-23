package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"nik-cli/crypt"
	"nik-cli/logger"
	"strconv"
	"strings"
	"time"
)

func Run(every, at, psw string) {
	logfile := fmt.Sprintf("./%s.log", time.Now().Format("2006_01_02"))
	Log, err := logger.InitLogger(logfile)
	if err != nil {
		panic(err.Error())
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(every)
	atTime := startAtTime(at)
	s.StartAt(atTime)

	st := crypt.NewStorage(psw, "config.cfg")
	if err := st.Load(); err != nil {
		Log.Printf("config: %s\n", err.Error())
		panic(err.Error())
	}

	_, err = s.Do(func() {
		syncJob(st.Data, Log)
	})
	if err != nil {
		Log.Printf("scheduler: %s\n", err.Error())
		panic(err.Error())
	}

	s.StartAsync()
}

func startAtTime(at string) time.Time {
	if at == "" {
		return time.Now().Add(1 * time.Second)
	}
	atArr := strings.Split(at, ":")
	n := time.Now().Add(1 * time.Second)
	if len(atArr) != 2 {
		return n
	}

	h, err := strconv.Atoi(atArr[0])
	if err != nil || h > 23 {
		return n
	}

	m, err := strconv.Atoi(atArr[1])
	if err != nil || m > 59 {
		return n
	}

	return time.Date(n.Year(), n.Month(), n.Day(), h, m, 0, 0, n.Location())
}
