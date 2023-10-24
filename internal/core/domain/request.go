package domain

// Pagination struct
type Pagination struct {
	Limit  int `json:"limit" query:"limit" validate:"gte=-1,lte=100"`
	Offset int `json:"offset" query:"offset"`
}

// SortMethod struct
type SortMethod struct {
	Asc     bool   `json:"asc" query:"asc"`
	OrderBy string `json:"order_by" query:"order_by"`
}
