package piscine

func CountIf(f func(string) bool, tab []string) int {
	n := 0
	for _, i := range tab {
		if f(i) == true {
			n += 1
		}
	}
	return n
}
