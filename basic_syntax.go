package main

import "fmt" // format package

func main() {
	name := "Abu Sayed" // var name string
	age := 20           // var age int
	num1 := 1
	num2 := 1
	const COUNTRY = "BD" // const variable
	fmt.Println("My Name is ", name)
	fmt.Println("My Age is ", age)
	fmt.Printf("%v is a student \n", name)
	fmt.Println(COUNTRY)

	fmt.Println("Enter you number")
	fmt.Scanf("%v %v\n", &num1, &num2)

	fmt.Printf("your num1 is %v, and num2 is %v\n", num1, num2)
}
