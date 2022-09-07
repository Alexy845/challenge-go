package main

import "github.com/01-edu/z01"

func main() {
	for lettre := 0; lettre >= 9; lettre++ {
		z01.PrintRune(lettre)
	}
	z01.PrintRune('\n')
}
