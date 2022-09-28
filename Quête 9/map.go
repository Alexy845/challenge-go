package piscine

func Map(f func(int) bool, a []int) []bool {
	l := []bool{}
	for _, i := range a {
		l = append(l, f(i))
	}
	return l
}
