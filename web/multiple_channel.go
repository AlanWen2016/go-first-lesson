package main

import "fmt"

func fibonacci(c, quit chan int) {
	fmt.Println(c)
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	// fmt.Println(c)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// fmt.Println(i)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
