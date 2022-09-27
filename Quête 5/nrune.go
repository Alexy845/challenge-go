package piscine

func NRune(s string, n int) rune {
	c := []rune(s)
	for i := range c {
		if i == n-1 {
			return c[i]
		}
	}
	return 0
}
