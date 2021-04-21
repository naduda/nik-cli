package scheduler

import (
	"fmt"
	"nik-cli/lms"
	"nik-cli/lms/scheduler/model"
	"strconv"
	"time"
)

func syncJob(conf model.LmsConfig) error {
	for _, id := range conf.Lms.Ids {
		d := time.Now().Add(24 * time.Hour)
		date := d.Format("02.01.2006")
		gpeeId := strconv.Itoa(id.Gpee)
		if err := lms.Sync(date, conf.Lms.Login, conf.Lms.Psw, id.Login, id.Psw, gpeeId, id.Lms); err != nil {
			Log.Printf("sync: stationId = %s, err -> %s\n", gpeeId, err.Error())
		} else {
			Log.Printf("%s - Success!\n", gpeeId)
			fmt.Println(gpeeId, "success")
		}
	}
	return nil
}
