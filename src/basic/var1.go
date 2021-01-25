package main

import "fmt"

/*
内容：变量的声明和赋值、使用
*/

func main() {
	/*
		变量：一小块内存，用于存储数据，在程序运行过程中可以改变

		使用：
			step1：变量的声明，也叫定义
			step2：变量的访问，赋值和取值
		go：静态语言：强类型语言
	*/
	//第一种：定义变量，然后进行赋值
	var num1 int
	num1 = 30
	fmt.Printf("num1的数值是：%d\n", num1)
	// 写在一行
	var num2 int = 15
	fmt.Printf("num2的数值是：%d\n", num2)

	// 第二种：类型推断
	var name = "王二狗"
	fmt.Printf("类型是：%T，数值是：%s\n", name, name)

	// 第三种：简短定义，也叫简短声明 省略var
	num := 300
	fmt.Println(num)

	// 多个变量声明
	// 1:先定义，在赋值
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n int = 200, 100
	fmt.Println(m, n)

	// 2：类型推断
	var n1, f1, s1 = 100, 3.14, "Go"
	fmt.Println(n1, f1, s1)

	var (
		studentName = "李小花"
		age         = 18
		sex         = "女"
	)
	fmt.Printf("学生姓名：%s，年龄：%d，性别：%s\n", studentName, age, sex)
}
