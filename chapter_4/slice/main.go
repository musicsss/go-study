package main

import "fmt"

func main() {
	months := [...]string{1: "January", 12: "December", 22: "nil"}

	fmt.Println(len(months[6:9]), cap(months[6:9]))
}
