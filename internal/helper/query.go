package helper

import (
	"example-service/internal/constants"
	"example-service/internal/model"
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// BuildPagination for handler pagination
func BuildPagination(page, limit int) *model.Pagination {
	if limit == -1 {
		return &model.Pagination{
			Page:   1,
			Limit:  -1,
			Offset: 0,
		}
	}
	if limit <= 0 || limit > constants.PaginationLimit {
		limit = constants.PaginationLimit
	}

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

// BuildQuery for handler want to get data
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

func BuildFilters(filters []*model.Filter) (fields []string, values []interface{}) {
	for _, filter := range filters {
		if filter.Value == nil {
			fields = append(fields, fmt.Sprintf("%s %s NULL", filter.Key, filter.Method))
			continue
		}
		fields = append(fields, fmt.Sprintf("%s %s ?", filter.Key, filter.Method))
		values = append(values, filter.Value)
	}
	return fields, values
}

func BuildSearchFilter(q string, searchFields ...string) (filters []*model.Filter) {
	if q == "" {
		return nil
	}
	q = normalize(q)
	for _, searchField := range searchFields {
		filters = append(filters, &model.Filter{
			Key:    fmt.Sprintf("lower(%v)", searchField),
			Value:  strings.ToLower(fmt.Sprint("%" + q + "%")),
			Method: "LIKE",
		})
	}
	return filters
}

func BuildJoins(tableName string, joins []*model.Join) (joinsQuery, whereOnJoin string, selectData []string) {
	selectData = []string{fmt.Sprintf("%s.*", tableName)}
	if len(joins) == 0 || tableName == "" {
		return
	}
	var sliceQuery []string
	var sliceWhere []string
	for _, join := range joins {
		if len(join.Select) > 0 {
			selectData = append(selectData, join.Select...)
		}
		if join.Condition != "" {
			sliceWhere = append(sliceWhere, join.Condition)
		}
		sliceQuery = append(sliceQuery, fmt.Sprintf("%s %s on %s.%s = %s.%s", join.Type, join.Table, join.Table, join.Key, join.OriginalTable, join.OriginalKey))
	}

	return strings.ToLower(strings.Join(sliceQuery, "\n")), strings.ToLower(strings.Join(sliceWhere, " AND ")), selectData
}

func normalize(str string) string {
	trans := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(trans, str)
	result = strings.ReplaceAll(result, "đ", "d")
	result = strings.ReplaceAll(result, "Đ", "D")
	return result
}
