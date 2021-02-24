package main

import "fmt"

func main() {
	/*
		指针：存储了另外一个变量的内存地址的变量
	*/

	// 1.定义一个int类型的变量
	a := 10
	fmt.Println("a的数值是：", a)
	fmt.Printf("a的地址是：%p\n", &a) // a的地址是：0x14000018050
	//2 创建一个指针变量，存储变量a的地址
	var p1 *int
	fmt.Println(p1)                           // <nil> 空指针
	p1 = &a                                   // p1指向了a的地址
	fmt.Println("p1的数值：", p1)                 // p1的数值： 0x14000018050  就是a的地址
	fmt.Printf("p1自己的地址：%p\n", &p1)           // p1自己的地址：0x1400000e030
	fmt.Println("p1的数值，是a的地址，该地址存储的数据：", *p1) // 获取p1存储地址，对应的数值 10

	// 3 操作变量，更改数值，并不改变地址
	a = 100
	fmt.Printf("%p\n", &a)

	// 4 通过指针，改变变量的数值
	*p1 = 200
	fmt.Println(a) // 200

	// 5 指针的指针
	var p2 **int    // 将p1的地址存储到p2
	fmt.Println(p2) // <nil>
	p2 = &p1
	fmt.Printf("%T，%T，%T\n", a, p1, p2)               // int，*int，**int
	fmt.Println("p2的数值：", p2)                         // p2的数值： 0x1400000e030
	fmt.Printf("p2的地址：%p\n", &p2)                     // p2的地址：0x1400000e038
	fmt.Println("p2中存储的地址，对应的数值，就是p1的地址对应的数据：", *p2)  // a的地址
	fmt.Println("p2中存储的地址，对应的数值，就是p1的地址对应的数据：", **p2) // 200

}
