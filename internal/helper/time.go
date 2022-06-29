package helper

import (
	"time"

	"gorm.io/gorm"
)

func TimeToString(input time.Time) string {
	return input.Format(time.RFC3339)
}

func NullTimeToString(input gorm.DeletedAt) *string {
	if !input.Valid {
		return nil
	}

	inputAsString := input.Time.Format(time.RFC3339)

	return &inputAsString
}
