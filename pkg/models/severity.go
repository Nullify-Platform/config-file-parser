package models

const (
	SeverityCritical = "CRITICAL"
	SeverityHigh     = "HIGH"
	SeverityMedium   = "MEDIUM"
	SeverityLow      = "LOW"
	SeverityUnknown  = "UNKNOWN"
)

var severityToInt = map[string]int{
	SeverityUnknown:  0,
	SeverityLow:      1,
	SeverityMedium:   2,
	SeverityHigh:     3,
	SeverityCritical: 4,
}

func CompareSeverity(severity1, severity2 string) int {
	val1, ok1 := severityToInt[severity1]
	val2, ok2 := severityToInt[severity2]

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
