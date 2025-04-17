package models

const (
	PriorityUrgent        string = "URGENT"
	PriorityImportant     string = "IMPORTANT"
	PriorityInformational string = "INFORMATIONAL"
)

var priorityToInt = map[string]int{
	PriorityInformational: 0,
	PriorityImportant:     1,
	PriorityUrgent:        2,
}

func ComparePriority(priority1, priority2 string) int {
	val1, ok1 := priorityToInt[priority1]
	val2, ok2 := priorityToInt[priority2]

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
