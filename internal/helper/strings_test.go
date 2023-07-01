package helper

import (
	"testing"
)

func TestIndexOfSliceString(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		input1 string
		output int
	}{
		{
			name:   "test case exist",
			input:  []string{"5", "10", "15", "20", "25", "30"},
			input1: "20",
			output: 3,
		},
		{
			name:   "test case not exist",
			input:  []string{"5", "10", "15", "20", "25", "30"},
			input1: "19",
			output: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOfSliceString(tt.input, tt.input1); got != tt.output {
				t.Errorf("IndexOfSliceString() = %v, want %v", got, tt.output)
			}
		})
	}
}

func TestRemoveAccents(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "test case 1",
			input:  "Nguyễn Thanh Việt",
			output: "Nguyen Thanh Viet",
		},
		{
			name:   "test case 1",
			input:  "Đỗ Đào Duy",
			output: "Do Dao Duy",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAccents(tt.input); got != tt.output {
				t.Errorf("Normalize() = %v, want %v", got, tt.output)
			}
		})
	}
}
