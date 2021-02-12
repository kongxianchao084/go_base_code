package main

import "fmt"

func main() {
	/*
		数组  定长
			1：概念：存储一组相同数据类型的数据结构
			2：语法
				var 数组名 [长度] 数据类型
				var 数组名 = [长度] 数据类型{元素1，元素2。。。}
				数组名 := [...]数组类型{元素1，元素2。。。。}
			3：通过下标访问 从0开始访问
			4：长度和容量 相等关系，区别于切片

		数组内存分析：
			数组开辟完内存地址，会存放零值(通过下标查找)
			数组的内存地址是数组的首地址，因为数组的地址是连续的，因此效率很高
	*/

	// step1：创建数组
	var arr1 [4]int
	// step2:数组的访问
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	//fmt.Println(arr1[4]) invalid array index 4 (out of bounds for 4-element array)

	fmt.Println("数组的长度：", len(arr1)) // 容器中实际存储的数据量
	fmt.Println("数组的容量：", cap(arr1)) // 容器中能够存储的最大的数量
	//因为数组定长，长度和容量相同
	arr1[0] = 100             // 可以通过下标来修改数组元素
	fmt.Printf("%p\n", &arr1) // 0x14000194000
	fmt.Println(arr1[0])

	// 数组的其他创建方式
	var a [4]int   // 同 var a = [4] int
	fmt.Println(a) // [0 0 0 0]

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b) // [1 2 3 4]

	// 对应位值可以为零
	var c = [5]int{1, 2, 4}
	fmt.Println(c) // [1 2 4 0 0]

	// 指定对应下标存储值
	var d = [5]int{1: 1, 3: 2}
	fmt.Println(d) // [0 1 0 2 0]

	// 字符串类型零值
	var e = [5]string{"rose", "王二狗", "ruby"}
	fmt.Println(e) // [rose 王二狗 ruby  ]

	// 不指定数组长度
	f := [...]int{1, 2, 3, 4, 5}
	fmt.Println(f)      // [1 2 3 4 5]
	fmt.Println(len(f)) // 5
	g := [...]int{1: 3, 6: 5}
	fmt.Println(g)      // [0 3 0 0 0 0 5]
	fmt.Println(len(g)) // 7

}
