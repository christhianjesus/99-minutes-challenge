package utils

func Contains(v string, a []string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}

	return false
}
