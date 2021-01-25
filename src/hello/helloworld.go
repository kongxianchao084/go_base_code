package main

import (
	// 包名称太长，可以起别名
	f "fmt"
	// _ 操作，只调用包中init函数，不使用包内函数
	// _ "github.com/xxx/yyy"
)

/*
go run helloworld.go     直接运行go程序
go build helloworld.go   在$GOPATH/bin下生成二进制文件helloworld
*/
func main() {
	// 打印并换行
	f.Println("hello world")
	// 打印不换行
	f.Print("hello world")
}
