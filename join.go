package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i, c := range strs {
		str += c
		if i < len(strs)-1 {
			str += sep
		}
	}
	return str
}
