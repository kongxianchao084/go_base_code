package main

import "fmt"

func main() {
	/*
		逻辑运算符：操作数必须是bool，运算结果也是bool
		逻辑与：&&
		逻辑或：||
		逻辑非：!
	*/

	// 逻辑与
	f1 := true
	f2 := false
	f3 := true
	f4 := false
	res1 := f1 && f2
	fmt.Printf("res1: %t\n", res1)
	res2 := f1 && f2 && f3
	fmt.Printf("res2: %t\n", res2)
	/*
		res1: false
		res2: false
	*/

	// 逻辑或
	res3 := f1 || f2
	res4 := f2 || f4
	fmt.Printf("res3: %t\n", res3)
	fmt.Printf("res4: %t\n", res4)
	/*
		res3: true
		res4: false
	*/

	// 逻辑非
	fmt.Printf("f1:%t,!f1:%t\n", f1, !f1)
	fmt.Printf("f2:%t,!f2:%t\n", f2, !f2)
	/*
		f1:true,!f1:false
		f2:false,!f2:true
	*/
}
