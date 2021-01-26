package main

import "fmt"

func main() {
	/*
		条件语句：if  只能作用于bool类型
		语法格式：
			1> 写法1
			if 条件表达式{
			//	布尔表达式为true时执行
			}
			2> 写法2
			if 条件表达式{
			//	布尔表达式为true时执行
			} else {
			//	布尔表达式为false时执行
			}
			3> 写法3
			if 条件1 {
			//	条件1成立
			} else if 条件2 {
			// 条件1不成立，条件2成立
			} else if 条件3 {
			//	条件1不成立，条件2不成立，条件3成立
			} else {
			// 上述所有条件均不成立
			}
	*/
	fmt.Println("================")

	//1.给定一个数字，如果大于10，输出打印这个数字大于10
	num := 9
	if num > 10 {
		fmt.Println("大于10")
	}
	fmt.Println("main...over...")

	fmt.Println("================")

	//2.给定一个成绩，如果大于等于60，就是及格，否则就是不及格
	score := 0
	fmt.Println("请输入您的成绩：")
	fmt.Scanln(&score)
	if score >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}
	fmt.Println("main...over...")

	fmt.Println("================")

	//3.给定性别，如果是男，就去男厕所，否则就去女厕所
	//男 true 女 false
	sex := true
	if sex == true {
		fmt.Println("去男厕所")
	} else if sex == false {
		fmt.Println("去女厕所")
	} else {
		fmt.Println("请输入正确的性别")
	}
	fmt.Println("main...over...")

	fmt.Println("================")

}
