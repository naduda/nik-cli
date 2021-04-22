package scheduler

import (
	"log"
	"nik-cli/lms"
	"nik-cli/lms/scheduler/model"
	"strconv"
	"time"
)

func syncJob(conf []model.ConfigLms, logger *log.Logger) {
	for _, item := range conf {
		for _, id := range item.Ids {
			d := time.Now().Add(24 * time.Hour)
			date := d.Format("02.01.2006")
			gpeeId := strconv.Itoa(id.Gpee)
			if err := lms.Sync(date, item.Login, item.Psw, id.Login, id.Psw, gpeeId, id.Lms); err != nil {
				logger.Printf("sync: stationId = %s, err -> %s\n", gpeeId, err.Error())
			} else {
				logger.Printf("%s - Success!\n", gpeeId)
			}
		}
	}
}
