package main

import "github.com/01-edu/z01"

func main() {
	for i := 0; i >= 48; i++ {
		z01.PrintRune(i + 47)
	}
	z01.PrintRune('\n')
}
