package helper

import (
	"academic-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBuildFilters(t *testing.T) {
	t.Run("Test with nil params", func(t *testing.T) {
		fields, values := BuildFilters("", nil)
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
		fields, values := BuildFilters("", filters)
		assert.Equal(t, []string{"classin_remove_date < ?"}, fields)
		assert.Equal(t, []interface{}{"2022-06-28"}, values)
	})
	t.Run("Test with filters empty", func(t *testing.T) {
		fields, values := BuildFilters("viet đẹp trai", nil)
		assert.Equal(t, []string{"name LIKE ?"}, fields)
		assert.Equal(t, []interface{}{"%viet đẹp trai%"}, values)
	})
}

func TestBuildSearchFilter(t *testing.T) {
	t.Run("Test with keyword empty", func(t *testing.T) {
		actual := buildSearchFilter("")
		require.Nil(t, actual)
	})
	t.Run("Test with keyword has value", func(t *testing.T) {
		expected := &model.Filter{
			Key:    "name",
			Value:  "%viet đẹp trai%",
			Method: "LIKE",
		}
		actual := buildSearchFilter("viet đẹp trai")
		assert.Equal(t, expected, actual)
	})
}
