package repository

import (
	"context"
	"example-service/internal/constants"
	"example-service/internal/helper"
	"example-service/internal/model"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

//go:generate mockery --name=TaskRepository --case=snake
type TaskRepository interface {
	Find(ctx context.Context, query *model.Query) ([]*model.Task, int64, error)
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: db,
	}
}

type taskRepositoryImpl struct {
	db *gorm.DB
}

func (i *taskRepositoryImpl) task(ctx context.Context) *gorm.DB {
	return i.db.WithContext(ctx).Debug().Model(&model.Task{})
}

func (i *taskRepositoryImpl) Find(ctx context.Context, query *model.Query) (tasks []*model.Task, total int64, err error) {
	fields, values := helper.BuildFilters(query.Filters)
	//Filters
	tx := i.task(ctx).Where(strings.Join(fields, " AND "), values...)
	//Preloads
	searchFields, searchValues := helper.BuildFilters(helper.BuildSearchFilter(query.Q, query.SearchFields...))
	if len(query.Preloads) > 0 {
		for _, item := range query.Preloads {
			tx = tx.Preload(item)
		}
	}
	//Joins
	if len(query.Joins) > 0 {
		joinsQuery, joinsWhere, selectData := helper.BuildJoins("tasks", query.Joins)
		tx = tx.Joins(joinsQuery)
		//Select fields, both of on joins and main fields you want.
		if len(selectData) > 0 {
			query.Select = append(query.Select, selectData...)
		}
		tx = tx.Where(joinsWhere, searchValues...)
	}
	tx = tx.Where(strings.Join(searchFields, " OR "), searchValues...)
	//Sorting
	if query.Sort != nil {
		tx.Order(fmt.Sprintf("%s %s", query.Sort.Key, query.Sort.SortBy))
	}
	//Count total
	if query.HaveCount {
		if err := tx.Count(&total).Error; err != nil {
			return nil, 0, err
		}
	}
	//Pagination
	if query.Pagination == nil {
		query.Pagination = helper.BuildPagination(constants.PaginationPage, constants.PaginationLimit)
	} else {
		query.Pagination = helper.BuildPagination(query.Pagination.Page, query.Pagination.Limit)
	}
	err = tx.
		Select(strings.Join(query.Select, ",")).
		Offset(query.Pagination.Offset).
		Limit(query.Pagination.Limit).
		Find(&tasks).Error

	return tasks, total, err
}
