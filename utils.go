package golesque

/*
 * Contains utility functions
 */

// Converts value to int. true -> 1, false -> 0
func boolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

// Converts value to bool. 0 -> false, _ -> true
func intToBool(value int) bool {
	if value == 0 {
		return false
	} else {
		return true
	}
}
