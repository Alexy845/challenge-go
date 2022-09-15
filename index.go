package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return -1
	}
	for t := range s {
		l := s[t:len(toFind)]
		if l == toFind {
			return t
		}
	}
	return -1
}
