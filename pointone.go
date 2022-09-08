package piscine

import "fmt"

func PointOne(n *int) {
	var a int = 1
	var ap *int //
	ap = &a
	fmt.Printf("Valeur de l'adresse %p: %d\n", ap, *ap)
}
