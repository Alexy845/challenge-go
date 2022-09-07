package main

import "github.com/01-edu/z01"

func main() {
	for i := 49; i <= 57; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
