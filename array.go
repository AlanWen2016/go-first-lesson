package main

import "fmt"

func main() {

	// 1.数组声明方法 ： var variable_name [SIZE] variable_type

	// 声明一个数组，默认填充 0
	var test [10]float32
	fmt.Println(test)

	// 初始化数组
	var balance = [5]float32{1000.110, 2, 3.4, 7.0, 23.1}
	fmt.Println(balance)

	// 访问数组元素
	fmt.Println(balance[0], balance[2])

	// 完整的例子
	example()
	example2()

}

// 一维数组的赋值和访问
func example() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = 100 + i
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func example2() {
	var a = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}

	fmt.Println(a[1])    // [4 5 6 7]
	fmt.Println(a[1][2]) // 6

}
