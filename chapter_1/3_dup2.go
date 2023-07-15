// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open 函数返回两个值。第一个值是被打开的文件（*os.File），其后被 Scanner 读取。
			// os.Open 返回的第二个值是内置 error 类型的值。如果 err 等于内置值nil那么文件被成功打开。
			// 读取文件，直到文件结束，然后调用 Close 关闭该文件，并释放占用的所有资源。
			// 相反的话，如果 err 的值不是 nil，说明打开文件时出错了。这种情况下，错误值描述了所遇到的问题。
			// 我们的错误处理非常简单，只是使用 Fprintf 与表示任意类型默认格式值的动词 %v，向标准错误流打印一条信息，
			// 然后 dup 继续处理下一个文件；continue 语句直接跳到 for 循环的下个迭代开始执行。
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
