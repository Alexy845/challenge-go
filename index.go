package piscine

func Index(s string, toFind string) int {
	for t := range toFind {
		var l string = s[t : len(toFind)-1]
		if l == toFind {
			return t
		} else if len(toFind) == 0 {
			return -1
		}
	}
	return -1
}
