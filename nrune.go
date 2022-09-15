package piscine

func NRune(s string, n int) rune {
	c := []rune(s)
	for i := range c {
		if i == n {
			return c[i]
		}
	}
	return 0
}
