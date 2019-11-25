package main

import "fmt"

//返回 X+Y 和 X*Y
func Computer(X, Y int) (int, int) {
	return X+Y, X*Y
}

func main() {
	x := 10
	y := 20

	a, b := Computer(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, a)
	fmt.Printf("%d * %d = %d\n", x, y, b)
}