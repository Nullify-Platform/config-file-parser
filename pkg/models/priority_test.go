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

		// To be deprecated, for backwards compatibility
		{"NEGIGIBLE", "URGENT", -1},
		{"NEGIGIBLE", "IMPORTANT", -1},
		{"NEGIGIBLE", "INFORMATIONAL", 0},
		{"NEGIGIBLE", "UNKNOWN", 1},
		{"NEGIGIBLE", "NONE", 1},
		{"NEGIGIBLE", "", 1},
		{"", "NEGIGIBLE", -1},
		{"NEGIGIBLE", "NEGIGIBLE", 0},
		{"LOW", "URGENT", -1},
		{"LOW", "IMPORTANT", -1},
		{"LOW", "INFORMATIONAL", 0},
		{"LOW", "UNKNOWN", 1},
		{"LOW", "NONE", 1},
		{"LOW", "", 1},
		{"", "LOW", -1},
		{"LOW", "LOW", 0},
		{"MEDIUM", "URGENT", -1},
		{"MEDIUM", "IMPORTANT", -1},
		{"MEDIUM", "INFORMATIONAL", 1},
		{"MEDIUM", "UNKNOWN", 1},
		{"MEDIUM", "NONE", 1},
		{"MEDIUM", "", 1},
		{"", "MEDIUM", -1},
		{"MEDIUM", "MEDIUM", 0},
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
