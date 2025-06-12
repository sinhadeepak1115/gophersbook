package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	// Created a new Router
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandeler)
	mux.HandleFunc("POST /users", api.creatUserHandeler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
