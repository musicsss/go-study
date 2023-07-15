// Boiling prints the boiling point of water
// 每个源文件中以包的声明语句开始，说明该源文件是属于哪个包。
package main

// 包声明语句之后是import语句导入依赖的其它包，
// 然后是包一级的类型、变量、常量、函数的声明语句，包一级的各种类型的声明语句的顺序无关紧要。
import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %gC\n", f, c)
}
