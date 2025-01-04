package main

import "fmt"

type Student struct {
	id   int
	name string
}

func main() {

	teacher := make(map[int]Student)

	teacher[1] = Student{1, "Teacher :: Abu Sayed"}
	fmt.Println(teacher)
	students := map[int]Student{}

	fmt.Println(students)

	students[3] = Student{3, "Boka Chuda"}

	fmt.Println(students)

	delete(students, 2)

	fmt.Println(students)

	for k, v := range students {
		fmt.Println(k, v.name)
	}

}
