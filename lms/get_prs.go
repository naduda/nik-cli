package lms

import (
	"bytes"
	"encoding/json"
	"net/http"
	"nik-cli/lms/model"
)

func (l *Lms) Get(date string, id int) (model.GetResponse, error) {
	res := model.GetResponse{}

	rb := model.NewRequestBody("getPRS", date, id)
	rbBytes, err := json.Marshal(rb)
	if err != nil {
		return res, err
	}
	req, err := http.NewRequest("POST", Url+"/api/generation/prs/get/", bytes.NewBuffer(rbBytes))
	if err != nil {
		return res, err
	}
	for _, cookie := range l.Cookies {
		req.AddCookie(cookie)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return res, err
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
