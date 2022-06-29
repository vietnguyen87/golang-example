package helper

import (
	"example-service/internal/constants"
	"example-service/internal/model"
	"fmt"
)

//BuildSort for handler sorting
func BuildSort(key string, isASC bool) (sort *model.Sort) {
	sortBy := constants.SORT_BY_DESC
	if isASC {
		sortBy = constants.SORT_BY_ASC
	}
	return &model.Sort{
		Key:    key,
		SortBy: sortBy,
	}
}

//BuildPagination for handler pagination
func BuildPagination(page, limit int) *model.Pagination {
	if limit <= 0 || limit > constants.PAGINATION_LIMIT {
		limit = constants.PAGINATION_LIMIT
	}

	if page <= 0 {
		page = constants.PAGINATION_PAGE
	}
	offset := (page - 1) * limit

	return &model.Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

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
//BuildFilters build filters from list fields and value
func BuildFilters(q string, filters []*model.Filter) (fields []string, values []interface{}) {
	if q != "" {
		filter := buildSearchFilter(q)
		filters = append(filters, filter)
	}
	for _, filter := range filters {
		fields = append(fields, fmt.Sprintf("%s %s ?", filter.Key, filter.Method))
		values = append(values, filter.Value)
	}
	return fields, values
}

func buildSearchFilter(q string) *model.Filter {
	if q == "" {
		return nil
	}
	return &model.Filter{
		Key:    "name",
		Value:  fmt.Sprint("%" + q + "%"),
		Method: "LIKE",
	}
}
