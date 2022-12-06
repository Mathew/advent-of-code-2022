package structures

func MapValues[M map[K]V, K comparable, V any](dict M) []V {
	values := make([]V, 0, len(dict))

	for _, v := range dict {
		values = append(values, v)
	}

	return values
}

func MapKeys[M map[K]V, K comparable, V any](dict M) []K {
	keys := make([]K, 0, len(dict))

	for k := range dict {
		keys = append(keys, k)
	}

	return keys
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
	intersects := map[V]bool{}

	for _, i := range a {
		m[i] = true
	}

	for _, j := range b {
		if _, ok := m[j]; ok {
			intersects[j] = true
		}
	}

	return MapKeys(intersects)
}

func Sequence(start, end int) []int {
	is := make([]int, end-start+1)

	for i := start; i < end+1; i++ {
		is[i-start] = i
	}

	return is
}

func Set[T comparable](is ...T) []T {
	seen := map[T]bool{}

	for _, i := range is {
		seen[i] = true
	}

	return MapKeys(seen)
}
