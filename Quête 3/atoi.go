package piscine

func Atoi(s string) int {
	srune1 := []rune(s)
	srune := []rune{}
	neg := false
	pos := false
	for v, c := range srune1 {
		if int(c) == 45 {
			if v == 0 {
				if neg == true {
					return 0
				} else {
					neg = true
				}
			} else {
				return 0
			}
		} else if int(c) == 43 {
			if v == 0 {
				if neg == true {
					return 0
				}
				if pos == true {
					return 0
				}
				pos = true
			} else {
				return 0
			}
		} else {
			srune = append(srune, c)
		}
	}
	sint := []int{}
	for _, c := range srune {
		if int(c)-48 <= 9 && int(c)-48 >= 0 {
			sint = append(sint, int(c)-48)
		} else {
			return 0
		}
	}
	result := 0
	for _, c := range sint {
		result = result*10 + c
	}
	if neg == true {
		return result / -1
	}
	return result
}
