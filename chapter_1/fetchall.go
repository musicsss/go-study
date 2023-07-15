package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()     // 记录程序开始时间
	ch := make(chan string) // 创建一个字符串类型的通道
	// 并发地调用fetch函数，将结果发送到通道ch中
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	// 从通道ch中接收结果并打印
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	// 打印程序运行时间
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch函数发送HTTP GET请求并将结果发送到通道ch中
func fetch(url string, ch chan<- string) {
	start := time.Now() // 记录请求开始时间
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 发送错误信息到通道ch中
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 关闭响应体，防止资源泄漏
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // 发送错误信息到通道ch中
		return
	}
	secs := time.Since(start).Seconds()                  // 计算请求的时间
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) // 发送结果到通道ch中
}
