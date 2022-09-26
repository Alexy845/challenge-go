package main

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = %d, y = %d\n", points.x, points.y)
}

func printStr(s string, i1, i2 int) {
	panic("unimplemented")
}
