package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID uint64 `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Filter struct {
	Key    string
	Value  interface{}
	Method string
}

type Sort struct {
	Key    string
	SortBy string
}

type Query struct {
	Q          string
	Filters    []*Filter
	Sort       *Sort
	Pagination *Pagination
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
