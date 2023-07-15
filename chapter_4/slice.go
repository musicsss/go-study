package main

import "fmt"

func main() {
	months := []string{1: "January", 12: "December"}

	months = append(months, "nil")
	fmt.Println(months)
}
