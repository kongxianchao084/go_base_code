package main

import "fmt"

func main() {
	/*
		存储特点
			A：存储的是无序的键值对
			B：键不能重复，并且和value值一一对应的
				map中key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错
	*/

	// 1.创建map
	var map1 map[int]string // map[key类型]值类型
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go": 98, "Python": 87}
	rating := map[string]float64{"Go": 5}
	fmt.Println(map1) // map[] 未初始化 nil map 不能存放键值对
	fmt.Println(map2) // map[] 不是nil map
	fmt.Println(map3) // map[Go:98 Python:87]
	fmt.Println(rating)

	// nil map不能存放键值对
	//map1[0] = "hello"  panic: assignment to entry in nil map

	fmt.Println(map1 == nil) // true
	fmt.Println(map2 == nil) // false

	// 3 存储键值对到map
	if map1 == nil {
		map1 = make(map[int]string)
	}

	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "王二狗"
	fmt.Println(map1)
	fmt.Printf("%T\n", map1)

	// 4.获取数据
	// 根据key获取对应的value，如果key存在，获取value数值；如果key不存在，获取value对应数据类型的零值
	fmt.Println(map1[3])
	fmt.Println(map1[5]) // "" 空字符串

	v1, ok := map1[3]
	if ok {
		fmt.Println("对应的数值是：", v1)
	} else {
		fmt.Println("操作的key不存在，获取到的是零值：", v1)
	}

	// 5.修改数据
	// 如果key不存在，则添加数据

	// 6.删除数据
	// 如果key不存在，则对map没影响
	delete(map1, 3)
	delete(map1, 30)

	// 7.map长度  len函数 对应key个数

}
