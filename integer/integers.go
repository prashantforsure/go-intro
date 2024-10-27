package main

import "fmt"

func AddN(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello from main function!")
	integer()
}

func integer() {
	fmt.Println(AddN(6, 0))
}
