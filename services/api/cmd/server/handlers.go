package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/elton451/panel-monorepo/services/api/internal/models"
)

func (a *Application) listOrders(w http.ResponseWriter, r *http.Request) {
	var filters models.Filter[models.OrderFilters]

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := a.Orders.ListAllOrders(filters)
	if err != nil {
		fmt.Fprintf(w, "Hi! Mistakes were made: %s", err)
	}

	fmt.Fprintf(w, "Listing Orders: %+v", orders)
	json.NewEncoder(w).Encode(orders)
}
