package main

import "fmt"

func main() {
	/*
		slice := arr[start:end]
			切片中的数据: [start,end]
			arr[:end] 从开头到end
			arr[start:] 从start到末尾

		从已有的数组上，直接创建切片，该切片的底层数组就是当前的数组
			长度: start到end切割的数据长度
			容量：start到数组的末尾的长度
	*/
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("...1.已有数组直接创建切片...")
	s1 := a[:5]  // 1-5
	s2 := a[3:8] // 4-8
	s3 := a[5:]  // 6-10
	s4 := a[:]   // 1-10
	fmt.Println("a: ", a)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", s3)
	fmt.Printf("%p\n", s4)
	/*
		a: 0x14000130000
		s1: 0x14000130000
		s2: 0x14000130018
		s3: 0x14000130028
		s4: 0x14000130000
	*/

	fmt.Println("...2.长度和容量...")
	/*
		s1长度：start->end 为5
		s1容量：start->末尾 为10
	*/
	fmt.Printf("s1: len:%d,cap:%d\n", len(s1), cap(s1)) // s1: len:5,cap:10
	/*
		s2长度：start->end 为5
		s2容量：start->末尾 为3-10 7
	*/
	fmt.Printf("s2: len:%d,cap:%d\n", len(s2), cap(s2)) // s2: len:5,cap:7
	fmt.Printf("s3: len:%d,cap:%d\n", len(s3), cap(s3)) // s3: len:5,cap:5
	fmt.Printf("s4: len:%d,cap:%d\n", len(s4), cap(s4)) // s4: len:10,cap:10

	fmt.Println("...3.更改数组的内容...")
	a[4] = 100
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", s3)
	fmt.Printf("%p\n", s4)
	/*
		切片是引用类型：修改一个都会改变，因为底层是同一个数组
		[1 2 3 4 100 6 7 8 9 10]
		[1 2 3 4 100]
		[4 100 6 7 8]
		[6 7 8 9 10]
		[1 2 3 4 100 6 7 8 9 10]
		0x14000020050
		0x14000020050
		0x14000020068
		0x14000020078
		0x14000020050
	*/

	fmt.Println("...3.更改数组的内容...")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", s3)
	fmt.Printf("%p\n", s4)
	/*
		[1 2 3 4 100 1 1 1 1 10]
		[1 2 3 4 100 1 1 1 1]
		[4 100 1 1 1]
		[1 1 1 1 10]
		[1 2 3 4 100 1 1 1 1 10]
		0x1400013a000
		0x1400013a000
		0x1400013a018
		0x1400013a028
		0x1400013a000

	*/

	fmt.Println("...4.扩容...")
	s1 = append(s1, 2, 2, 2, 2, 2)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s2)
	/*
		扩容指向了新的底层数组
		[1 2 3 4 100 1 1 1 1 10]
		[1 2 3 4 100 1 1 1 1 2 2 2 2 2]
		[4 100 1 1 1]
		[1 1 1 1 10]
		[1 2 3 4 100 1 1 1 1 10]
	*/
}
