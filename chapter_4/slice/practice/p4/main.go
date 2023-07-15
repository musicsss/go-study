// 编写一个rotate函数， 实现通过一次循环完成旋转
package main

import "fmt"

func rotate(s []int, k int) {
	n := len(s)
	k %= n
	reverse(s[:k])
	reverse(s[k:])
	reverse(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	rotate(arr, 2)
	fmt.Println(arr)
}
