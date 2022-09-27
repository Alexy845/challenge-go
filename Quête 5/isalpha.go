package piscine

func IsAlpha(s string) bool {
	for _, i := range s {
		if !(i >= 48 && i <= 122) {
			return false
		}
	}
	return true
}
