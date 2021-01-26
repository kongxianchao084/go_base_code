package main

import "fmt"

func main() {
	/*
		switch其他写法：
		1: 省略switch后的变量，相当于直接作用于true上 布尔类型

		2:case后可以跟多个数值:满足一个即可
		switch 变量{
		case 数值1,数值2,数值3:
		case 数值4,数值5:
		}

		3:switch初始化：注意作用域即可
		switch 初始化语句;变量{
		}
	*/
	/*
		练习：
		成绩：
		[0-59] 不及格
		[60-69] 及格
		[70-79] 中
		[80-89] 良好
		[90-100] 优秀
	*/
	score := 80
	switch {
	case score < 60:
		fmt.Println("不及格")
	case 60 <= score && score < 70:
		fmt.Println("及格")
	case 70 <= score && score < 80:
		fmt.Println("中")
	case 80 <= score && score < 90:
		fmt.Println("良好")
	case 90 <= score && score <= 100:
		fmt.Println("优秀")
	default:
		fmt.Println("操作有误")
	}

	fmt.Println("=============")
	/*
		一个月天数
		1，3，5，7，8，10，12   31
		4，6，9，11 30
		2  28/29
	*/
	year := 2021
	month := 2
	day := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		/*
			普通闰年：是4的倍数，不是100的倍数
			世纪闰年：是400的倍数
		*/
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		fmt.Println("请输入正确的月份")
	}
	fmt.Printf("%d-%d 的天数：%d\n", year, month, day)

	fmt.Println("=============")
	// 注意作用域
	switch language := "golang"; language {
	case "golang":
		fmt.Println("Go语言")
	case "java":
		fmt.Println("Java语言")
	case "python":
		fmt.Println("Python语言")
	default:
		fmt.Println("请输入正确的语言类型")
	}
}
