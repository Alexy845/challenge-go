package piscine

func IsAlpha(s string) bool {
	for _, i := range s {
		if !(i >= 48 && i <= 57) {
			return false
		}
	}
	return true
}
