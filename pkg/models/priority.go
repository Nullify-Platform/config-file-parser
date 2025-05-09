package models

const (
	PriorityUrgent        string = "URGENT"
	PriorityImportant     string = "IMPORTANT"
	PriorityInformational string = "INFORMATIONAL"
	PriorityUnknown       string = "UNKNOWN"

	// To be deprecated
	PriorityNegligible string = "NEGIGIBLE"
	PriorityLow        string = "LOW"
	PriorityMedium     string = "MEDIUM"
	// important is repeated
	// urgent is repeated
)

var priorityToInt = map[string]int{
	PriorityUnknown:       0,
	PriorityInformational: 1,
	PriorityImportant:     3,
	PriorityUrgent:        4,

	// To be deprecated
	PriorityNegligible: 1,
	PriorityLow:        1,
	PriorityMedium:     2,
}

func ComparePriority(priority1, priority2 string) int {
	val1, ok1 := priorityToInt[priority1]
	val2, ok2 := priorityToInt[priority2]

	// If priority1 is NONE and priority2 is not in the map, treat them as equal
	if priority1 == PriorityUnknown && !ok2 {
		return 0
	}
	// If priority2 is NONE and priority1 is not in the map, treat them as equal
	if priority2 == PriorityUnknown && !ok1 {
		return 0
	}

	if !ok1 && !ok2 {
		return 0
	} else if !ok1 {
		return -1
	} else if !ok2 {
		return 1
	}

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	} else {
		return 0
	}
}

func SeverityToPriority(severity string) string {
	switch severity {
	case "CRITICAL":
		return PriorityUrgent
	case "HIGH", "MEDIUM":
		return PriorityImportant
	case "LOW":
		return PriorityInformational
	default:
		return PriorityInformational
	}
}
