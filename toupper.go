package piscine

func ToUpper(s string) string {
	for _, i := range s {
		if i >= 97 && i <= 122 {
			i += 32
		}
	}
	return ""
}
