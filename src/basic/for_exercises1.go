package main

import "fmt"

func main() {
	/*
		for循环练习题
	*/
	// 1：打印58-23数字
	for i := 28; i <= 58; i++ {
		fmt.Println(i)
	}

	// 2：求1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("1-100的和：%d\n", sum)

	// 3：打印1-100内，能够被3整除，但是不能被5整除的数字，统计被打印数字的个数，每行打印5个
	num := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			num += 1
			fmt.Print(i, " ")
			if num%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Println(num)

	fmt.Println("====打印99乘法表====")
	// 4：打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Printf("%dx%d=%d\n", i, j, i*j)
			} else {
				fmt.Printf("%dx%d=%d ", i, j, i*j)
			}
		}
	}

	fmt.Println("标准代码")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		// 直接打印一个换行 简洁
		fmt.Println()
	}
}
