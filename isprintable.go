package piscine

func IsPrintable(s string) bool {

	for _, i := range s {
		if !(i >= 33 && i <= 126) {
			return false
		}
	}
	return true
}
