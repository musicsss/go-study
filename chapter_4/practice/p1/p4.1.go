// 编写一个函数，计算两个SHA256哈希码中不同bit的数目
package main

import (
	"crypto/sha256"
	"fmt"
)

func CountDifferentBits(c1, c2 [32]byte) int {
	count := 0
	for i := 0; i < len(c1); i++ {
		diff := c1[i] ^ c2[i] // 使用异或运算符获取不同的bit
		for diff != 0 {
			count++
			diff &= diff - 1 // 清除最右边的1，继续统计不同的bit
		}
	}
	return count
}

func main() {
	hash1 := sha256.Sum256([]byte("hello"))
	hash2 := sha256.Sum256([]byte("world"))

	diffBits := CountDifferentBits(hash1, hash2)
	fmt.Println(diffBits)
}
