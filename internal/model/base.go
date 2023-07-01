package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID uint64 `json:"id" gorm:"primaryKey"`

	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}

type Filter struct {
	Key    string      `json:"key"`
	Value  interface{} `json:"value"`
	Method string      `json:"method"`
}

type Sort struct {
	Key    string
	SortBy string
}

type Query struct {
	Q            string
	Select       []string
	SearchFields []string
	Filters      []*Filter
	Preloads     []string
	Joins        []*Join
	Pagination   *Pagination
	Sort         *Sort
	HaveCount    bool
}

type Pagination struct {
	Page, Limit, Offset int
}

func (q *Query) SetQ(keyword string) {
	q.Q = keyword
}

func (q *Query) SetFilters(filters []*Filter) {
	q.Filters = filters
}

func (q *Query) SetSort(sort *Sort) {
	q.Sort = sort
}

func (q *Query) SetPagination(pagination *Pagination) {
	q.Pagination = pagination
}

func (q *Query) SetHaveCount(haveCount bool) {
	q.HaveCount = haveCount
}

func (q *Query) SetSearchFields(searchFields []string) {
	q.SearchFields = searchFields
}

type CountGroupBy[T any] struct {
	Field T     `json:"field"`
	Total int64 `json:"total"`
}

type Join struct {
	Key           string   `json:"key"`
	OriginalKey   string   `json:"originalKey"`
	Type          string   `json:"type"`
	Table         string   `json:"table"`
	OriginalTable string   `json:"originalTable"`
	Condition     string   `json:"condition,omitempty"`
	Select        []string `json:"select"`
}
