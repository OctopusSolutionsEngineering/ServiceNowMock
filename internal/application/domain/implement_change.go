package domain

var ImplementationCount map[string]int = map[string]int{}

// CheckChangeRequest will return true after 5 requests to check the change request.
// This means all change requests eventually move to the implementation phase
func CheckChangeRequest(id string) bool {
	if count, ok := ImplementationCount[id]; !ok {
		ImplementationCount[id] = 1
	} else {
		ImplementationCount[id] = count + 1
	}

	return ImplementationCount[id] >= 5
}
