package main

import "fmt"

func main() {
	/*
		map遍历
			for循环
			for。。。range
	*/

	map1 := map[int]string{1: "红孩儿", 2: "小钻风"}
	map2 := make(map[int]string)
	map2[1] = "白骨精"
	map2[2] = "白素贞"

	// 1。遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println("---------------")
	for i := 1; i <= len(map2); i++ {
		fmt.Println(i, map2[i])
	}
}
