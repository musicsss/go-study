// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作
package main

import "fmt"

func removeAdjacentDuplicates(strings []string) []string {
	index := 0
	for i := 1; i < len(strings); i++ {
		if strings[i] != strings[index] {
			index++
			strings[index] = strings[i]

		}
	}

	return strings[:index+1]
}

func main() {
	array := []string{"hello", "hello", "good", "nice", "yes", "yes"}

	fmt.Println(removeAdjacentDuplicates(array))
}
