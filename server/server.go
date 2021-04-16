package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct {
	port        int
	disableCors bool
	handlers    map[string]http.HandlerFunc
}

func NewInstance(port int, disableCors bool) Server {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	handlers := map[string]http.HandlerFunc{
		"/": http.FileServer(http.Dir(path + "/web")).ServeHTTP,
	}

	return Server{port, disableCors, handlers}
}

func (s *Server) AddHandler(endpoint string, f http.HandlerFunc) {
	s.handlers[endpoint] = f
}

func (s *Server) Run() error {
	for path, handler := range s.handlers {
		handlerFunc := disableCors(handler, s.disableCors)
		http.Handle(path, handlerFunc)
	}

	addr := fmt.Sprintf(":%d", s.port)
	fmt.Printf("Server is running on port: %d\nRun your app: http://localhost%s\n", s.port, addr)
	return http.ListenAndServe(addr, nil)
}

func disableCors(h http.HandlerFunc, disable bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if disable {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
		}
		h(w, r)
	}
}
