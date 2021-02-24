package main

import "fmt"

func main() {
	/*
		数组指针：首先是一个指针，里面存储一个数组的地址
		*[4]Type

		指针数组：首先是一个数组，存储的数据类型是指针
		[4]*Type

		*[5]float64: 数组指针：一个存储了5个浮点类型数据的数组的指针
		*[3]string：数组指针：一个存储了3个字符串类型的数组的指针
		[3]*string：指针数组：一个存储3个字符串指针地址的数组
		[5]*float64：指针数组：一个存储5个float64类型数据指针地址的数组
		*[5]*float64：数组指针：一个数组的指针，存储了5个float类型的数据的指针地址的数组的指针
		*[3]*string：数组指针：一个数组的指针，存储了3个string类型的数据的指针地址的数组的指针
		**[4]string：数组指针：存储了4个字符串数据的数组的指针的指针
		**[4]*string：数组指针：存储了4个字符串的指针地址的数组，的指针的指针
	*/
	// 数组指针
	//1.创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("arr1的地址；%p\n", &arr1) // arr1的地址；0x1400001a040
	//2.创建一个指针，存储该数组的地址 --> 数组指针
	var p1 *[4]int
	p1 = &arr1                      // 数组长度必须为4，且为int
	fmt.Println(p1)                 // &[1 2 3 4]
	fmt.Println(*p1)                // 获取指针对应数组的值
	fmt.Printf("arr1的地址：%p\n", p1)  // arr1的地址：0x1400001a040
	fmt.Printf("p1自己的地址：%p\n", &p1) // p1自己的地址：0x1400000e030

	//3.根据数组指针，操作数组
	(*p1)[0] = 100
	fmt.Println(arr1) // [100 2 3 4]
	// 简化写法
	p1[0] = 200
	fmt.Println(arr1) // [200 2 3 4]

	// 指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Printf("arr2类型：%T\n", arr2) // arr2类型：[4]int
	fmt.Printf("arr3类型：%T\n", arr3) // arr3类型：[4]*int
	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)
	/*
		[100 2 3 4]
		1    a的值没变，因为修改的是arr2数组元素，a不会被修改
	*/
	//arr[3] := 0     报错，因为arr3是指针数组，内部存储的是指针，不能赋值int数据进去
	*arr3[0] = 200 // arr3[0]是指针，*arr3[0] 修改指针对应地址（a的地址），对应的数值：就是a变量的值 1
	fmt.Println(arr3)
	fmt.Println(a) // 200

	b = 1000
	fmt.Println(arr2) // [100 2 3 4]  arr2存储的是数值
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i]) // 获取a b c d地址对应的数值
		/*
			200
			1000    b的值已经被修改了
			3
			4
		*/
	}

}
