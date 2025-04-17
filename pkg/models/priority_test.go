package models

import "testing"

func TestComparePriority(t *testing.T) {
	tests := []struct {
		priority1 string
		priority2 string
		expected  int
	}{
		{"INVALID", "URGENT", -1},
		{"UNKNOWN", "IMPORTANT", -1},
		{"INFORMATIONAL", "URGENT", -1},
		{"IMPORTANT", "URGENT", -1},
		{"URGENT", "IMPORTANT", 1},
		{"IMPORTANT", "INFORMATIONAL", 1},
		{"URGENT", "UNKNOWN", 1},
		{"IMPORTANT", "INVALID", 1},
		{"UNKNOWN", "UNKNOWN", 0},
		{"INFORMATIONAL", "INFORMATIONAL", 0},
		{"IMPORTANT", "IMPORTANT", 0},
		{"URGENT", "URGENT", 0},
	}

	for _, test := range tests {
		result := ComparePriority(test.priority1, test.priority2)
		if result != test.expected {
			t.Errorf("ComparePriority(%s, %s) = %d; want %d",
				test.priority1, test.priority2, result, test.expected)
		}
	}
}
