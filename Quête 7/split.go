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
		if len(s[i:]) > len(sep) && sep == s[i:i+len(sep)] {
			counter = len(sep) - 1
			tab = append(tab, str)
			str = ""
		} else {
			str += string(s[i])
		}
	}
	tab = append(tab, str)
	return tab
}
