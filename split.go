package piscine

func Split(s, sep string) []string {
	counter := 0
	str := ""
	tab := []string{}
	for i := 0; i < len(s)-len(sep); i++ {
		if counter > 0 {
			counter--
			continue
		}
		if sep == s[i:i+len(sep)] {
			counter = counter + 1
			tab = append(tab, str)
			str = ""
		} else {
			str += string(s[i])
		}
	}
	str += s[len(s)-len(sep):]
	tab = append(tab, str)
	return tab
}
