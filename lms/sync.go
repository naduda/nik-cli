package lms

import (
	"nik-cli/gpee"
	"time"
)

func Sync(lmsDate, lmsLogin, lmsPassword, gpeeLogin, gpeePassword, gpeeStationId string, lmsStationId int) error {
	inst, err := NewLms(lmsLogin, lmsPassword)
	if err != nil {
		return err
	}
	d, err := time.Parse("02.01.2006", lmsDate)
	if err != nil {
		return err
	}
	date := d.Format("2006-01-02")
	r, err := inst.Get(date, lmsStationId)
	if err != nil {
		return err
	}

	gIns, err := gpee.NewGpee(gpeeLogin, gpeePassword)
	if err != nil {
		return err
	}
	data, err := gIns.HistoryPerDate(gpeeStationId, d.Format("02.01.2006"))
	if err != nil {
		return err
	}
	return inst.Put(date, lmsStationId, r.Prs.Version, data)
}
