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

func Min(ints ...int) int {
	min := ints[0]

	for _, i := range ints[1:] {
		if i < min {
			min = i
		}
	}

	return min
}

func CapSignum(i, max int) int {
	if i < 0 {
		if i < -max {
			return -max
		}
	} else if i > max {
		return max
	}

	return i
}
