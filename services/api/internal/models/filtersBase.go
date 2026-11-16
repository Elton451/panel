package models

type Filter[T any] struct {
	Page     int
	PageSize int
}
