package main

import "fmt"

func main() {
	/*
		多维数组
	*/
	arr1 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(arr1)         // [[1 2 3 4] [5 6 7 8] [9 10 11 12]]
	fmt.Println(arr1[2][3])   // 12 第三个一维数组的第4个值
	fmt.Println(len(arr1))    // 3 二维数组的长度
	fmt.Println(len(arr1[0])) // 4 第一个一维数组的长度

	// for。。。range
	for i, v := range arr1 {
		fmt.Println(i, "===", v)
	}
	/*
		0 === [1 2 3 4]
		1 === [5 6 7 8]
		2 === [9 10 11 12]
	*/
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Println(arr1[i][j])
		}
	}
}
