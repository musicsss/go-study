package main

import "fmt"

func main() {
	defer fmt.Println("defer func")
	fmt.Println("hello world!")
	defer fmt.Println("second defer func")
	fmt.Println("nice to meet you!")
}
