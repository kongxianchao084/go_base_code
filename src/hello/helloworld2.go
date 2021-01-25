package main

import (
	// 包名称太长，可以起别名
	f "fmt"
	// _ 操作，只调用包中init函数，不使用包内函数
	// _ "github.com/xxx/yyy"
)

/*
go源码文件：
1.命令源码文件
声明自己属于main代码包，包含无参数声明和结果声明的main函数
命令源码文件被安装以后，GOPATH如果只有一个工作区，那么相应的可执行文件会被存放当前工作区的bin文件夹下；如果有多个工作区，就会安装到GOBIN
->指向的目录下
命令源码文件是Go程序的入口
命令源码文件应该单独放在一个代码包

多个工作区
1 go run可以运行
2 go build/go install 都会报错：这也证明了，多个命令源码文件可以单独go run，但是不能go build/go install
./helloworld2.go:22:6: main redeclared in this block
	previous declaration at ./helloworld.go:13:6
2.库源码文件
库源码文件就是不具备上述两个特征的源码文件。存在于某个代码包中的普通的源码文件。
库源码文件被安装后，相应的归档文件（.a）会被存放到当前工作区的pkg平台相关目录下
*/
func main() {
	// 打印并换行
	f.Println("我是第二个hello world文件.")
}
