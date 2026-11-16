package main

import "net/http"

func (a *Application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/orders", a.listOrders)

	return a.logRequest(mux)
}
