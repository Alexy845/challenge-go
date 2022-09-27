package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 || len(s) == 0 {
		return 0
	}
	for v := range s {
		if v+len(toFind) >= len(s) {
			break
		}
		var slice string = s[v : v+len(toFind)]
		if slice == toFind {
			return v
		}
	}
	return -1
}
