package config

import (
	"httpCpntroller"
	"net/http"
)

func CreateHttpMuxer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/tasks", &httpCpntroller.TaskHandler{})
	return mux
}
