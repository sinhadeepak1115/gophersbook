package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		}
	case http.MethodPost:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("post request received"))
		case "/users":
			w.Write([]byte("user created"))
		}
	case http.MethodPut:
		w.Write([]byte("put request received"))
	}
}

func main() {
	api := &api{addr: ":8080"}
	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}
	srv.ListenAndServe()
}
