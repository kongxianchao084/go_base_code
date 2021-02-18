package main

import (
	"fmt"
)

func main() {
	/*
		map结合slice使用：
			1。创建map用于存放人信息
				name，age，sex，address
			2。每个map存储一个人信息
			3。将这些map存放到slice中
			4。打印遍历输出
	*/
	map1 := make(map[string]string)
	map1["name"] = "王二狗"
	map1["age"] = "30"
	map1["sex"] = "男性"
	map1["address"] = "北京市xx路xx号"

	map2 := make(map[string]string)
	map2["name"] = "李小华"
	map2["age"] = "20"
	map2["sex"] = "女性"
	map2["address"] = "上海市xx路xx号"

	map3 := map[string]string{"name": "ruby", "age": "30", "sex": "女性", "address": "杭州市"}

	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)

	for i, val := range s1 {
		fmt.Printf("第%d个人的信息是：\n", i+1)
		fmt.Printf("\t姓名：%s\n", val["name"])
		fmt.Printf("\t年龄：%s\n", val["age"])
		fmt.Printf("\t性别：%s\n", val["sex"])
		fmt.Printf("\t地址：%s\n", val["address"])
	}

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{4, 5, 6}
	fmt.Println(arr1, arr2)

	slice1 := make([]int, 1, 3)
	fmt.Println(slice1) // [0]
	slice2 := []int{7, 7, 8}
	fmt.Println(slice2)

	Map1 := make(map[string]string)
	Map2 := map[string]string{"Go": "222"} // 必须吃屎话
	fmt.Println(Map1, Map2)
}
