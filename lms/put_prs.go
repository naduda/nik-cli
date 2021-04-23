package lms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	m "nik-cli/gpee/model"
	"nik-cli/lms/model"
	"strconv"
	"time"
)

func (l *Lms) Put(date string, id, ver int, data []m.HistoryDataRow, lmsData map[int]float64) error {
	payload := transform(date, id, data, lmsData)
	payload.Data.SourceVersion = ver

	rbBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", Url+"/api/generation/prs/put/", bytes.NewBuffer(rbBytes))
	if err != nil {
		return err
	}
	for _, cookie := range l.Cookies {
		req.AddCookie(cookie)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return nil
	}
	text := fmt.Sprintf("lms: putPRS - responseCode: %d", resp.StatusCode)
	return errors.New(text)
}

func transform(date string, id int, data []m.HistoryDataRow, lmsData map[int]float64) model.RequestBody {
	res := model.NewRequestBody("putPRS", date, id)
	d, _ := time.Parse("2006-01-02", date)
	for i, r := range data {
		v, err := strconv.ParseFloat(r.E, 64)
		if v < 0 || v == 0 {
			v = 0
		}
		v *= 1000
		if err == nil {
			for j := 0; j < 4; j++ {
				res.Data.Body[4*i+j+1] = getValueByTime(d, 4*i+j+1, v, lmsData[4*i+j+1])
			}
		}
	}
	return res
}

func getValueByTime(d time.Time, t int, gpeeValue, lmsValue float64) float64 {
	cy, cm, cd := time.Now().Date()
	y, mm, dd := d.Date()
	today := cy == y && cm == mm && cd == dd
	if !today || updatable(t) {
		return gpeeValue
	}
	return lmsValue
}

func updatable(t int) bool {
	c := time.Now().Hour()*4 + 1 + time.Now().Minute()/15
	return t >= c+4
}
