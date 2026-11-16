package models

type PaginatedResponse[T any] struct {
	Data       []T
	Page       int
	PageSize   int
	TotalPages int
	TotalItems int
}

func NewPaginatedResponse[T any](data []T, page, pageSize, total int) PaginatedResponse[T] {
	if data == nil {
		data = []T{} // serialize as [] instead of null
	}
	totalPages := 0
	if pageSize > 0 {
		totalPages = (total + pageSize - 1) / pageSize
	}
	return PaginatedResponse[T]{
		Data:       data,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: total,
		TotalPages: totalPages,
	}
}
