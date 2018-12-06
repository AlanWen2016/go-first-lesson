package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	compay string
	money  float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Printf("La la la la ...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.compay, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-xxx"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-xxx"}, "Harvard", 1000}
	sam := Employee{Human{"Sam", 36, "444-555-xxx"}, "Goland Inc", 1000}
	Tom := Employee{Human{"Tome", 18, "545-454-235"}, "Google Inc.", 457485}

	var i Men
	i = mike
	fmt.Println("This is mike, a Student:")
	i.SayHi()
	i.Sing("June Rain")

	i = Tom
	fmt.Println("This is Tom, am Employee: ")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what haapens")

	// 一定一个切片
	x := make([]Men, 3)

	// 切片赋值
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

}
