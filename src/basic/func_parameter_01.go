package main

import "fmt"

func main() {
	/*
		参数传递过程：
		数据类型：
			1：按照数据类型来分
				基本数据类型：
					int、float、string、bool
				符合数据类型：
					array、slice、map、struct、interface
			2：按照数据的存储特点来份
				值类型的数据：操作的是数据本身
					int、float、bool、string、array
				引用类型的数据：操作的是数据的地址
					slice、map、chan

		参数传递：
			A：值传递：传递的是数据的副本，修改数据，对于原始的数据没有影响的
				值类型的数据，默认都是值传递：基础类型，array，struct
			B：引用传递：传递的是数据的地址，修改数据，原始数据也会发生改变

	*/
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前，数组的数据：", arr1)
	func1(arr1)
	fmt.Println("函数调用后，数组的数据：", arr1)
	/*
		值传递：修改数组后会开辟新的空间
		函数调用前，数组的数据： [1 2 3 4]
		函数中，数组的数据： [1 2 3 4]
		函数中，数组的数据修改后： [100 2 3 4]
		函数调用后，数组的数据： [1 2 3 4]
	*/
	fmt.Println("------------------------")
	s1 := []int{1, 2, 3, 4}
	fmt.Println("函数调用前，切片的数据：", s1)
	func2(s1)
	fmt.Println("函数调用后，切片的数据：", s1)
	/*
		函数调用前，切片的数据： [1 2 3 4]
		函数中，切片的数量： [1 2 3 4]
		函数中，切片修改后的数据： [100 2 3 4]
		函数调用后，切片的数据： [100 2 3 4]
	*/
}

// 值传递：开辟新的空间
func func1(arr2 [4]int) { // arr2 = arr1 将arr1的值复制给arr2，开辟新的内存
	fmt.Println("函数中，数组的数据：", arr2)
	arr2[0] = 100
	fmt.Println("函数中，数组的数据修改后：", arr2)
}

// 引用传递：传递的地址
func func2(s2 []int) { // s2 = s1 将s1的地址传递给s2 修改s2 s1的数值也会改变
	fmt.Println("函数中，切片的数量：", s2)
	s2[0] = 100
	fmt.Println("函数中，切片修改后的数据：", s2)
}
