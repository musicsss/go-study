// Dup1 print the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map 存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作。
	// 键可以是任意类型，只要其值能用 == 运算符比较，最常见的例子是字符串；值则可以是任意类型。
	// 这个例子中的键是字符串，值是整数。
	// 内置函数 make 创建空 map，此外，它还有别的作用。
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// map 中不含某个键时不用担心，首次读到新行时，等号右边的表达式 counts[line] 的值将被计算为其类型的零值，对于 int 即 0。
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	// 每次迭代得到两个结果，键和其在 map 中对应的值。map 的迭代顺序并不确定，从实践来看，该顺序随机，每次运行都会变化
	for line, n := range counts {
		// if 语句条件两边也不加括号，但是主体部分需要加。if 语句的 else 部分是可选的，在 if 的条件为 false 时执行
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
