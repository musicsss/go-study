// 编写一个程序wordfreq程序， 报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 设置输入分隔符为单词
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// 创建一个map来存储单词频率
	wordFreq := make(map[string]int)

	// 逐个读取单词并增加频率计数
	for scanner.Scan() {
		word := scanner.Text()
		wordFreq[word]++
	}

	// 打印每个单词和对应的频率
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
