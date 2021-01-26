package main

import "fmt"

func main() {
	/*
		switch语句
		语法结构：
			switch 变量名{
			case 数值1：分支1
			case 数值2：分支2
			case 数值3：分支3
			。。。
			default：
				最后一个分支
			}
		注意事项：
		1.switch可以作用域其他类型上，case后的数值必须和switch作用的变量类型一致
		2.case是无序的
		3.case后的数值是唯一的
	*/
	num := 5
	switch num {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	default:
		fmt.Println("请输入正确的季度")
	}
	fmt.Println("main...over...")

	// 模拟计算器
	num1 := 0
	num2 := 0
	opr := ""
	fmt.Println("请输入一个整数：")
	fmt.Scanln(&num1)
	fmt.Println("请在输入一个整数：")
	fmt.Scanln(&num2)
	fmt.Println("请输入一个操作：+、-、*、/")
	fmt.Scanln(&opr)

	switch opr {
	case "+":
		fmt.Printf("%d %s %d = %d", num1, opr, num2, num1+num2)
	case "-":
		fmt.Printf("%d %s %d = %d", num1, opr, num2, num1-num2)
	case "*":
		fmt.Printf("%d %s %d = %d", num1, opr, num2, num1*num2)
	case "/":
		fmt.Printf("%d %s %d = %d", num1, opr, num2, num1/num2)
	default:
		fmt.Println("请输入正确的操作符：+、-、/、*")
	}
}
