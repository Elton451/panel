package main

import "net/http"

func (a *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/orders")

	return mux
}
