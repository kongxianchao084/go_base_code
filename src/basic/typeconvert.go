package main

import "fmt"

func main() {
	var a int8
	a = 10

	var b int16
	b = int16(a)
	fmt.Println(a, b) // 10 10

	f1 := 4.83
	var c int
	c = int(f1)
	fmt.Println(f1, c) // 4.83 4

	//b1 := true
	//a = int8(b1) cannot convert b1 (type bool) to type int8
}
