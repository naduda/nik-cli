package handlers

import (
	"net/http"
	"os"
)

var TestFlag string

func ExitHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
