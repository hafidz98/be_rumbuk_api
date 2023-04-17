package service

type WebResponse struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   interface{}     `json:"data,omitempty"`
	Meta   *PaginationMeta `json:"meta,omitempty"`
	Links  interface{}     `json:"links,omitempty"`
}

type PaginationMeta struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total_page"`
}
