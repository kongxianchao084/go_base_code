package main

import "fmt"

/*
常量：iota关键字
*/

func main() {
	/*
		iota：特殊的常量，可以被编译器自动修改的常量
		每当定义一个const，iota的初始值为0
		每当定义一个常量，iota自动加1
		直到下一个const出现，iota清零
	*/
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	/*
		0 1 2
		此处定义了三个常量，从0开始依次加1
	*/

	const (
		d = iota // 0
		e
	)
	fmt.Println(d, e)
	/*
		0 1
		遇到下一个const，iota清零
		e默认和上一行一致，e = iota，定义了一个常量，加1，等于1
	*/

	// 枚举
	const (
		MALE    = iota // 0
		FEMALE         // 1
		UNKNOWN        // 2
	)
	fmt.Println(MALE, FEMALE, UNKNOWN)
}
