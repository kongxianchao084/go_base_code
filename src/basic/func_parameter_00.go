package main

import "fmt"

func main() {
	/*
		可变参数
			语法：
				参数名 ... 参数类型
			实质：nums是切片(长度可变的数组)
		注意事项：
			A：如果一个函数的参数是可变参数，同时还有其他的参数，可变参数要放在
				参数列表的最后
			B：一个函数的参数列表中最多只能有一个可变参数
	*/
	getSum(1, 2, 3)

	s1 := []int{1, 2, 3, 4, 5}
	getSum(s1...)
}

func getSum(nums ...int) {
	// fmt.Printf("%T\n",nums) // []int
	fmt.Println(nums)
	/*
		[1 2 3]
		[1 2 3 4 5]
	*/
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
}
