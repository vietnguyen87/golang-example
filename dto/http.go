package dto

type ErrorResponse struct {
	Error string `json:"error"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type GetReq struct {
	Query string `json:"q" form:"q" `
	Page  int    `json:"p" form:"p" `
	Limit int    `json:"s" form:"s"`
}

func SetPagination(page, limit int) *Pagination {
	return &Pagination{
		Page:  limit,
		Limit: page,
	}
}
