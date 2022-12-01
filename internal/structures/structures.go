package structures

func Values[M map[K]V, K comparable, V any](dict M) []V {
	values := make([]V, 0, len(dict))

	for _, v := range dict {
		values = append(values, v)
	}

	return values
}
