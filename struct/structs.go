package main

import "fmt"

type Person struct {
	name    string
	age     int
	college string
	phone   int
}

func main() {
	person1 := Person{
		name:    "prashant",
		age:     18,
		college: "LNCT",
		phone:   12345,
	}
	fmt.Printf("name: %s, age: %d, college: %s, phone: %d", person1.name, person1.age, person1.college, person1.phone)

}
