package models

import "testing"

func TestCompareSeverity(t *testing.T) {
	tests := []struct {
		severity1 string
		severity2 string
		expected  int
	}{
		{"INVALID", "HIGH", -1},
		{"UNKNOWN", "CRITICAL", -1},
		{"LOW", "CRITICAL", -1},
		{"MEDIUM", "CRITICAL", -1},
		{"HIGH", "CRITICAL", -1},
		{"CRITICAL", "HIGH", 1},
		{"HIGH", "MEDIUM", 1},
		{"MEDIUM", "LOW", 1},
		{"LOW", "UNKNOWN", 1},
		{"CRITICAL", "UNKNOWN", 1},
		{"HIGH", "INVALID", 1},
		{"UNKNOWN", "UNKNOWN", 0},
		{"LOW", "LOW", 0},
		{"MEDIUM", "MEDIUM", 0},
		{"HIGH", "HIGH", 0},
		{"CRITICAL", "CRITICAL", 0},
	}

	for _, test := range tests {
		result := CompareSeverity(test.severity1, test.severity2)
		if result != test.expected {
			t.Errorf("CompareSeverity(%q, %q) = %d; expected %d", test.severity1, test.severity2, result, test.expected)
		}
	}
}
