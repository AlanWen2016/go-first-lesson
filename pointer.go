package main

import "fmt"

func main() {

	// 声明实际变量
	var a int = 10

	// 声明指针变量
	var ip *int

	fmt.Printf("a 变量的地址是: %x\n", &a)

	// 指针变量赋值
	ip = &a
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	if ptr == nil {
		fmt.Println("ptr 是空指针")
	}

}
