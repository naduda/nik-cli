package gpee

import (
	"net/http"
	"net/url"
)

const LoginUrl = "https://www.gpee.com.ua/login/try"

func Login(login, password string) ([]*http.Cookie, error) {
	resp, err := http.PostForm(LoginUrl, url.Values{
		"login":    {login},
		"password": {password},
	})
	if err != nil {
		return nil, err
	}
	return resp.Cookies(), err
}
