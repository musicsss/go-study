package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{3, 2, 1}

	fmt.Println(arr1 == arr2)
}
