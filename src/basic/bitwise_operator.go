package main

import "fmt"

func main() {
	/*
		位运算符
			将数值，转为二进制后，按位操作
		按位&：对应位的值都为1才为1，有一个0就为0
		按位｜：对应位的值都为-才为0，有一个1就为1
		异或^：
			二元：a ^ b
				对应位的值不同为1，相同为0
			一元：^a
				按位取反
					1---0 0---1
		位清空&^：
			对于 a &^ b
				对于b上的每个数值
				如果为0，则取a对应位上的值
				如果为1，则结果位就取0

		位移运算符
		<<：按位左移  放大2^n次方
		>>：按位右移  缩小2^n次方
	*/

	a := 60
	b := 13
	fmt.Printf("a:%d,%b\n", a, a)
	fmt.Printf("b:%d,%b\n", b, b)
	/*
		0011 1100
		0000 1101
	*/
	// 1.按位&
	res1 := a & b
	fmt.Printf("%d & %d = %d\n", a, b, res1)
	/*
		60 & 13 = 12
	*/

	// 2.按位｜
	res2 := a | b
	fmt.Printf("%d | %d = %d\n", a, b, res2)
	/*
		60 | 13 = 61
	*/

	// 3.异或^
	res3 := a ^ b
	fmt.Printf("%d ^ %d = %d\n", a, b, res3)

	// 4.位清空&^
	res4 := a &^ b
	fmt.Printf("%d &^ %d = %d\n", a, b, res4)
	/*
		60 &^ 13 = 48
	*/

	// 5.左移、右移
	c := 8
	res5 := c << 2
	res6 := c >> 2
	fmt.Println(res5, res6) // 32 2
}
