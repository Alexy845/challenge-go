package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lst := os.Args[1:]
	for _, c := range lst {
		arr := []rune(c)                // Conversion de Args[0] (qui contient le programmname) en tableau de rune
		for i := len(arr); i > 0; i-- { // On parcourt le tableau à partir de la seconde position pour esquiver le ./
			z01.PrintRune(arr[i])
		}
		z01.PrintRune('\n')
	}
}
