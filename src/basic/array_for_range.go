package main

import "fmt"

func main() {
	/*
		数组的遍历
			依次访问数组中的元素
			方法一：arr[0],arr[1],arr[2]...

			方法二：通过循环，配合下标
			for i:=0;i<len(arr);i++{
				arr[i]
			}

			方法三：使用range
			range,词义"范围"
			不需要操作数组的下标，到达数组的末尾，自动结束for range循环
			每次都数组中获取下标和对应的数值
	*/
	// for循环
	arr1 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// for。。。range。。。
	for i, v := range arr1 {
		fmt.Printf("下标是：%d，数值是：%d\n", i, v)
	}
	/*
		下标是：0，数值是：1
		下标是：1，数值是：2
		下标是：2，数值是：3
		下标是：3，数值是：4
		下标是：4，数值是：5
	*/
}
