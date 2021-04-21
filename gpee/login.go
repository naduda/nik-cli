package gpee

import (
	"net/http"
	"net/url"
)

const LoginUrl = "https://www.gpee.com.ua/login/try"

type Gpee struct {
	Cookies []*http.Cookie
}

func NewGpee(login, password string) (Gpee, error) {
	res := Gpee{}
	cookies, err := loginFunc(login, password)
	if err != nil {
		return res, err
	}
	res.Cookies = cookies
	return res, nil
}

func loginFunc(login, password string) ([]*http.Cookie, error) {
	resp, err := http.PostForm(LoginUrl, url.Values{
		"login":    {login},
		"password": {password},
	})
	if err != nil {
		return nil, err
	}
	return resp.Cookies(), err
}
