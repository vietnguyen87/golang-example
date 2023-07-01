package builder

import (
	"example-service/dto"
	"example-service/internal/constants"
	"example-service/internal/model"
	"fmt"
)

//BuildQuery for handler want to get data
func BuildQuery(q string, filters []*model.Filter, sort *model.Sort, pagination *model.Pagination) *model.Query {
	query := &model.Query{}
	//Search by keyword
	query.SetQ(q)
	//Build Filters
	query.SetFilters(filters)
	//Build sort
	if sort != nil && sort.Key != "" {
		query.SetSort(sort)
	}
	query.SetPagination(pagination)
	return query
}

//Using for repo query data
func buildSearchFilter(q string) *dto.Filter {
	if q == "" {
		return nil
	}
	return &dto.Filter{
		Key:    "name",
		Value:  fmt.Sprint("%" + q + "%"),
		Method: "LIKE",
	}
}

func BuildFilters(filtersReq []*dto.Filter) (filters []*model.Filter) {
	for _, f := range filtersReq {
		filter := model.Filter{
			Key:    f.Key,
			Value:  buildFilter(f),
			Method: f.Method,
		}
		filters = append(filters, &filter)
	}

	return filters
}

//BuildSort for handler sorting
func BuildSort(sortReq *dto.Sort) (sort *model.Sort) {
	sortBy := constants.SortByDesc
	if sortReq.GetIsAsc() {
		sortBy = constants.SortByAsc
	}
	return &model.Sort{
		Key:    sortReq.GetKey(),
		SortBy: sortBy,
	}
}

func buildFilter(v *dto.Filter) (data interface{}) {
	if x, ok := v.Value.(float64); ok {
		return x
	}
	if x, ok := v.Value.(int64); ok {
		return x
	}
	if x, ok := v.Value.(string); ok {
		return x
	}
	if x, ok := v.Value.(bool); ok {
		return x
	}
	if x, ok := v.Value.([]interface{}); ok {
		return x
	}
	if x, ok := v.Value.([]string); ok {
		return x
	}
	return data
}

//BuildPagination for handler pagination
func BuildPagination(paginationParam *dto.Pagination) *model.Pagination {
	limit := int(paginationParam.GetLimit())
	if limit == 0 || limit > constants.PaginationLimit {
		limit = constants.PaginationLimit
	}

	page := int(paginationParam.GetPage())
	if page <= 0 {
		page = constants.PaginationPage
	}
	offset := (page - 1) * limit
	return &model.Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

func BuildErrorResponse(err error, message string) *dto.ErrorResp {
	return &dto.ErrorResp{
		Error:   err,
		Message: message,
	}
}
