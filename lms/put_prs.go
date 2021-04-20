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
)

func (l *Lms) Put(date string, id, ver int, data []m.HistoryDataRow) error {
	payload := transform(date, id, data)
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

func transform(date string, id int, data []m.HistoryDataRow) model.RequestBody {
	res := model.NewRequestBody("putPRS", date, id)
	for i, r := range data {
		v, err := strconv.ParseFloat(r.E, 64)
		if v < 0 || v == 0 {
			v = 0
		}
		v *= 1000
		if err == nil {
			res.Data.Body[4*i+1] = v
			res.Data.Body[4*i+2] = v
			res.Data.Body[4*i+3] = v
			res.Data.Body[4*i+4] = v
		}
	}
	return res
}
