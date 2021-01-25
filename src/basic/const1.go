package main

import "fmt"

/*
内容：常量的使用
*/

func main() {
	/*
		常量：
		1.概念：同变量类似，程序执行过程中数值不能改变
		2.语法：
			显示类型定义
			隐式类型定义
		3.常数
			固定的数值：100，"abc"
	*/

	// 1:定义常量
	const PATH string = "http://www.baidu.com"
	const PI = 3.14
	fmt.Println(PATH)
	fmt.Println(PI)

	// 2:尝试修改常量的数值
	//PATH = "http://www.baidu1.com" cannot assign to PATH (declared const)

	// 3:定义一组常量(类型推断)
	const C1, C2, C3 = 100, 3.14, "haha"
	const (
		MALE    = 0
		FEMALE  = 1
		UNKNOWN = 3
	)

	// 4:常量定义完可以不使用，不会报错
	const M1 = 200
	const m1, m2 = 2, 3

	// 5：一组常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		a int = 100
		b
		c = 300
	)
	fmt.Println(a, b, c)
	/*
		100 100 300
	*/

	// 6:枚举类型：使用常量组作为枚举类型。一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)

}
