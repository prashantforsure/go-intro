package main

import "fmt"

type Person struct {
	name    string
	age     int
	college string
	phone   int
}
type Address struct {
	city    string
	state   string
	country string
}

type Man struct {
	name    string
	gender  bool
	address Address
}

func main() {
	person1 := Person{
		name:    "prashant",
		age:     18,
		college: "LNCT",
		phone:   12345,
	}
	man1 := Man{
		name:   "adarsh",
		gender: true,
		address: Address{
			city:    "New York",
			state:   "NY",
			country: "USA",
		},
	}
	fmt.Printf("name: %s, age: %d, college: %s, phone: %d\n", person1.name, person1.age, person1.college, person1.phone)
	fmt.Printf("name: %s, gender: %t, city: %s, state: %s", man1.name, man1.gender, man1.address.city, man1.address.state)

}
