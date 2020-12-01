package Utils

func IntArrayFind(slice []int, value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func IntArrayContain(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Int64ArrayFind(slice []int64, value int64) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func Int64ArrayContain(slice []int64, value int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func StrArrayContain(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
