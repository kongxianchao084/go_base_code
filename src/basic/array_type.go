package main

import "fmt"

func main() {
	/*
		数值的数据类型：值类型
			[size]type
		值类型：理解为存储的数值本身
			将数据传递给其他变量，传递的是数据的副本（备份）
			int，float，string，bool，array
	*/

	//1.数据类型
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [3]float64{2.15, 3.18, 6.19}
	arr3 := [4]int{5, 6, 7, 8}
	arr4 := [2]string{"hello", "world"}
	fmt.Printf("%T\n", arr1) // [4]int
	fmt.Printf("%T\n", arr2) // [3]float64
	fmt.Printf("%T\n", arr3) // [4]int
	fmt.Printf("%T\n", arr4) // [2]string

	//2.赋值
	num := 10
	num2 := num // 值传递
	num2 = 20
	fmt.Println(num, num2) // 10 20

	// 3.数组
	arr5 := arr1      // 值传递：修改arr5，不会影响arr1数值
	fmt.Println(arr5) // [1 2 3 4]
	arr5[0] = 2
	fmt.Println(arr1, arr5) // [1 2 3 4]  [2 2 3 4]

	// 4 使用 == 比较大小
	arr6 := [4]int{1, 2, 3, 4}
	fmt.Println(arr6 == arr1) // 比较数据下标对应的数值是否相等
	//fmt.Println(arr1 == arr2) invalid operation: arr1 == arr2 (mismatched types [4]int and [3]float64)

}
