package lms

import (
	"nik-cli/lms/handlers"
	"nik-cli/server"
)

func InitHandlers(s *server.Server) {
	s.AddHandler("/exit", handlers.ExitHandler)
}
