package piscine

func IsUpper(s string) bool {
	c := []rune(s)
	for _, i := range c {
		if !(i >= 65 && i <= 90) {
			return false
		}
	}
	return true
}
