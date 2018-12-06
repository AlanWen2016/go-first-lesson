package main

import "fmt"

func SumTwoValue(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	x := 3
	y := 4

	xPlusY, xTimesY := SumTwoValue(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPlusY)
	fmt.Printf("%d + %d = %d\n", x, y, xTimesY)
}
