package main

import "fmt"

// n := 30
var n1 int = 30

func main() {
	/*
		作用域：变量可以使用的范围
			局部变量：函数内部定义的变量，叫做就局部变量
				变量在哪里定义，就只能在那个范围使用，超过这个范围，我们就认为变量被销毁。
			全局变量：函数外部定义的变量，叫做全局变量
				所有的函数都可以使用，共享一份变量，修改了就是修改了

	*/

	// 定义在main函数中，所以n的作用域就是min函数的范围
	n := 10
	fmt.Println(n) // 10

	if a := 1; a <= 10 {
		n := 20
		fmt.Println(a) // 1
		fmt.Println(n) // 20
	}
	//fmt.Println(a)  undefined: a  a的作用域在if语句中
	fmt.Println(n) // 10
	fmt.Println(n1)
}
