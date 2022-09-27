package piscine

func ToUpper(s string) string {
	str := ""
	srune := []rune(s)
	for _, i := range srune {
		if i >= 97 && i <= 122 {
			str += string(i - 32)
		} else {
			str += string(i)
		}
	}
	return str
}
