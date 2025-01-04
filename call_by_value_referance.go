package main

import "fmt"

func callbyvalue(x int) {
	x = 20
}

func callByReference(x *int) {
	*x = 20
}

func main() {
	x := 10

	fmt.Println("Before calling function x is: ", x)
	callbyvalue(x)
	fmt.Println("After calling callbyvalue x is: ", x)

	fmt.Println("Calling callByReference x is: ", x)
	callByReference(&x)
	fmt.Println("After Calling callByReference x is: ", x)

}
