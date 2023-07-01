package helper

import (
	"example-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

func TestBuildFilters(t *testing.T) {
	t.Run("Test with nil params", func(t *testing.T) {
		fields, values := BuildFilters(nil)
		assert.Equal(t, []string(nil), fields)
		assert.Equal(t, []interface{}(nil), values)
	})
	t.Run("Test with keyword empty", func(t *testing.T) {
		filters := []*model.Filter{
			{
				Key:    "classin_remove_date",
				Value:  time.Unix(1656402196, 0).Format("2006-01-02"),
				Method: "<",
			},
		}
		fields, values := BuildFilters(filters)
		assert.Equal(t, []string{"classin_remove_date < ?"}, fields)
		assert.Equal(t, []interface{}{"2022-06-28"}, values)
	})
	t.Run("Test with filters empty", func(t *testing.T) {
		fields, values := BuildFilters(nil)
		assert.Equal(t, []string{"name LIKE ?"}, fields)
		assert.Equal(t, []interface{}{"%viet đẹp trai%"}, values)
	})
}

func TestBuildSearchFilter(t *testing.T) {
	t.Run("Test with keyword empty", func(t *testing.T) {
		actual := BuildSearchFilter("", []string{"name"})
		require.Nil(t, actual)
	})
	t.Run("Test with keyword has value", func(t *testing.T) {
		expected := []*model.Filter{
			{
				Key:    "lower(name)",
				Value:  "%viet đẹp trai%",
				Method: "LIKE",
			},
		}
		actual := BuildSearchFilter("viet đẹp trai", []string{"name"})
		expected[0].Value = "%viet dep trai%"
		assert.Equal(t, expected, actual)
	})
}

func TestBuildJoinsV2(t *testing.T) {
	tableName := "exam.exams"
	t.Run("Test with n join query and select", func(t *testing.T) {
		joins := []*model.Join{
			&model.Join{
				Key:           "exam_id",
				OriginalKey:   "id",
				Type:          "LEFT JOIN",
				Table:         "exam.exam_courses",
				OriginalTable: "exam.exams",
				Condition:     "",
				Select:        nil,
			},
			&model.Join{
				Key:           "lesson_id",
				OriginalKey:   "lesson_id",
				Type:          "INNER JOIN",
				Table:         "laravel.classes",
				OriginalTable: "exam.exam_courses",
				Condition:     "laravel.classes.from > exam.exam_courses.expired_at",
				Select:        []string{"laravel.classes.from, exam.exam_courses.course_id, exam.exam_courses.expired_at"},
			},
		}

		joinsQuery, joinsWhere, selectData := BuildJoins(tableName, joins)
		expectedSelectData := []string{"exam.exams.*", "laravel.classes.from, exam.exam_courses.course_id, exam.exam_courses.expired_at"}
		expectedJoinsQuery := "left join exam.exam_courses on exam.exam_courses.exam_id = exam.exams.id\ninner join laravel.classes on laravel.classes.lesson_id = exam.exam_courses.lesson_id"
		expectedJoinsWhere := "laravel.classes.from > exam.exam_courses.expired_at"
		if !reflect.DeepEqual(selectData, expectedSelectData) {
			t.Errorf("Expected: %v, Got: %v", expectedSelectData, selectData)
		}

		if joinsQuery != expectedJoinsQuery {
			t.Errorf("Expected: %v, Got: %v", expectedJoinsQuery, joinsQuery)
		}

		if joinsWhere != expectedJoinsWhere {
			assert.Equal(t, expectedJoinsWhere, joinsWhere)
		}
	})

}
