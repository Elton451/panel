package models

import (
	"database/sql"
	"fmt"
	"time"
)

type OrderStatus string

const (
	StatusWaitingPayment OrderStatus = "WAITING_PAYMENT"
	StatusPreparing      OrderStatus = "PREPARING"
	StatusSent           OrderStatus = "SENT"
	StatusOnRoute        OrderStatus = "ON_ROUTE"
	StatusReceived       OrderStatus = "RECEIVED"
	StatusCancelled      OrderStatus = "CANCELLED"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case StatusWaitingPayment, StatusPreparing, StatusSent,
		StatusOnRoute, StatusReceived, StatusCancelled:
		return true
	}
	return false
}

type Order struct {
	Id        int         `json:"id"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UserId    int         `json:"user_id"`
}

type OrderFilters struct {
	Status OrderStatus
	UserId int
}

type OrderModel struct {
	Db *sql.DB
}

func (m *OrderModel) ListAllOrders(filters Filter[OrderFilters]) (PaginatedResponse[Order], error) {
	query := `
		SELECT id, status, created_at, user_id, COUNT(*) OVER AS total
		FROM orders
		ORDER BY created_at DESC, id DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := m.Db.Query(query, filters.PageSize, (filters.Page-1)*filters.PageSize)
	if err != nil {
		fmt.Println("Error on querying the Orders")
		return PaginatedResponse[Order]{}, err
	}

	defer rows.Close()

	var orders []Order

	total := 0
	for rows.Next() {
		var order Order

		err := rows.Scan(&order.Id, &order.Status, &order.CreatedAt, &order.UserId, &total)
		if err != nil {
			return PaginatedResponse[Order]{}, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return PaginatedResponse[Order]{}, err
	}

	return NewPaginatedResponse(orders, filters.Page, filters.PageSize, total), nil
}
