package main

import (
	"fmt"
)

func main() {
	fmt.Println("main function called")

	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(2, 5)
}

func init() {
	fmt.Println("init function called before main")
}
