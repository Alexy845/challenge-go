package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := []rune(os.Args[1])
	for _, t := range arguments[4:] {
		z01.PrintRune(t)
	}
	z01.PrintRune('\n')
}
