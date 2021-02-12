package main

import "fmt"

func main() {
	/*
		数组array
			存储一组相同数据类型的数据结构
			特点：定长
		切片slice
			同数组类似，也叫做变长数组或者动态数组
			特点：变长
			是一个引用类型的容器，指向了一个顶层数组

		make()
			func make(t Type, size...IntegerType) Type
			第一个参数：类型 slice，map，chan
			第二个参数：长度len
				实际存储元素的数量
			第三个长度：容量cap
				最多能够存储的元素数量
	*/
	// 1.数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	// 2.切片
	var s1 = []int{1, 2, 3, 4} // 变长
	fmt.Println(s1)
	fmt.Printf("%T,%T\n", arr1, s1) // [4]int,[]int

	s2 := []int{4, 5, 6, 7}
	fmt.Println(s2)

	s3 := make([]int, 3, 8)
	fmt.Printf("类型：%T,长度：%d,容量：%d\n", s3, len(s3), cap(s3)) // 类型：[]int,长度：3,容量：8
	fmt.Println(s3)                                         // [0 0 0]
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3) // [1 2 3]

	//append()
	s4 := make([]int, 0, 5)
	fmt.Println(s4) // []
	s4 = append(s4, 1, 2)
	fmt.Println(s4) // [1 2]
	s4 = append(s4, 3, 4, 5, 6, 7)
	fmt.Println(s4) // [1 2 3 4 5 6 7]
	// 将另外一个切片值赋值给切片
	s4 = append(s4, s3...)
	fmt.Println(s4) // [1 2 3 4 5 6 7 1 2 3]

	// 遍历切片
	for i, v := range s4 {
		fmt.Println(i, "==>", v)
	}
	/*
		0 ==> 1
		1 ==> 2
		2 ==> 3
		3 ==> 4
		4 ==> 5
		5 ==> 6
		6 ==> 7
		7 ==> 1
		8 ==> 2
		9 ==> 3
	*/
}
