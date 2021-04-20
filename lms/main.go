package lms

import (
	"net/http"
	"nik-cli/lms/handlers"
	"nik-cli/server"
)

type Lms struct {
	Cookies []*http.Cookie
}

func NewLms(name, psw string) (Lms, error) {
	res := Lms{}
	cookies, err := res.Login(name, psw)
	if err != nil {
		return res, err
	}
	res.Cookies = cookies
	return res, nil
}

func InitHandlers(s *server.Server) {
	s.AddHandler("/exit", handlers.ExitHandler)
}
