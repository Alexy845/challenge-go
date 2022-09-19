package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nomprog := []rune(os.Args[0])
	for _, t := range nomprog[2:] {
		z01.PrintRune(t)
	}
	z01.PrintRune('\n')
}
