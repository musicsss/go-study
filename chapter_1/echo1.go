// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	// os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。
	// 程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量
	// os.Args 变量是一个字符串（string）的 切片（slice）
	"os"
)

func main() {
	// var 声明定义了两个 string 类型的变量 s 和 sep。变量会在声明时直接初始化。
	// 如果变量没有显式初始化，则被隐式地赋予其类型的 零值（zero value），数值类型是 0，字符串类型是空字符串 ""。
	var s, sep string
	// os.Args 的第一个元素：os.Args[0]，是命令本身的名字
	// for initialization; condition; post { ...}
	// for 循环三个部分不需括号包围。大括号强制要求，左大括号必须和 post 语句在同一行。
	// initialization 语句是可选的，在循环开始前执行。initalization 如果存在，必须是一条 简单语句（simple statement），
	// 即，短变量声明、自增语句、赋值语句或函数调用。condition 是一个布尔表达式（boolean expression），
	// 其值在每次循环迭代开始时计算。如果为 true 则执行循环体语句。post 语句在循环体执行结束后执行，之后再次对 condition 求值。condition 值为 false 时，循环结束。
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
