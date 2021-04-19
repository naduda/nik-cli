package lms

import (
	"nik-cli/lms/handlers"
	"nik-cli/server"
)

type Lms struct {
	TokenName  string
	TokenValue string
}

func NewLms() Lms {
	return Lms{}
}

func InitHandlers(s *server.Server) {
	s.AddHandler("/exit", handlers.ExitHandler)
}
