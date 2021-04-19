package lms

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var Url = "https://lms.ua.energy"

func (lms *Lms) Login(name, psw string) ([]*http.Cookie, error) {
	cookies, token, err := getToken()
	if err != nil {
		return nil, err
	}

	data := url.Values{
		"csrfmiddlewaretoken": {token},
		"username":            {name},
		"password":            {psw},
		"next":                {"/"},
	}

	req, err := http.NewRequest("POST", Url+"/login/", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp.Cookies(), nil
}

func getToken() ([]*http.Cookie, string, error) {
	resp, err := http.Get(Url)
	if err != nil {
		return nil, "", err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)

	body = body[strings.Index(body, "csrfmiddlewaretoken")+20:]
	body = body[0:strings.Index(body, ">")]
	t := strings.Split(body, "=")
	if len(t) < 2 {
		return nil, "", errors.New("login: can't find csrfmiddlewaretoken")
	}
	token := strings.Replace(t[1], "\"", "", -1)
	return resp.Cookies(), token, nil
}
