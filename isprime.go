package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for c := 2; c < nb; c++ {
		if nb%c == 0 {
			return false
		}
	}
	return true
}
