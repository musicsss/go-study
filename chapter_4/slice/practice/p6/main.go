// 编写一个函数， 原地将一个UTF-8编码的[]byte类型的slice中相邻的空格(参考unicode.IsSpace)替换成一个空格返回
package main

import (
	"fmt"
	"unicode"
)

func removeAdjacentRepeatingSpace(bytes []byte) []byte {
	index := 0
	for i := 1; i < len(bytes); i++ {
		if !(unicode.IsSpace(rune(bytes[i])) && unicode.IsSpace(rune(bytes[index]))) {
			index++
			bytes[index] = bytes[i]
		}
	}
	return bytes[:index+1]
}

func main() {
	bytes := []byte{'a', 'b', ' ', ' ', 'd', 'e'}
	fmt.Println(string(removeAdjacentRepeatingSpace(bytes)))
}
