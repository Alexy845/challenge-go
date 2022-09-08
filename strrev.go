package piscine

func StrRev(s string) string {
	sString := []rune(s)
	f := ""
	for c := len(sString) - 1; c >= 0; c-- {
		f += string(sString[c])
	}
	return f
}
