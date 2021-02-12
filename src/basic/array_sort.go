package main

import "fmt"

func main() {
	/*
		冒泡排序
			依次排序相邻的元素，将大的移到右边
			第一轮：最大的数值已经移到最右边
			第二轮：倒数第二大的数值移到倒数第二位位置
			。。。 依次类推
	*/

	arr1 := [5]int{15, 23, 8, 10, 7}
	// 循环轮数:一共5个数，循环4次即可 1<=i<=4
	for i := 1; i < len(arr1); i++ {
		// 每一轮相邻数值对比
		for j := 0; j < len(arr1)-i; j++ {
			if arr1[j] > arr1[j+1] {
				arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
			}
		}
	}
	fmt.Println(arr1)
}
