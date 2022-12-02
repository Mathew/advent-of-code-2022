package structures

func MapValues[M map[K]V, K comparable, V any](dict M) []V {
	values := make([]V, 0, len(dict))

	for _, v := range dict {
		values = append(values, v)
	}

	return values
}

func MapInvert[M map[K]V, K comparable, V comparable](dict M) map[V]K {
	m := map[V]K{}

	for k, v := range dict {
		m[v] = k
	}

	return m
}
