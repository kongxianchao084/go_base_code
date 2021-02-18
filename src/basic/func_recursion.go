package main

import "fmt"

func main() {
	/*
		递归函数：一个函数自己调用自己
	*/
	res := getSum(5)
	fmt.Println(res)

	// 2.fibonacci数列
	/*
		1 2 3 4 5 6 7  8  9  10 。。。
		1 1 2 3 5 8 13 21 34 55。。。
	*/
	res1 := getFibonacci(8)
	fmt.Println(res1)
}

func getSum(n int) int {
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}

func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFibonacci(n-1) + getFibonacci(n-2)
}
