package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return -1
	}

	for t := range s {
		var l string = s[t : len(toFind)-1]
		if l == toFind {
			return t
		}
	}
	return -1
}
