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
