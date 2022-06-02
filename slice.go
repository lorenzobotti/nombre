package nombre

func contains[T comparable](arr []T, needle T) bool {
	for _, item := range arr {
		if item == needle {
			return true
		}
	}

	return false
}
