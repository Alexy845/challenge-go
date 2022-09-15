package piscine

func AlphaCount(s string) int {
	c := []rune(s)
	e := 0
	for _, i := range c {
		if i >= 97 && i <= 122 || i >= 65 && i <= 90 {
			e += 1
		}
	}
	return e
}
