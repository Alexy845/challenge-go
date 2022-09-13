package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 0
	}
	if IsPrime(nb) {
		return nb
	}
	return FindNextPrime(nb + 1)
}
