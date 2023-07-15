package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设置路由处理函数
	http.HandleFunc("/", handler)
	// 启动服务器并监听端口
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	// 将请求的URL路径返回给客户端
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
