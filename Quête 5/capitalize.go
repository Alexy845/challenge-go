package piscine

func Capitalize(s string) string {
	str := ""
	oi := []rune(s)
	for i, c := range []rune(s) {
		if IsLower(string(c)) {
			if i == 0 {
				str += ToUpper(string(c))
			} else if !(IsUpper(string(oi[i-1]))) && !(IsNumeric(string(oi[i-1]))) && !(IsLower(string(oi[i-1]))) {
				str += ToUpper(string(c))
			} else {
				str += string(c)
			}
		} else if IsUpper(string(c)) {
			if i == 0 {
				str += string(c)
			} else if IsUpper(string(oi[i-1])) || IsLower(string(oi[i-1])) || IsNumeric(string(oi[i-1])) {
				str += ToLower(string(c))
			} else {
				str += string(c)
			}
		} else {
			str += string(c)
		}
	}
	return str
}
