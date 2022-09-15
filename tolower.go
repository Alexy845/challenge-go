package piscine

func ToLower(s string) string {
	str := ""
	srune := []rune(s)
	for _, i := range srune {
		if i >= 65 && i <= 90 {
			str += string(i + 32)
		} else {
			str += string(i)
		}
	}
	return str
}
