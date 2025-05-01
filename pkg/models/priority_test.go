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
		{"NONE", "INVALID", 0},
		{"NONE", "UNKNOWN", 0},
		{"UNKNOWN", "", 0},
		{"", "UNKNOWN", 0},
	}

	for _, test := range tests {
		result := ComparePriority(test.priority1, test.priority2)
		if result != test.expected {
			t.Errorf("ComparePriority(%s, %s) = %d; want %d",
				test.priority1, test.priority2, result, test.expected)
		}
	}
}

func TestSeverityToPriority(t *testing.T) {
	tests := []struct {
		name     string
		severity string
		want     string
	}{
		{"Critical to Urgent", SeverityCritical, PriorityUrgent},
		{"High to Important", SeverityHigh, PriorityImportant},
		{"Medium to Important", SeverityMedium, PriorityImportant},
		{"Low to Informational", SeverityLow, PriorityInformational},
		{"Unknown to Informational", SeverityUnknown, PriorityInformational},
		{"Empty to Informational", "", PriorityInformational},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SeverityToPriority(tt.severity); got != tt.want {
				t.Errorf("SeverityToPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
