package builder

import (
	"example-service/dto"
	"github.com/stretchr/testify/assert"
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
