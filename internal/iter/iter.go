package iter

func Filter[T comparable](f func(T) bool, items ...T) []T {
	result := []T{}

	for _, i := range items {
		if f(i) {
			result = append(result, i)
		}
	}

	return result
}
