package maths

func Sum(ints ...int) int {
	total := 0
	for _, i := range ints {
		total += i
	}

	return total
}

func Max(ints ...int) int {
	max := ints[0]

	for _, i := range ints[1:] {
		if i > max {
			max = i
		}
	}

	return max
}
