package piscine

func BasicAtoi2(s string) int {
	srune := []rune(s)
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
	return result
}
