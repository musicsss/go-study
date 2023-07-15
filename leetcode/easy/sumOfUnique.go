// 1748. 唯一元素的和
// 给你一个整数数组 nums 。数组中唯一元素是那些只出现 恰好一次 的元素。
//
// 请你返回 nums 中唯一元素的 和 。
package main

import "fmt"

func sumOfUnique(nums []int) int {

	intMap := make(map[int]int)
	sum := 0
	for _, i := range nums {
		intMap[i]++
		if intMap[i] == 1 {
			sum += i
		} else if intMap[i] == 2 {
			sum -= i
		}
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumOfUnique(nums))
}
