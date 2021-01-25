package main

import "fmt"

func main() {
	// 1.等于相关
	var a int
	a = 3
	fmt.Println(a) // 3

	a += 4         // a++相等
	fmt.Println(a) // 7

	a -= 3
	fmt.Println(a) // 4
	a *= 2
	fmt.Println(a) // 8
	a /= 3
	fmt.Println(a) // 2
	a %= 1
	fmt.Println(a) // 0
}
