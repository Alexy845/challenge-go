package piscine

func SplitWhiteSpaces(s string) []string {
	a := 0
	b := 0
	c := []string{}
	for i, d := range s {
		if d == ' ' {
			if s[i-1] == ' ' {
				b = i + 2
			} else {
				b = i
				c = append(c, s[a:b])
				a = i + 1
			}
		}
	}
	c = append(c, s[a:])
	return c
}
