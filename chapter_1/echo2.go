// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 每次循环迭代， range产生一对值；索引以及在该索引处的元素值。
	// 这个例子不需要索引，但 range 的语法要求，要处理元素，必须处理索引。一种思路是把索引赋值给一个临时变量（如 temp）然后忽略它的值，
	// 但 Go 语言不允许使用无用的局部变量（local variables），因为这会导致编译错误。
	// Go 语言中这种情况的解决方法是用 空标识符（blank identifier），即 _（也就是下划线）。空标识符可用于在任何语法需要变量名但
	// 程序逻辑不需要的时候（如：在循环里）丢弃不需要的循环索引，并保留元素值。
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
