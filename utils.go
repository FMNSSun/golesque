package golesque

func boolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

func intToBool(value int) bool {
	if value == 0 {
		return false
	} else {
		return true
	}
}
