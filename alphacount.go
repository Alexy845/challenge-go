package piscine

func AlphaCount(s string) int {
	c := []rune(s)
	e := 0
	for i := range c {
		if i >= 97 || i <= 122 {
			e += 1
		}
	}
	return 0
}
