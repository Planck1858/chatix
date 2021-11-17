package utils

type RequestOptionsSchema struct {
	PaginationSchema
}

type PaginationSchema struct {
	Page    int `schema:"page"`
	PerPage int `schema:"per_page"`
}
