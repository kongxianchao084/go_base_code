package main

import "fmt"

/*
	函数的类型
		func(参数列表的数据类型)(返回值列表的数据类型)
*/

func main() {
	fmt.Printf("%T\n", func1) // func()
	fmt.Printf("%T\n", func2) // func(int) int
	fmt.Printf("%T\n", func3) // func(float64, int, int) (int, int)
}

func func1() {}
func func2(a int) int {
	return 0
}
func func3(a float64, b, c int) (int, int) {
	return 0, 0
}
