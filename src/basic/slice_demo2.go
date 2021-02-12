package main

import "fmt"

func main() {
	/*
		slice:
			1.每一个切片引用了一个底层数组
			2.切片本身不存储任何数据，都是这个底层数组存储，所以修改切片也就是修改这个数组中的数据
			3.当向切片中添加数据，如果没有超过容量，直接添加，如果超过容量，自动扩容（成倍增长）
			4.切片一旦扩容，就是重新指向一个新的底层数组
	*/
	s1 := []int{1, 2, 3}
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:3,cap:3
	fmt.Printf("%p\n", s1)                          // 0x14000132018 这里s1不要加&,加了取的就是底层数组的地址，不是切片的地址

	// 扩容：指向新的底层数组 地址发生改变
	// 扩容：容量成倍数增加
	s1 = append(s1, 4, 5)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:5,cap:6
	fmt.Printf("%p\n", s1)                          // 0x14000122030
}
