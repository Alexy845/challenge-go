package piscine

func MakeRange(min, max int) []int {
	o := make([]int, max-min)
	if min > max {
		return []int(nil)
	}

	for i := 0; i < max-min; i++ {
		o[i] = i + min
	}
	return o
}
