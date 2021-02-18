package main

import "fmt"

func main() {
	getSum(20)
	getSum(10)
	getAdd(20, 30)
}

func getSum(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和是：%d\n", n, sum)
}

func getAdd(a, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}
