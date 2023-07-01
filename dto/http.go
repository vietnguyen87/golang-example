package dto

import "example-service/internal/model"

type ErrorResp struct {
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ListReq struct {
	Q string `json:"q" form:"q" query:"q"`
	// Sort field - support single field
	Sort *Sort `json:"sort,omitempty"`
	// Filter - support multiple field with AND condition
	Filters []*Filter `json:"filters,omitempty"`
	//Preloads - support multiple preload
	Preloads []string `json:"preloads,omitempty"`
	//Joins - support multiple join
	Joins []*model.Join `json:"joins,omitempty"`
	//query string
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Filter struct {
	Key    string      `json:"key"`
	Method string      `json:"method"`
	Value  interface{} `json:"value"`
}

type Sort struct {
	Key   string `json:"key,omitempty"`
	IsAsc bool   `json:"is_asc,omitempty"`
}

type Pagination struct {
	Limit int32 `json:"limit" form:"limit"`
	Page  int32 `json:"page" form:"page"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (p *ListReq) SetPagination(page, limit int32) {
	p.Pagination = &Pagination{
		Limit: limit,
		Page:  page,
	}
}

func (p *ListReq) SetSort(key string, isAsc bool) {
	if key == "" {
		return
	}
	p.Sort = &Sort{
		Key:   key,
		IsAsc: isAsc,
	}
}

type ListResp[T any] struct {
	Data       T           `json:"data"`
	Total      int64       `json:"total"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Metadata   *Metadata   `json:"metadata,omitempty"`
}

type Metadata struct {
	Pagination Pagination `json:"pagination"`
	Total      int64      `json:"total"`
}

func (p *ListReq) GetFilters() []*Filter {
	if p != nil {
		return p.Filters
	}
	return nil
}

func (p *ListReq) GetSort() *Sort {
	if p != nil {
		return p.Sort
	}
	return nil
}

func (p *ListReq) GetPagination() *Pagination {
	if p != nil {
		return p.Pagination
	}
	return nil
}

func (p *Pagination) GetLimit() int32 {
	if p != nil {
		return p.Limit
	}
	return 0
}

func (p *Pagination) GetPage() int32 {
	if p != nil {
		return p.Page
	}
	return 0
}

func (f *Filter) GetKey() string {
	if f != nil {
		return f.Key
	}
	return ""
}

func (f *Filter) GetMethod() string {
	if f != nil {
		return f.Method
	}
	return ""
}

func (s *Sort) GetIsAsc() bool {
	if s != nil {
		return s.IsAsc
	}
	return false
}

func (s *Sort) GetKey() string {
	if s != nil {
		return s.Key
	}
	return ""
}
