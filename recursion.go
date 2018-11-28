package main

import "fmt"

func Factorical(n uint64) (result uint64) {

	if n > 0 {
		result = n * Factorical(n-1)
		return result
	}
	return 1
}

func main() {
	var i int = 4
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorical(uint64(i)))
}
