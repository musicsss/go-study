package main

import (
	"fmt"
	"go-study/chapter_6/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
}
