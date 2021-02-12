package main

import "fmt"

func main() {
	/*
		深拷贝：拷贝的是数据本身
			值类型的数据，默认都是深拷贝：array、int、float、string、bool、struct
		浅拷贝：拷贝的是数据地址
			导致多个变量指向同一块内存
			引用类型的数据，默认都是浅拷贝：slice、map
		因为切片是引用类型的数据，直接拷贝的是地址
	*/

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0)
	for i := 0; i < len(s1); i++ { // s2扩容，会重新生成新的底层数组，地址和s1是不一样的
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Printf("s1:%p\n", s1) //s1:0x1400001a060
	fmt.Println(s2)
	fmt.Printf("s2:%p\n", s2) //s2:0x1400001a080

	s1[0] = 100
	fmt.Println(s1) // [100 2 3 4]
	fmt.Println(s2) // [1 2 3 4]

	fmt.Println("-----copy-------")
	// copy() 深拷贝
	s3 := []int{7, 8, 9}
	fmt.Println(s2)
	fmt.Println(s3)
	// 将s2拷贝到s3
	//fmt.Println("-----copy s2->s3 -------")
	//copy(s3,s2)
	//fmt.Println(s2)  // [1 2 3 4]
	//fmt.Println(s3)  // [1 2 3 4]
	// 将s3拷贝到s2
	// 因为s2长度大于s3,因此随便装
	fmt.Println("-----copy s3->s2 -------")
	copy(s2, s3)
	fmt.Println(s2)           // [1 2 3 4]
	fmt.Println(s3)           // [1 2 3 4]
	fmt.Printf("s2:%p\n", s2) //s2:0x140000aa020
	fmt.Printf("s3:%p\n", s3) //s3:0x140000a2018

}
