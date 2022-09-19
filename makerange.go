package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return nil
	}
	o := make([]int, min, max)
	for i := min; i < max; i++ {
		o[i] = i + 1
	}
	return o
}
