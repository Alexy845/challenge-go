package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	sString := []rune(s)
	for _, c := range sString {
		z01.PrintRune(rune(c))
	}
}
