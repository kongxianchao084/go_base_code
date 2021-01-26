package main

import "fmt"

func main() {
	/*
		break:跳出整个循环
		continue:跳出当前循环
	*/
	fmt.Println("====break=======")
	for i := 1; i <= 9; i++ {
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
	/*
		====break=======
		1
		2
		3
		4
		5
		6
	*/

	fmt.Println("====continue=======")
	for j := 1; j <= 9; j++ {
		if j == 7 {
			continue
		}
		fmt.Println(j)
	}
	/*
		====continue=======
		1
		2
		3
		4
		5
		6
		8
		9
	*/
}
