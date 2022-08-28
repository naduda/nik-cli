package gpee

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"net/url"
	"nik-cli/gpee/model"
	"strconv"
	"strings"
)

const (
	sendDateListUrl        = "https://www.gpee.com.ua/index.php/maket_rdn/send_date"
	sendDateRefreshListUrl = "https://www.gpee.com.ua/index.php/maket_vdr/send_date_vdr"
	profileUrl             = "https://www.gpee.com.ua/index.php/maket_rdn/new_request"
	profileRefreshUrl      = "https://www.gpee.com.ua/index.php/maket_vdr/new_request_vdr"
)

//goland:noinspection GoNilness
func (g *Gpee) HistoryPerDate(stationId, date string) ([]model.HistoryDataRow, error) {
	var res []model.HistoryDataRow

	rows, err := historyList(date, g.Cookies, false)
	if err != nil {
		return res, err
	}

	for _, r := range rows {
		if r.Id != stationId {
			continue
		}
		histData, err := historyData(r.Id, r.Code, date, g.Cookies, false)
		if err != nil {
			return res, err
		}

		for i, r := range histData {
			if i > 23 {
				break
			}
			r.Date = date
			r.Hour = strconv.Itoa(i + 1)
			res = append(res, r)
		}
	}

	rows, err = historyList(date, g.Cookies, true)
	if err != nil {
		return res, err
	}

	stopIdx := 23
	for _, r := range rows {
		if r.Id != stationId {
			continue
		}
		histData, err := historyData(r.Id, r.Code, date, g.Cookies, true)
		if err != nil {
			return res, err
		}

		nsi := 0
		for i, r := range histData {
			if i > stopIdx {
				break
			}
			if r.E == "" {
				nsi = i
				continue
			}

			res[i].E = r.E
			res[i].P = r.P
		}
		stopIdx = nsi
	}

	return res, nil
}

func historyData(id, code, date string, cookies []*http.Cookie, refresh bool) ([]model.HistoryDataRow, error) {
	data := url.Values{
		"date_req": {date},
		"objs_req": {id},
		"ml_id":    {code},
	}

	rUrl := profileUrl
	if refresh {
		rUrl = profileRefreshUrl
		data = url.Values{
			"date_req_vdr": {date},
			"objs_req_vdr": {id},
			"ml_id":        {code},
		}
	}

	req, err := http.NewRequest("POST", rUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	setRequestHeaders(req, data, cookies)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	content, err := respContent(resp.Body)
	if err != nil {
		return nil, err
	}
	content = content[strings.Index(content, "<table ") : strings.Index(content, "</table>")+8]
	content = strings.Replace(content, "<br>", "", -1)

	h := model.Table{}
	err = xml.NewDecoder(bytes.NewBuffer([]byte(content))).Decode(&h)
	if err != nil {
		return nil, err
	}

	return historyDataFromRows(h.Body.Rows), err
}

func historyDataFromRows(rows []model.Row) []model.HistoryDataRow {
	var res []model.HistoryDataRow

	for _, r := range rows {
		row := model.HistoryDataRow{
			E: trim(r.Cell[1].Content, ""),
			P: trim(r.Cell[2].Content, ""),
		}
		res = append(res, row)
	}

	return res
}

func trim(t string, miss string) string {
	t = strings.Replace(t, " ", "", -1)
	t = strings.Replace(t, "\r", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "\t", "", -1)
	if t == "" {
		return miss
	}
	return t
}

func historyList(date string, cookies []*http.Cookie, refresh bool) ([]model.HistoryListRow, error) {
	data := url.Values{
		"date": {date},
	}

	rUrl := sendDateListUrl
	if refresh {
		rUrl = sendDateRefreshListUrl
	}
	req, err := http.NewRequest("POST", rUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	setRequestHeaders(req, data, cookies)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	content, err := respContent(resp.Body)
	if err != nil {
		return nil, err
	}

	if startIdx := strings.Index(content, "<table"); startIdx >= 0 {
		// removing trash from content to prevent decode xml error - in other way Rows is nil
		content = content[startIdx:]
	}

	h := model.Table{}
	err = xml.NewDecoder(bytes.NewBuffer([]byte(content))).Decode(&h)
	if err != nil {
		return nil, err
	}

	return historyListFromRows(h.Body.Rows), nil
}

func historyListFromRows(rows []model.Row) []model.HistoryListRow {
	var res []model.HistoryListRow

	for _, r := range rows {
		row := model.HistoryListRow{
			Id:     r.Cell[1].Content,
			Code:   r.Cell[2].Content,
			Date:   r.Cell[4].Content,
			Status: r.Cell[5].Content,
		}
		res = append(res, row)
	}

	return res
}

func setRequestHeaders(req *http.Request, data url.Values, cookies []*http.Cookie) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Accept-Encoding", "gzip")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
}
