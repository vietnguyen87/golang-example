package builder

import (
	"example-service/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBuildFilters(t *testing.T) {
	t.Run("Test with nil params", func(t *testing.T) {
		values := BuildFilters(nil)
		assert.Equal(t, []*dto.Filter([]*dto.Filter(nil)), values)
	})
	t.Run("Test with keyword empty", func(t *testing.T) {
		filters := []*dto.Filter{
			{
				Key:    "classin_remove_date",
				Value:  time.Unix(1656402196, 0).Format("2006-01-02"),
				Method: "<",
			},
		}
		values := BuildFilters(filters)
		assert.Equal(t, filters, values)
	})
}

func TestBuildSearchFilter(t *testing.T) {
	t.Run("Test with keyword empty", func(t *testing.T) {
		actual := buildSearchFilter("")
		require.Nil(t, actual)
	})
	t.Run("Test with keyword has value", func(t *testing.T) {
		expected := &dto.Filter{
			Key:    "name",
			Value:  "%viet đẹp trai%",
			Method: "LIKE",
		}
		actual := buildSearchFilter("viet đẹp trai")
		assert.Equal(t, expected, actual)
	})
}
