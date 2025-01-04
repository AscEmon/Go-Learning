package main

import (
	"fmt"
	"slices"
)

type Member struct {
	id   int
	name string
	age  int
}

func displayInfo(member Member) {
	fmt.Println("Display Information")
	fmt.Printf("ID = %v Name = %v Age = %v\n", member.id, member.name, member.age)
}

func (m *Member) increaseAge(age int) {
	if m.id == 1 {
		m.age = m.age + age
	}

}
func main() {

	member1 := Member{2, "Abu Sayed", 20}
	member2 := Member{1, "Boka Chuda", 30}
	displayInfo(member1)
	member1.increaseAge(25)
	displayInfo(member1)

	displayInfo(member2)
	member2.increaseAge(40)
	displayInfo(member2)

	members := []Member{
		{1, "Abu Sayed", 20},
		{2, "Boka Chuda", 30},
		{3, "Boka Chuda", 40},
	}

	fmt.Println(members)

	members = append(members, Member{4, "Boka Chuda", 50})
	slices.Reverse(members)
	fmt.Println(members)

}
