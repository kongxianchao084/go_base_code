package main

import "fmt"

func main() {
	/*
		1：
		定义一个匿名函数，直接进行调用。通常只会使用一次。
		如果匿名函数被赋值给变量，可以调用多次
	*/
	// 匿名函数
	func() {
		fmt.Println("我是一个匿名函数")
	}()

	fun3 := func() {
		fmt.Println("我也是一个匿名函数")
	}
	fun3()

	/*
		2：带参数的匿名函数
	*/
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2) // 1 2
	/*
		3：定义带返回值匿名函数
	*/
	res1 := func(a, b int) int {
		return a + b
	}(10, 20) // 匿名函数调用了，将执行结果给res1
	fmt.Println(res1) // 30

	res2 := func(a, b int) int {
		return a + b
	} // 匿名函数未调用，将函数体给res1=2
	fmt.Println(res2) // 0x104b7dbd0

	/*
		4：回调函数：将匿名函数作为另一个函数的参数
	*/
	fmt.Printf("%T\n", add)  // func(int, int) int
	fmt.Printf("%T\n", oper) // func(int, int, func(int, int) int) int

	res3 := oper(10, 20, add)
	fmt.Println(res3)

	/*
		5：闭包：将匿名函数作为另一个函数的返回值  局部变量+内层函数
		外层函数会保留内层函数局部变量，生命周期会发生改变.
	*/
	res5 := increment()      // res5 = fun
	fmt.Printf("%T\n", res5) // func() int
	v1 := res5()
	fmt.Println(v1) // 1
	v2 := res5()
	fmt.Println(v2) // 2
}

// 加法运算  回调函数：作为参数传递给oper
func add(a, b int) int {
	return a + b
}

// oper:高阶函数
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun) // 10 20 0x1001df5d0
	res4 := fun(a, b)
	return res4 // 30
}

func increment() func() int { // 外层函数
	// 1 定义一个局部变量
	i := 0
	// 2 定义一个匿名函数，给变量自增并返回
	fun := func() int { // 内层函数
		i++
		return i
	}
	// 3 返回该匿名函数
	return fun
}
