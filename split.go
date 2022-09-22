package piscine

func Split(s, sep string) []string {
	counter := 0
	str := ""
	tab := []string{}
	for i := 0; i < len(s); i++ {
		if counter > 0 {
			counter--
			continue
		}
		if sep == s[i:i+len(sep)] {
			counter = len(sep)
			tab = append(tab, str)
			str = ""
		} else {
			str += string(s[i])
		}
	}
	str += string(s[len([]rune(s))-1])
	tab = append(tab, str)
	return tab
}
