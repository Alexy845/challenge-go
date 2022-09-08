package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string {
	sString := []rune(s)
	for c := len(sString) - 1; c >= 0; c-- {
		z01.PrintRune(sString[c])
	}
}
