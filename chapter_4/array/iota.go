package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

var symbol = [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

func main() {

	fmt.Println(RMB, symbol[RMB])
}
