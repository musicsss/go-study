// Go语言的代码通过包(package)组织，包类似于其它语言里的库（libraries）或者模块（modules）。
// 一个包由位于单个目录下的一个或多个 .go 源代码文件组成，目录定义包的作用。每个源文件都以一条 package 声明语句开始，
// 这个例子里就是 package main，表示该文件属于哪个包，紧跟着一系列导入（import）的包，之后是存储在这个文件里的程序语句。
package main

import "fmt"

// main包比较特殊。它定义了一个独立可执行的程序，而不是一个库。
// 在main里的main()函数也很特殊，它是整个程序执行时的入口。

// 一个函数的声明由 func 关键字、函数名、参数列表、返回值列表
//（这个例子里的 main 函数参数列表和返回值都是空的）以及包含在大括号里的函数体组成
func main() {
	// fmt包 包含了格式化输出、接收输入的函数。
	// Println 是其中一个基础函数，可以打印以空格间隔的一个或多个值，并在最后添加一个换行符，从而输出一整行。
	fmt.Println("Hello, World")
}
