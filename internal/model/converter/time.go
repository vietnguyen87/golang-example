package converter

import (
	"time"

	"gorm.io/gorm"
)

func timeToString(input time.Time) string {
	return input.Format(time.RFC3339)
}

func nullTimeToString(input gorm.DeletedAt) *string {
	if !input.Valid {
		return nil
	}

	inputAsString := input.Time.Format(time.RFC3339)

	return &inputAsString
}
