package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return nil
	}
	o := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		o[i] = i + min
	}
	return o
}
