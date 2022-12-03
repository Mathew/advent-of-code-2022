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

func Intersection[M []V, V comparable](a M, b M) M {
	m := map[V]bool{}
	intersects := []V{}

	for _, i := range a {
		m[i] = true
	}

	for _, j := range b {
		if _, ok := m[j]; ok {
			intersects = append(intersects, j)
		}
	}

	return intersects
}
