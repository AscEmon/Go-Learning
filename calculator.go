// package main

// import "fmt"

// func main() {
// 	i := true

// 	for i {
// 		var num1, num2 float64
// 		var operator string
// 		fmt.Println("Enter the first number: ")
// 		fmt.Scanf("%f\n", &num1)
// 		fmt.Println("Enter the second number: ")
// 		fmt.Scanf("%f\n", &num2)
// 		fmt.Println("Enter the operator: ")
// 		fmt.Scanf("%s\n", &operator)

// 		result := calculate(num1, num2, operator)
// 		if result != 0 {
// 			fmt.Println("The result is: ", result)
// 		}

// 	}

// }

// func calculate(num1 float64, num2 float64, operator string) float64 {
// 	switch operator {
// 	case "+":
// 		return num1 + num2
// 	case "-":
// 		return num1 - num2
// 	case "*":
// 		return num1 * num2
// 	case "/":
// 		return num1 / num2
// 	default:
// 		fmt.Println("Invalid operator")
// 		return 0
// 	}
// }
