package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) == 0 || len(a) == 1 {
		return true
	}
	croissant_aux_beurre := true
	if a[0] > a[1] {
		croissant_aux_beurre = false
	}
	for i := 1; i < len(a); i++ {
		if croissant_aux_beurre == true {
			if f(a[i-1], a[i]) > 0 {
				return false
			}
		} else {
			if f(a[i-1], a[i]) < 0 {
				return false
			}
		}
	}
	return true
}
