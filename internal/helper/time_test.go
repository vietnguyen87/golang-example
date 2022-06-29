package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const RFC3339 = "2006-01-02T15:04:05Z07:00"

func TestTimeToString(t *testing.T) {
	t.Run("Test string Equal", func(t *testing.T) {
		timeFormat := time.Unix(1656349199, 0).Format(RFC3339)
		stringTime := TimeToString(time.Unix(1656349199, 0))
		//2006-01-02T15:04:05Z07:00
		assert.Equal(t, timeFormat, stringTime)
	})
}
